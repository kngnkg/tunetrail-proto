resource "aws_ecr_repository" "api" {
  name = "tunetrail-api"
  /**
  開発フェーズでは、イメージのタグを変更することが多いため、
  イメージのタグを変更可能にしておく。
  */
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecr_repository" "frontend" {
  name                 = "tunetrail-frontend"
  image_tag_mutability = "MUTABLE"
}