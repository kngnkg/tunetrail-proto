variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "private_subnet_ids" {
  description = "List of IDs of private subnets"
  type        = list(string)
  default     = []
}

variable "image_tag" {
  description = "The tag of the Docker image"
  type        = string
}

variable "s3_arn" {
  description = "The ARN of the S3 bucket"
  type        = string
}

variable "s3_bucket" {
  description = "The name of the S3 bucket"
  type        = string
}

variable "db_address" {
  description = "The address of the database"
  type        = string
}

variable "db_port" {
  description = "The port of the database"
  type        = number
}

variable "db_name" {
  description = "The name of the database"
  type        = string
}

variable "db_user" {
  description = "The user of the database"
  type        = string
}

variable "db_password" {
  description = "The password for the database"
  type        = string
}
