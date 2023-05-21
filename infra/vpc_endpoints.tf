# ECRのAPIを呼び出すためのVPCエンドポイント
# イメージのメタデータを取得したり、イメージの認証トークンを取得するために使用される。
resource "aws_vpc_endpoint" "ecr_api" {
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.api"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private1.id, aws_subnet.private2.id]
  security_group_ids  = [aws_security_group.sg.id]
  private_dns_enabled = true
}

# Dockerイメージのプッシュ/プルを行うためのVPCエンドポイント
resource "aws_vpc_endpoint" "ecr_dkr" {
  vpc_id              = aws_vpc.main.id
  service_name        = "com.amazonaws.ap-northeast-1.ecr.dkr"
  vpc_endpoint_type   = "Interface"
  subnet_ids          = [aws_subnet.private1.id, aws_subnet.private2.id]
  security_group_ids  = [aws_security_group.sg.id]
  private_dns_enabled = true
}