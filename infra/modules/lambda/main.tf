# --------------------------------------------------
# Lambda関数用のECRリポジトリ
# --------------------------------------------------
resource "aws_ecr_repository" "migration" {
  name                 = "tunetrail-migration"
  image_tag_mutability = "MUTABLE"
}

# --------------------------------------------------
# マイグレーション用のLambda関数
# --------------------------------------------------
resource "aws_lambda_function" "migration" {
  function_name = "migration_lambda"
  image_uri     = "${aws_ecr_repository.migration.repository_url}:${var.image_tag}" # ECRのリポジトリURL
  role          = aws_iam_role.lambda_exec.arn
  timeout       = 60
  memory_size   = 128
  package_type  = "Image"
  vpc_config {
    subnet_ids         = var.private_subnet_ids
    security_group_ids = [aws_security_group.migration_sg.id]
  }
  environment {
    variables = {
      ENV                   = "prod"
      TUNETRAIL_S3_BUCKET   = var.s3_bucket
      TUNETRAIL_DB_HOST     = var.db_address
      TUNETRAIL_DB_PORT     = tostring(var.db_port)
      TUNETRAIL_DB_USER     = var.db_user
      TUNETRAIL_DB_PASSWORD = var.db_password
      TUNETRAIL_DB_NAME     = var.db_name
      DRY_RUN               = "false"
      ENABLE_DROP_TABLE     = "true"
    }
  }
}

# --------------------------------------------------
# マイグレーション時に起動するLambda用のセキュリティグループ
# --------------------------------------------------
resource "aws_security_group" "migration_sg" {
  name        = "migration_sg"
  description = "Security Group for migration"
  vpc_id      = var.vpc_id
  # VPCエンドポイント用のインバウンドルールの設定
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # VPC内からのアクセスのみ許可
  }
  # アウトバウンドルールの設定
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPへのアクセスを許可
  }
}

# --------------------------------------------------
# tunetrail-migration の CloudWatch Logs の設定
# --------------------------------------------------
resource "aws_cloudwatch_log_group" "migration_log_group" {
  name              = "/aws/lambda/${aws_lambda_function.migration.function_name}"
  retention_in_days = 14 # 14日間ログを保持する
}

# --------------------------------------------------
# lambdaの実行ロール
# --------------------------------------------------
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

# --------------------------------------------------
# lambdaの実行ロールにアタッチするポリシー
# --------------------------------------------------

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
        Resource = "${aws_cloudwatch_log_group.migration_log_group.arn}:*",
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

# LambdaがS3にアクセスするために必要な権限を持つポリシー
resource "aws_iam_policy" "lambda_s3_access" {
  name        = "lambda_s3_access"
  description = "Allows lambda function to access S3"
  policy      = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:PutObject",
        "s3:ListBucket"
      ],
      "Resource": [
        "${var.s3_arn}/*",
        "${var.s3_arn}"
      ]
    }
  ]
}
EOF
}

# Lambdaの実行ロールに CloudWatch Logs への書き込み権限を持つポリシーをアタッチ
resource "aws_iam_role_policy_attachment" "cloudwatch_logs_policy_attach_lambda" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_cloudwatch.arn
}

# Lambdaの実行ロールにS3へのアクセス権限をアタッチ
resource "aws_iam_role_policy_attachment" "lambda_s3_access_policy_attach" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_s3_access.arn
}
