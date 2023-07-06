# tunetrail-api の CloudWatch Logs の設定
resource "aws_cloudwatch_log_group" "api_log_group" {
  name              = "tunetrail-api"
  retention_in_days = 14 # 14日間ログを保持する
}

# tunetrail-webapp の CloudWatch Logs の設定
resource "aws_cloudwatch_log_group" "webapp_log_group" {
  name              = "tunetrail-webapp"
  retention_in_days = 14
}
