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

output "tags" {
  value = provider::corefunc::map_key_if(var.pii, local.tags, "pii", true)
}

#=> {
#      application = "example"
#      environment = "production"
#      pii         = true
#    }
