terraform {
  required_version = "~> 1.0"

  required_providers {
    corefunc = {
      source  = "northwood-labs/corefunc"
      version = "~> 2.0"
    }
  }
}

provider "corefunc" {
  # There are no configuration options
}
