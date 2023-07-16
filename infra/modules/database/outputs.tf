output "address" {
  description = "The address of the database"
  value       = aws_db_instance.tunetrail.address
}

output "port" {
  description = "The port of the database"
  value       = aws_db_instance.tunetrail.port
}

output "username" {
  description = "The username for the database"
  value       = aws_db_instance.tunetrail.username
}

output "name" {
  description = "The name of the database"
  value       = aws_db_instance.tunetrail.name
}
