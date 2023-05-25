# ECRのAPIを呼び出すためのVPCエンドポイント
# イメージのメタデータを取得したり、イメージの認証トークンを取得するために使用される。
resource "aws_vpc_endpoint" "ecr_api" {
  count               = var.use_resources ? 1 : 0
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.api"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private1.id, aws_subnet.private2.id]
  security_group_ids  = [aws_security_group.api_sg.id, aws_security_group.frontend_sg.id]
  private_dns_enabled = true
}

# Dockerイメージのプッシュ/プルを行うためのVPCエンドポイント
resource "aws_vpc_endpoint" "ecr_dkr" {
  count               = var.use_resources ? 1 : 0
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.dkr"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private1.id, aws_subnet.private2.id]
  security_group_ids  = [aws_security_group.api_sg.id, aws_security_group.frontend_sg.id]
  private_dns_enabled = true
}

# S3用のVPCエンドポイント
# ECRのイメージをプッシュ/プルする際に、S3のバケットを使用するために必要。
resource "aws_vpc_endpoint" "s3" {
  count             = var.use_resources ? 1 : 0
  vpc_id            = aws_vpc.main.id
  service_name      = "com.amazonaws.ap-northeast-1.s3"
  vpc_endpoint_type = "Gateway"
  route_table_ids   = [aws_route_table.private.id]
}

# CloudWatch Logs用のVPCエンドポイント
resource "aws_vpc_endpoint" "cloudwatch_logs" {
  count               = var.use_resources ? 1 : 0
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.ap-northeast-1.logs"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private1.id, aws_subnet.private2.id]
  security_group_ids  = [aws_security_group.api_sg.id, aws_security_group.frontend_sg.id]
  private_dns_enabled = true
}

