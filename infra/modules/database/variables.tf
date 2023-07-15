variable "vpc_id" {
  description = "The ID of the VPC"
  type        = string
}

variable "private_subnet_ids" {
  description = "List of IDs of private subnets"
  type        = list(string)
  default     = []
}

variable "db_port" {
  description = "The port of the database"
  type        = number
}

variable "db_password" {
  description = "The password for the database"
  type        = string
}
