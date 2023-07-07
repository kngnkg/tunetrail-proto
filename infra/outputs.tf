output "use_resources" {
  value       = var.use_resources
  description = "Current status of the resources"
}

output "webapp_image_tag" {
  value       = var.webapp_image_tag
  description = "The tag of the webapp image"
}

output "restapi_image_tag" {
  value       = var.restapi_image_tag
  description = "The tag of the REST API image"
}

output "migration_image_tag" {
  value       = var.migration_image_tag
  description = "The tag of the migration image"
}
