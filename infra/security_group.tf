# ALB用のセキュリティグループ
resource "aws_security_group" "alb_sg" {
  name        = "allow_http_https"
  description = "Allow HTTP and HTTPS inbound traffic"
  vpc_id      = aws_vpc.main.id

  # HTTPに対するインバウンドルール
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPからのアクセスを許可
  }

  # HTTPSに対するインバウンドルール
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPからのアクセスを許可
  }
}

# ECSタスク用のセキュリティグループ
resource "aws_security_group" "sg" {
  name        = "ecs_tasks_sg"
  description = "Security Group for ECS Tasks"
  vpc_id      = aws_vpc.main.id

  # DBへのアクセス用のインバウンドルールの設定
  ingress {
    from_port   = 5432 # DBのポート番号
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # VPC内からのアクセスのみ許可
  }

  # アウトバウンドルールの設定
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPへのアクセスを許可
  }
}
