# Application Load Balancer (ALB)
resource "aws_lb" "tunetrail" {
  name               = "tunetrail-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_sg.id]
  subnets            = [aws_subnet.public1.id, aws_subnet.public2.id] # ALBを配置するパブリックサブネット
}

# HTTPリスナー
# HTTPでの接続はHTTPSへリダイレクトする
resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.tunetrail.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

# HTTPSリスナー
# ACMで発行した証明書を使用して、HTTPSでの接続を受け付ける
resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.tunetrail.arn # 上記で作成したALBを指定
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"       # デフォルトのセキュリティポリシーを指定
  certificate_arn   = aws_acm_certificate.tunetrail.arn # ACMで発行した証明書のARNを指定

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.alb_tg.arn # フォワード先のターゲットグループ
  }
}

# webapp用のリスナールール
# webappのドメインにアクセスした場合に、webapp用のターゲットグループにフォワードする
resource "aws_lb_listener_rule" "webapp" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 101

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.alb_tg_webapp.arn
  }

  condition {
    host_header {
      values = [var.webapp_domain]
    }
  }
}

# REST APIコンテナ用のリスナールール
# REST APIのドメインにアクセスした場合に、REST APIコンテナ用のターゲットグループにフォワードする
resource "aws_lb_listener_rule" "restapi" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.alb_tg_restapi.arn
  }

  condition {
    host_header {
      values = [var.restapi_domain]
    }
  }
}

# デフォルトのターゲットグループ
# デフォルトではwebappに接続する
resource "aws_lb_target_group" "alb_tg" {
  name     = "default-target-group"
  port     = var.webapp_port
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id # VPCを指定
}

# webapp用のターゲットグループ
resource "aws_lb_target_group" "alb_tg_webapp" {
  name        = "webapp-target-group"
  port        = var.webapp_port
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id # VPCを指定
  target_type = "ip"

  health_check {
    enabled             = true
    healthy_threshold   = 2  # 2回連続で正常なレスポンスを返すとヘルスチェックをパス
    unhealthy_threshold = 2  # 2回連続で異常なレスポンスを返すとヘルスチェックを不合格
    timeout             = 5  # 5秒以内にレスポンスを返さない場合はヘルスチェックを不合格
    interval            = 30 # 30秒ごとにヘルスチェックを実施
    path                = "/health"
    matcher             = "200-399" # 200番台と300番台のレスポンスを正常とする
    port                = var.webapp_port
    protocol            = "HTTP"
  }
}

# restapi用のターゲットグループ
resource "aws_lb_target_group" "alb_tg_restapi" {
  name        = "restapi-target-group"
  port        = var.restapi_port
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id # VPCを指定
  target_type = "ip"

  health_check {
    enabled             = true
    healthy_threshold   = 2  # 2回連続で正常なレスポンスを返すとヘルスチェックをパス
    unhealthy_threshold = 2  # 2回連続で異常なレスポンスを返すとヘルスチェックを不合格
    timeout             = 5  # 5秒以内にレスポンスを返さない場合はヘルスチェックを不合格
    interval            = 30 # 30秒ごとにヘルスチェックを実施
    path                = "/health"
    matcher             = "200-399" # 200番台と300番台のレスポンスを正常とする
    port                = var.restapi_port
    protocol            = "HTTP"
  }
}
