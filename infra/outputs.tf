output "use_resources" {
  value       = var.use_resources
  description = "Current status of the resources"
}

output "webapp_image_tag" {
  value       = var.webapp_image_tag
  description = "The tag of the webapp image"
}

output "api_image_tag" {
  value       = var.api_image_tag
  description = "The tag of the API image"
}
