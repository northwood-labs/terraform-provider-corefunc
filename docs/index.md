<!--
---
page_title: "Core Functions Provider"
subcategory: ""
description: |-
  Utilities that should have been Terraform core functions.
  While some of these can be implemented in HCL, some of them begin to
  push up against the limits of Terraform and the HCL2 configuration
  language. We also perform testing using the
  Terratest https://terratest.gruntwork.io framework on a regular basis.
  Exposing these functions as both a Go library as well as a Terraform
  provider enables us to use the same functionality in both our Terraform
  applies as well as while using a testing framework.
  Since Terraform doesn't have the concept of user-defined functions, the
  next step to open up the possibilities is to write a custom Terraform
  Provider which has the functions built-in, using Terraform's existing
  support for inputs and outputs.
  This does not add new syntax or constructs to Terraform. Instead it
  uses the existing concepts around Providers, Resources, Data Sources,
  Variables, and Outputs to expose new custom-built functionality.
  The goal of this provider is not to call any APIs, but to provide
  pre-built functions in the form of Data Sources.
---
-->

# Core Functions Provider

-> We intend to support both Terraform as well as OpenTF with this provider.

Utilities that should have been Terraform _core functions_.

While some of these _can_ be implemented in HCL, some of them begin to
push up against the limits of Terraform and the HCL2 configuration
language. We also perform testing using the
[Terratest](https://terratest.gruntwork.io) framework on a regular basis.
Exposing these functions as both a Go library as well as a Terraform
provider enables us to use the same functionality in both our Terraform
applies as well as while using a testing framework.

Since Terraform doesn't have the concept of user-defined functions, the
next step to open up the possibilities is to write a custom Terraform
Provider which has the functions built-in, using Terraform's existing
support for inputs and outputs.

**This does not add new syntax or constructs to Terraform.** Instead it
uses the _existing_ concepts around Providers, Resources, Data Sources,
Variables, and Outputs to expose new custom-built functionality.

The goal of this provider is not to call any APIs, but to provide
pre-built functions in the form of _Data Sources_.

## Setting-up the provider

```terraform
terraform {
  required_version = "~> 1.1"

  required_providers {
    corefunc = {
      source  = "northwood-labs/corefunc"
      version = "~> 1.0"
    }
  }
}

# There are no configuration options
provider "corefunc" {}
```

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->
