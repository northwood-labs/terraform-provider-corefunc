terraform {
  required_version = "~> 1.8"

  required_providers {
    corefunc = {
      source  = "northwood-labs/corefunc"
      version = "~> 1.5"
    }
  }
}

# There are no configuration options
provider "corefunc" {}
