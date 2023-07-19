output "user_pool_id" {
  description = "The ID of the user pool"
  type        = string
  value       = aws_cognito_user_pool.tunetrail.id
}

output "user_pool_client_id" {
  description = "The ID of the user pool client"
  type        = string
  value       = aws_cognito_user_pool_client.client.id
}
