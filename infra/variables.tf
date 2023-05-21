variable "api_domain" {
  description = "The domain name for the API"
  type        = string
  default     = "api.tune-trail.com"
}

variable "api_image_tag" {
  description = "value of the tag for the API image"
  type        = string
}

variable "db_password" {
  description = "The password for the DB instance"
  type        = string
}
