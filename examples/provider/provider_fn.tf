terraform {
  required_version = "~> 1.8" # Terraform
  # required_version = "~> 1.7" # OpenTofu

  required_providers {
    corefunc = {
      source  = "northwood-labs/corefunc"
      version = "~> 1.4"
    }
  }
}

provider "corefunc" {
  # There are no configuration options
}
