output "use_resources" {
  value       = var.use_resources
  description = "Current status of the resources"
}

output "frontend_image_tag" {
  value       = var.frontend_image_tag
  description = "The tag of the Frontend image"
}

output "api_image_tag" {
  value       = var.api_image_tag
  description = "The tag of the API image"
}
