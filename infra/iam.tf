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

# Lambda用のCloudWatch Logs への書き込み権限を持つポリシー
resource "aws_iam_policy" "lambda_cloudwatch" {
  name        = "CloudWatchLogsPolicyLambda"
  description = "Allow writing to CloudWatch Logs"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = "${aws_cloudwatch_log_group.migration_log_group.arn}",
        Effect   = "Allow"
      }
    ]
  })
}

# LambdaがVPC内で動作するために必要な権限を持つポリシー
resource "aws_iam_role_policy" "lambda_vpc_access" {
  name = "lambda_vpc_access"
  role = aws_iam_role.lambda_exec.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface"
        ],
        Resource = "*"
      }
    ]
  })
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

# lambdaの実行ロール
resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Principal = {
          Service = "lambda.amazonaws.com"
        },
        Effect = "Allow",
      },
    ]
  })
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

# lambdaの実行ロールに CloudWatch Logs への書き込み権限を持つポリシーをアタッチ
resource "aws_iam_role_policy_attachment" "cloudwatch_logs_policy_attach_lambda" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_cloudwatch.arn
}
