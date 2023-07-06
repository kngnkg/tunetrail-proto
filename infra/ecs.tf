# ECS Cluster の作成
resource "aws_ecs_cluster" "tunetrail" {
  name = "tunetrail"
}

# tunetrail-api サービスの設定
resource "aws_ecs_service" "api" {
  name            = "tunetrail-api"
  cluster         = aws_ecs_cluster.tunetrail.id
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = var.use_resources ? 2 : 0 # タスクの数
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_groups  = [aws_security_group.api_sg.id]
    assign_public_ip = false
  }
  # ロードバランサーの設定
  # ターゲットグループをアタッチすることで、ロードバランサーにターゲットとして登録される
  load_balancer {
    target_group_arn = aws_lb_target_group.alb_tg_api.arn
    container_name   = "tunetrail-api"
    container_port   = var.api_port
  }
}

# tunetrail-api タスク定義の作成
resource "aws_ecs_task_definition" "api" {
  container_definitions = jsonencode([{
    name  = "tunetrail-api",
    image = "${aws_ecr_repository.api.repository_url}:${var.api_image_tag}", # ECRのリポジトリURL
    portMappings = [{
      containerPort = var.api_port
      protocol      = "tcp"
    }],
    environment = [
      {
        name  = "TUNETRAIL_ENV"
        value = "prod"
      },
      {
        name  = "PORT"
        value = "80"
      },
      {
        name  = "TUNETRAIL_DB_HOST"
        value = "${aws_db_instance.tunetrail.address}"
      },
      {
        name  = "TUNETRAIL_DB_PORT"
        value = tostring(aws_db_instance.tunetrail.port)
      },
      {
        name  = "TUNETRAIL_DB_USER"
        value = "${aws_db_instance.tunetrail.username}"
      },
      {
        name  = "TUNETRAIL_DB_PASSWORD"
        value = "${aws_db_instance.tunetrail.password}"
      },
      {
        name  = "TUNETRAIL_DB_NAME"
        value = "${aws_db_instance.tunetrail.name}"
      },
    ],
    logConfiguration = {
      logDriver = "awslogs", # CloudWatch Logsを使用する
      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.api_log_group.name}",
        awslogs-region        = "ap-northeast-1",
        awslogs-stream-prefix = "api"
      }
    },
  }])

  family                   = "tunetrail-api" # タスク定義のファミリー名
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.execution_role.arn # タスク実行ロール
  task_role_arn            = aws_iam_role.task_role.arn      # タスクロール
}

# tunetrail-webapp サービスの設定
resource "aws_ecs_service" "webapp" {
  name            = "tunetrail-webapp"
  cluster         = aws_ecs_cluster.tunetrail.id
  task_definition = aws_ecs_task_definition.webapp.arn
  desired_count   = var.use_resources ? 2 : 0 # タスクの数
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_groups  = [aws_security_group.webapp_sg.id]
    assign_public_ip = false
  }

  # ロードバランサーの設定
  # ターゲットグループをアタッチすることで、ロードバランサーにターゲットとして登録される
  load_balancer {
    target_group_arn = aws_lb_target_group.alb_tg_webapp.arn
    container_name   = "tunetrail-webapp"
    container_port   = var.webapp_port
  }
}

# tunetrail-webapp タスク定義の作成
resource "aws_ecs_task_definition" "webapp" {
  container_definitions = jsonencode([{
    name  = "tunetrail-webapp",
    image = "${aws_ecr_repository.webapp.repository_url}:${var.webapp_image_tag}",
    portMappings = [{
      containerPort = var.webapp_port
    }],
    logConfiguration = {
      logDriver = "awslogs", # CloudWatch Logsを使用する
      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.webapp_log_group.name}",
        awslogs-region        = "ap-northeast-1",
        awslogs-stream-prefix = "webapp"
      }
    },
  }])

  family                   = "tunetrail-webapp"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.execution_role.arn
  task_role_arn            = aws_iam_role.task_role.arn
}
