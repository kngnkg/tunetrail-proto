# GitHub Actions 用の IAM ロール
resource "aws_iam_role" "github_actions" {
  name = "github_actions_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

# GitHub Actions 用の IAM ポリシー
# ECRからイメージをプル、プッシュ、
# ECSにデプロイなどに必要な権限を持つポリシー
resource "aws_iam_role_policy" "github_actions_policy" {
  name = "github_actions_policy"
  role = aws_iam_role.github_actions.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:GetRepositoryPolicy",
        "ecr:DescribeRepositories",
        "ecr:ListImages",
        "ecr:DescribeImages",
        "ecr:BatchGetImage",
        "ecr:GetLifecyclePolicy",
        "ecr:GetLifecyclePolicyPreview",
        "ecr:ListTagsForResource",
        "ecr:DescribeImageScanFindings",
        "ecr:InitiateLayerUpload",
        "ecr:UploadLayerPart",
        "ecr:CompleteLayerUpload",
        "ecr:PutImage",
        "ecs:RegisterTaskDefinition",
        "ecs:UpdateService",
        "ecs:DescribeServices"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

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

# execution_roleにポリシーをアタッチ
/**
ECSタスクが、例えばECRからイメージをプルしたり、
CloudWatchにログを送信したりするための権限がexecution_roleに付与される
*/
resource "aws_iam_role_policy_attachment" "ecsTaskExecutionRole_policy" {
  role       = aws_iam_role.execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# ECS タスク実行ロールに CloudWatch Logs への書き込み権限を付与
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
        "${aws_cloudwatch_log_group.api_log_group.arn}",
        "${aws_cloudwatch_log_group.frontend_log_group.arn}"
      ]
    }
  ]
}
EOF
}

# ECS タスク実行ロールに CloudWatch Logs への書き込み権限を付与
resource "aws_iam_role_policy_attachment" "cloudwatch_logs_policy_attach" {
  role       = aws_iam_role.execution_role.name
  policy_arn = aws_iam_policy.cloudwatch_logs_policy.arn
}
