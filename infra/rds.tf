# RDSのサブネットグループ
## 高可用性を担保するために、複数のAZにまたがるサブネットを指定する
resource "aws_db_subnet_group" "tunetrail" {
  name       = "tunetrail-db-subnet-group"
  subnet_ids = [aws_subnet.private1.id, aws_subnet.private2.id]

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
