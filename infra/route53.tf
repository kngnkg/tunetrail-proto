# Route53 Zone
resource "aws_route53_zone" "tunetrail" {
  name = "tune-trail.com"
}

# ALBへのエイリアスレコード (www)
resource "aws_route53_record" "www" {
  zone_id = aws_route53_zone.tunetrail.zone_id
  name    = "www.tune-trail.com"
  type    = "A"

  alias {
    name                   = aws_lb.tunetrail.dns_name
    zone_id                = aws_lb.tunetrail.zone_id
    evaluate_target_health = true
  }
}

# ALBへのエイリアスレコード (api)
resource "aws_route53_record" "api" {
  zone_id = aws_route53_zone.tunetrail.zone_id
  name    = "api.tune-trail.com"
  type    = "A"

  alias {
    name                   = aws_lb.tunetrail.dns_name
    zone_id                = aws_lb.tunetrail.zone_id
    evaluate_target_health = true
  }
}
