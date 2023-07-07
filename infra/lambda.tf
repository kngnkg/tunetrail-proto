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

  environment {
    variables = {
      TUNETRAIL_DB_HOST     = "${aws_db_instance.tunetrail.address}"
      TUNETRAIL_DB_PORT     = tostring(aws_db_instance.tunetrail.port)
      TUNETRAIL_DB_USER     = "${aws_db_instance.tunetrail.username}"
      TUNETRAIL_DB_PASSWORD = "${aws_db_instance.tunetrail.password}"
      TUNETRAIL_DB_NAME     = "${aws_db_instance.tunetrail.name}"
    }
  }
}
