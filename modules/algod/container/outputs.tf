output "container_id" {
  description = "The ID of the container that was created"
  value       = docker_container.this.id
}

output "container_image" {
  description = "The Image used for the container that was created"
  value       = docker_container.this.image
}
