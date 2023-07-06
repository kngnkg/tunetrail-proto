# tunetrail-api の CloudWatch Logs の設定
resource "aws_cloudwatch_log_group" "restapi_log_group" {
  name              = "tunetrail-restapi"
  retention_in_days = 14 # 14日間ログを保持する
}

# tunetrail-webapp の CloudWatch Logs の設定
resource "aws_cloudwatch_log_group" "webapp_log_group" {
  name              = "tunetrail-webapp"
  retention_in_days = 14
}
