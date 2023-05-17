# ECS Cluster の作成
resource "aws_ecs_cluster" "tunetrail" {
  name = "tunetrail"
}

# tunetrail-api サービスの設定
resource "aws_ecs_service" "api" {
  name            = "tunetrail-api"
  cluster         = aws_ecs_cluster.tunetrail.id
  task_definition = aws_ecs_task_definition.api.arn
  desired_count   = 2 # タスクの数
  network_configuration {
    subnets          = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_groups  = [aws_security_group.sg.id]
    assign_public_ip = false
  }
}

# tunetrail-api タスク定義の作成
resource "aws_ecs_task_definition" "api" {
  container_definitions = jsonencode([{
    name  = "tunetrail-api",
    image = "${aws_ecr_repository.api.repository_url}:latest", # ECRのリポジトリURL
    portMappings = [{
      containerPort = 8080
    }],
  }])

  family                   = "tunetrail-api" # タスク定義のファミリー名
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.execution_role.arn # タスク実行ロール
  task_role_arn            = aws_iam_role.task_role.arn      # タスクロール
}

# tunetrail-frontend サービスの設定
resource "aws_ecs_service" "frontend" {
  name            = "tunetrail-frontend"
  cluster         = aws_ecs_cluster.tunetrail.id
  task_definition = aws_ecs_task_definition.frontend.arn
  desired_count   = 2
  network_configuration {
    subnets          = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_groups  = [aws_security_group.sg.id]
    assign_public_ip = false
  }
}

# tunetrail-frontend タスク定義の作成
resource "aws_ecs_task_definition" "frontend" {
  container_definitions = jsonencode([{
    name  = "tunetrail-frontend",
    image = "${aws_ecr_repository.frontend.repository_url}:latest",
    portMappings = [{
      containerPort = 3000
    }],
  }])

  family                   = "tunetrail-frontend"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.execution_role.arn
  task_role_arn            = aws_iam_role.task_role.arn
}
