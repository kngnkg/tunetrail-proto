# RDS用のサブネットグループ
resource "aws_db_subnet_group" "tunetrail" {
  name       = "tunetrail-db-subnet-group"
  subnet_ids = [aws_subnet.private.id]

  tags = {
    Name = "tunetrail-db-subnet-group"
  }
}

# RDSインスタンス
resource "aws_db_instance" "tunetrail" {
  allocated_storage    = 20
  storage_type         = "gp2" # 汎用SSD
  engine               = "postgres"
  engine_version       = "15.2"
  instance_class       = "db.t3.micro"
  name                 = "tunetrail"
  username             = "tunetrail"
  password             = var.db_password
  parameter_group_name = "default.postgres15"
  db_subnet_group_name = aws_db_subnet_group.tunetrail.name
  publicly_accessible  = false # インターネットからの直接のアクセスを拒否
}
