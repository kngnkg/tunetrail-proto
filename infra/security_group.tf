# ALB用のセキュリティグループ
resource "aws_security_group" "alb_sg" {
  name        = "allow_http_https"
  description = "Allow HTTP and HTTPS inbound traffic"
  vpc_id      = aws_vpc.main.id

  # HTTPに対するインバウンドルール
  ingress {
    from_port   = var.restapi_port
    to_port     = var.restapi_port
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

  # アウトバウンドルールの設定
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPへのアクセスを許可
  }
}

# webapp用のセキュリティグループ
resource "aws_security_group" "webapp_sg" {
  name        = "webapp_sg"
  description = "Security Group for webapp Tasks"
  vpc_id      = aws_vpc.main.id

  # restapiへのアクセス用のインバウンドルールの設定
  ingress {
    from_port   = var.webapp_port
    to_port     = var.webapp_port
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # VPC内からのアクセスのみ許可
  }

  # VPCエンドポイント用のインバウンドルールの設定
  ingress {
    from_port   = 443
    to_port     = 443
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

# REST API用のセキュリティグループ
resource "aws_security_group" "restapi_sg" {
  name        = "restapi_sg"
  description = "Security Group for REST API Tasks"
  vpc_id      = aws_vpc.main.id

  # REST APIへのアクセス用のインバウンドルールの設定
  ingress {
    from_port   = var.restapi_port
    to_port     = var.restapi_port
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # VPC内からのアクセスのみ許可
  }

  # VPCエンドポイント用のインバウンドルールの設定
  ingress {
    from_port   = 443
    to_port     = 443
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
