# 無駄な課金を防ぐために、リソースを使用しない間はfalseにする。
variable "use_resources" {
  description = "Whether to use resources or data sources"
  type        = bool
  default     = true
}

variable "region" {
  type    = string
  default = "ap-northeast-1"
}

variable "webapp_image_tag" {
  description = "value of the tag for the webapp image"
  type        = string
}

variable "restapi_image_tag" {
  description = "value of the tag for the REST API image"
  type        = string
}

variable "migration_image_tag" {
  description = "value of the tag for the migration image"
  type        = string
}

variable "db_password" {
  description = "The password for the DB instance"
  type        = string
}

variable "db_schema_bucket_name" {
  description = "The name of the S3 bucket"
  type        = string
}

# variable "cognito_client_secret" {
#   description = "The secret for the Cognito user pool client"
#   type        = string
# }

variable "webapp_domain" {
  description = "The domain name for the webapp"
  type        = string
  default     = "tune-trail.com"
}

variable "webapp_port" {
  description = "The port for the webapp"
  type        = number
  default     = 3000
}

variable "restapi_domain" {
  description = "The domain name for the REST API"
  type        = string
  default     = "api.tune-trail.com"
}

variable "restapi_port" {
  description = "The port for the REST API"
  type        = number
  default     = 80
}

variable "db_port" {
  description = "The port for the DB"
  type        = number
  default     = 5432
}
