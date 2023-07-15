terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"
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
