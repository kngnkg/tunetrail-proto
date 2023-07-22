# ECS Cluster の作成
resource "aws_ecs_cluster" "tunetrail" {
  name = "tunetrail"
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


# tunetrail-restapi サービスの設定
resource "aws_ecs_service" "restapi" {
  name            = "tunetrail-restapi"
  cluster         = aws_ecs_cluster.tunetrail.id
  task_definition = aws_ecs_task_definition.restapi.arn
  desired_count   = var.use_resources ? 2 : 0 # タスクの数
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_groups  = [aws_security_group.restapi_sg.id]
    assign_public_ip = false
  }
  # ロードバランサーの設定
  # ターゲットグループをアタッチすることで、ロードバランサーにターゲットとして登録される
  load_balancer {
    target_group_arn = aws_lb_target_group.alb_tg_restapi.arn
    container_name   = "tunetrail-restapi"
    container_port   = var.restapi_port
  }
}

# tunetrail-restapi タスク定義の作成
resource "aws_ecs_task_definition" "restapi" {
  container_definitions = jsonencode([{
    name  = "tunetrail-restapi",
    image = "${aws_ecr_repository.restapi.repository_url}:${var.restapi_image_tag}", # ECRのリポジトリURL
    portMappings = [{
      containerPort = var.restapi_port
      protocol      = "tcp"
    }],
    environment = [
      {
        name  = "TUNETRAIL_ENV"
        value = "prod"
      },
      {
        name  = "PORT"
        value = tostring(var.restapi_port)
      },
      {
        name  = "TUNETRAIL_DB_HOST"
        value = module.database.address
      },
      {
        name  = "TUNETRAIL_DB_PORT"
        value = tostring(module.database.port)
      },
      {
        name  = "TUNETRAIL_DB_USER"
        value = module.database.username
      },
      {
        name  = "TUNETRAIL_DB_PASSWORD"
        value = var.db_password
      },
      {
        name  = "TUNETRAIL_DB_NAME"
        value = module.database.name
      },
      {
        name  = "TUNETRAIL_AWS_REGION"
        value = var.region
      },
      {
        name  = "COGNITO_USER_POOL_ID"
        value = "${module.auth.user_pool_id}"
      },
      {
        name  = "COGNITO_CLIENT_ID"
        value = "${module.auth.user_pool_client_id}"
      },
      {
        name  = "COGNITO_CLIENT_SECRET"
        value = var.cognito_client_secret
      },
    ],
    logConfiguration = {
      logDriver = "awslogs", # CloudWatch Logsを使用する
      options = {
        awslogs-group         = "${aws_cloudwatch_log_group.restapi_log_group.name}",
        awslogs-region        = "ap-northeast-1",
        awslogs-stream-prefix = "restapi"
      }
    },
  }])

  family                   = "tunetrail-restapi" # タスク定義のファミリー名
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.execution_role.arn # タスク実行ロール
  task_role_arn            = aws_iam_role.task_role.arn      # タスクロール
}
