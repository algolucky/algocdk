locals {
  # set tokens, if not already
  admin_token = var.admin_token == "" ? var.token : var.admin_token
  kmd_token   = var.kmd_token == "" ? local.admin_token : var.kmd_token

  # container environment variables
  container_env = compact([
    var.start_kmd ? "START_KMD=1" : "",
    var.start_kmd ? "KMD_PORT=${var.kmd_port}" : "",
    var.fast_catchup ? "FAST_CATCHUP=1" : "",
    var.start_kmd ? "KMD_TOKEN=${local.kmd_token}" : "",
    "TOKEN=${var.token}",
    "ADMIN_TOKEN=${local.admin_token}",
    var.network != "" ? "NETWORK=${var.network}" : ""
  ])

  container_ports = concat([
    {
      internal = 8080
      external = var.algod_port
    },
    var.start_kmd ? {
      internal = 7833
      external = var.kmd_port
    } : {}
  ])
}
