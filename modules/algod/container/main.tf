resource "docker_image" "this" {
  name = "${var.container_repo}:${var.container_tag}"
}

resource "docker_container" "this" {
  name  = var.name
  image = docker_image.this.image_id
  env   = local.container_env

  dynamic "ports" {
    for_each = local.container_ports

    content {
      internal = ports.value.internal
      external = ports.value.external
    }
  }
}
