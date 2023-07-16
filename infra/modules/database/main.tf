# RDSのサブネットグループ
## 高可用性を担保するために、複数のAZにまたがるサブネットを指定する
resource "aws_db_subnet_group" "tunetrail" {
  name       = "tunetrail-db-subnet-group"
  subnet_ids = var.private_subnet_ids
  tags = {
    Name = "tunetrail-db-subnet-group"
  }
}

# RDSインスタンス
resource "aws_db_instance" "tunetrail" {
  allocated_storage      = 10
  storage_type           = "gp2" # 汎用SSD
  engine                 = "postgres"
  engine_version         = "15.2"
  instance_class         = "db.t3.micro"
  name                   = "tunetrail"
  username               = "tunetrail"
  password               = var.db_password
  parameter_group_name   = "default.postgres15"
  db_subnet_group_name   = aws_db_subnet_group.tunetrail.name
  publicly_accessible    = false # インターネットからの直接のアクセスを拒否
  skip_final_snapshot    = true  # 削除時にスナップショットを作成しない
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
}

# RDS用のセキュリティグループ
resource "aws_security_group" "rds_sg" {
  name        = "rds_sg"
  description = "Security Group for RDS Instance"
  vpc_id      = var.vpc_id
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
