resource "aws_lb" "alb" {
  name               = "tunetrail-alb"
  internal           = false # 内部向けのALBではなく、インターネット向けのALB
  load_balancer_type = "application"
  security_groups    = [aws_security_group.alb_sg.id]
  subnets            = [aws_subnet.public1.id, aws_subnet.public2.id]
}

# ALBのリスナー
resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_lb.alb.arn
  port              = "80"
  protocol          = "HTTP"

  # デフォルトアクション（リクエストが他のルールにマッチしない場合に適用される）
  default_action {
    type             = "forward"                      # リクエストをフォワード
    target_group_arn = aws_lb_target_group.alb_tg.arn # フォワード先のターゲットグループ
  }
}

# ALBのターゲットグループ
resource "aws_lb_target_group" "alb_tg" {
  name     = "tunetrail-alb-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.main.id
}

# ALBに適用するセキュリティグループ
resource "aws_security_group" "alb_sg" {
  name        = "allow_http"
  description = "Allow HTTP inbound traffic"

  # インバウンドルールの設定
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # 任意のIPからのアクセスを許可
  }
}
