# Route53 Zone
resource "aws_route53_zone" "tunetrail" {
  name = "tune-trail.com"
}

# ACM証明書
# resource "aws_acm_certificate" "tunetrail" {
#   domain_name               = "*.tune-trail.com"
#   subject_alternative_names = ["tune-trail.com"]
#   validation_method         = "DNS"
#   lifecycle {
#     create_before_destroy = true
#   }
# }

# CNAMEレコード
# resource "aws_route53_record" "cert_validation" {
#   # 複数のドメインに対する検証レコードを作成する
#   # 各ドメイン名をキーとし、対応するレコード情報を値とするマップを作成
#   for_each = {
#     for dvo in aws_acm_certificate.tunetrail.domain_validation_options : dvo.domain_name => {
#       name   = dvo.resource_record_name
#       record = dvo.resource_record_value
#       type   = dvo.resource_record_type
#     }
#   }

#   allow_overwrite = true # 既存のレコードを上書きする
#   name            = each.value.name
#   records         = [each.value.record]
#   ttl             = 60
#   type            = each.value.type
#   zone_id         = aws_route53_zone.tunetrail.zone_id
# }

# ACM証明書のドメイン検証
# 検証が完了するまで待機する
# resource "aws_acm_certificate_validation" "cert" {
#   certificate_arn = aws_acm_certificate.tunetrail.arn
#   # aws_route53_record.cert_validation の全てのインスタンスを取得し
#   # そのfqdn属性をリストとして取得
#   validation_record_fqdns = values(aws_route53_record.cert_validation)[*].fqdn
# }

# # ALBへのエイリアスレコード (www)
# resource "aws_route53_record" "www" {
#   zone_id = aws_route53_zone.tunetrail.zone_id
#   name    = "www.tune-trail.com"
#   type    = "A"

#   alias {
#     name                   = aws_lb.tunetrail.dns_name
#     zone_id                = aws_lb.tunetrail.zone_id
#     evaluate_target_health = true
#   }
# }

# ALBへのエイリアスレコード (api)
# resource "aws_route53_record" "api" {
#   zone_id = aws_route53_zone.tunetrail.zone_id
#   name    = "api.tune-trail.com"
#   type    = "A"

#   alias {
#     name                   = aws_lb.tunetrail.dns_name
#     zone_id                = aws_lb.tunetrail.zone_id
#     evaluate_target_health = true
#   }
# }
