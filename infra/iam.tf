# ------------------------------
# IAM ポリシー
# ------------------------------

# ECS用のCloudWatch Logs への書き込み権限を持つポリシー
resource "aws_iam_policy" "cloudwatch_logs_policy" {
  name        = "CloudWatchLogsPolicy"
  description = "Allow writing to CloudWatch Logs"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": [
        "${aws_cloudwatch_log_group.restapi_log_group.arn}",
        "${aws_cloudwatch_log_group.webapp_log_group.arn}"
      ]
    }
  ]
}
EOF
}

# ------------------------------
# IAM ロール
# ------------------------------

# ECS タスク実行ロール
## タスクの実行に必要な基本的な権限を持つロール
resource "aws_iam_role" "execution_role" {
  name = "ecsTaskExecutionRole"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

# ECS タスクロール
## タスク内から起動されるコンテナがAWSリソースにアクセスするための権限を持つロール
resource "aws_iam_role" "task_role" {
  name = "ecsTaskRole"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

# ------------------------------
# IAM ロールポリシーアタッチメント
# ------------------------------

# ECSタスク実行ロールにECSタスクがECRからイメージをpullする権限をアタッチ
resource "aws_iam_role_policy_attachment" "ecsTaskExecutionRole_policy" {
  role       = aws_iam_role.execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# ECSタスク実行ロールにS3 ReadOnlyアクセスポリシーをアタッチ
## ECRからイメージをpullする際に必要
resource "aws_iam_role_policy_attachment" "s3_read_only" {
  role       = aws_iam_role.execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
}

# ECSタスク実行ロールに CloudWatch Logs への書き込み権限を持つポリシーをアタッチ
resource "aws_iam_role_policy_attachment" "cloudwatch_logs_policy_attach" {
  role       = aws_iam_role.execution_role.name
  policy_arn = aws_iam_policy.cloudwatch_logs_policy.arn
}
