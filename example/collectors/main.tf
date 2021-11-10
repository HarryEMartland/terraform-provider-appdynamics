provider "appdynamics" {
  secret              = var.secret
  controller_base_url = var.controller_url
}

terraform {
  required_version = "= 0.13.4"

  required_providers {

    appdynamics = {
      source  = "worldremit/appdynamics"
      version = "0.1.0-0"
    }
  }
}

resource appdynamics_collector test {
  name       = "example2"
  type       = "MYSQL"
  hostname   = "test2"
  username   = "user"
  password   = "paswd3"
  port       = 17
  agent_name = "test"
  enabled    = true
}

resource appdynamics_collector test2 {
  name       = "example1"
  type       = "MYSQL"
  hostname   = "test2"
  username   = "u"
  password   = "password2"
  port       = 17
  agent_name = "test"
  enabled    = true
}

variable secret {}

variable controller_url {}