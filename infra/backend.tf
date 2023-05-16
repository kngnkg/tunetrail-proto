terraform {
  backend "s3" {
    bucket = "tunetrail-terraform-state" # S3バケット名
    key    = "terraform.tfstate"         # ステートファイル名
    region = "ap-northeast-1"
  }
}
