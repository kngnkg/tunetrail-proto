variable "api_image_tag" {
  description = "value of the tag for the API image"
  type        = string
}

variable "db_password" {
  description = "The password for the DB instance"
  type        = string
}

variable "frontend_domain" {
  description = "The domain name for the frontend"
  type        = string
  default     = "www.tune-trail.com"
}

variable "frontend_port" {
  description = "The port for the frontend"
  type        = number
  default     = 3000
}

variable "api_domain" {
  description = "The domain name for the API"
  type        = string
  default     = "api.tune-trail.com"
}

variable "api_port" {
  description = "The port for the API"
  type        = number
  default     = 80
}

variable "db_port" {
  description = "The port for the DB"
  type        = number
  default     = 5432
}
