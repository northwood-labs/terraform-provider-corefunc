variable "pii" {
  description = "Whether or not this app processes personally identifiable information."
  type        = bool
  default     = false
}

locals {
  tags = {
    application = "example"
    environment = "production"
  }
}

data "corefunc_map_key_if" "tags" {
  condition = var.pii
  map       = local.tags
  key       = "pii"
  value     = true
}

#=> {
#      application = "example"
#      environment = "production"
#      pii         = true
#    }
