variable "container_repo" {
  description = "The specific container to use"
  type        = string
  default     = "docker.io/algorand/algod"
}

variable "container_tag" {
  description = "The specific tag of the container to use"
  type        = string
  default     = "latest"
}

variable "name" {
  description = "Canonical name to give the container"
  type        = string
  default     = "algod"
}

variable "algod_port" {
  description = "The host port to expose the algod REST API on"
  type        = number
  default     = 8080
}

variable "start_kmd" {
  description = "Whether or not to start the kmd service"
  type        = bool
  default     = false
}

variable "kmd_port" {
  description = "The host port to expose the kmd REST API on"
  type        = number
  default     = 8081
}

variable "kmd_token" {
  description = "The token to use for kmd REST API requests"
  type        = string
  default     = ""
}

variable "network" {
  description = "The name of the network "
  type        = string
  default     = "betanet"
}

variable "token" {
  description = "The token to use for algod REST API requests"
  type        = string
  default     = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}

variable "admin_token" {
  description = "The token to use for algod admin REST API requests"
  type        = string
  default     = ""
}

variable "fast_catchup" {
  description = "Whether or not to start fast-catchup once algod has started"
  type        = string
  default     = false
}
