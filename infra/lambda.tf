# マイグレーション用のLambda関数
resource "aws_lambda_function" "migration" {
  function_name = "migration_lambda"
  image_uri     = "${aws_ecr_repository.migration.repository_url}:${var.migration_image_tag}" # ECRのリポジトリURL
  role          = aws_iam_role.lambda_exec.arn
  # The Lambda's timeout value
  timeout = 60
  # Memory size for your function
  memory_size  = 128
  package_type = "Image"

  vpc_config {
    subnet_ids         = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_group_ids = [aws_security_group.migration_sg.id]
  }

  environment {
    variables = {
      ENV                   = "prod"
      TUNETRAIL_S3_BUCKET   = "${aws_s3_bucket.schema.bucket}"
      TUNETRAIL_DB_HOST     = module.database.address
      TUNETRAIL_DB_PORT     = tostring(module.database.port)
      TUNETRAIL_DB_USER     = module.database.username
      TUNETRAIL_DB_PASSWORD = var.db_password
      TUNETRAIL_DB_NAME     = module.database.name
      DRY_RUN               = "false"
    }
  }
}
