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

# ALBのターゲットグループ
resource "aws_lb_target_group" "alb_tg" {
  name     = "tunetrail-alb-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id # VPCを指定
}

# ALBに適用するセキュリティグループ
resource "aws_security_group" "alb_sg" {
  name        = "allow_http_https"
  description = "Allow HTTP and HTTPS inbound traffic"
  vpc_id      = aws_vpc.main.id

  # HTTPに対するインバウンドルール
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPからのアクセスを許可
  }

  # HTTPSに対するインバウンドルール
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPからのアクセスを許可
  }
}
