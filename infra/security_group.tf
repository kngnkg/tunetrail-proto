# ALB用のセキュリティグループ
resource "aws_security_group" "alb_sg" {
  name        = "allow_http_https"
  description = "Allow HTTP and HTTPS inbound traffic"
  vpc_id      = aws_vpc.main.id

  # HTTPに対するインバウンドルール
  ingress {
    from_port   = var.api_port
    to_port     = var.api_port
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

# Frontend用のセキュリティグループ
resource "aws_security_group" "frontend_sg" {
  name        = "frontend_sg"
  description = "Security Group for Frontend Tasks"
  vpc_id      = aws_vpc.main.id

  # APIへのアクセス用のインバウンドルールの設定
  ingress {
    from_port   = var.frontend_port
    to_port     = var.frontend_port
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

# API用のセキュリティグループ
resource "aws_security_group" "api_sg" {
  name        = "api_sg"
  description = "Security Group for API Tasks"
  vpc_id      = aws_vpc.main.id

  # APIへのアクセス用のインバウンドルールの設定
  ingress {
    from_port   = var.api_port
    to_port     = var.api_port
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

# RDS用のセキュリティグループ
resource "aws_security_group" "rds_sg" {
  name        = "rds_sg"
  description = "Security Group for RDS Instance"
  vpc_id      = aws_vpc.main.id

  # RDSへのアクセス用のインバウンドルールの設定
  ingress {
    from_port   = var.db_port # DBのポート番号
    to_port     = var.db_port
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
