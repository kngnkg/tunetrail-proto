# tunetrail-api の CloudWatch Logs の設定
resource "aws_cloudwatch_log_group" "api_log_group" {
  name              = "tunetrail-api"
  retention_in_days = 14 # 14日間ログを保持する
}

# tunetrail-frontend の CloudWatch Logs の設定
resource "aws_cloudwatch_log_group" "frontend_log_group" {
  name              = "tunetrail-frontend"
  retention_in_days = 14
}
