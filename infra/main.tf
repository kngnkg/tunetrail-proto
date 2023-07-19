terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = var.region
}

module "database" {
  source = "./modules/database"
  vpc_id = aws_vpc.main.id
  private_subnet_ids = [
    aws_subnet.private1.id,
    aws_subnet.private2.id
  ]
  db_port     = var.db_port
  db_password = var.db_password
}

module "migration_lambda" {
  source = "./modules/lambda"
  vpc_id = aws_vpc.main.id
  private_subnet_ids = [
    aws_subnet.private1.id,
    aws_subnet.private2.id
  ]
  image_tag   = var.migration_image_tag
  s3_arn      = aws_s3_bucket.schema.arn
  s3_bucket   = var.db_schema_bucket_name
  db_address  = module.database.address
  db_name     = module.database.name
  db_user     = module.database.username
  db_port     = var.db_port
  db_password = var.db_password
}

# module "auth" {
#   source = "./modules/auth"
# }
