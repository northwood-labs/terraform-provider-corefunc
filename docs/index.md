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

## Updating your Terraform lockfile

Running `terraform init` will download the provider and update the [_Dependency Lock File_](https://developer.hashicorp.com/terraform/language/files/dependency-lock) (`.terraform.lock.hcl`) for your _current_ OS and CPU architecture. If you have a team with multiple operating systems or multiple CPU architectures, the _Dependency Lock File_ will be incomplete, and other members on the team won't be able to use it.

In order to resolve this, you can use the `terraform providers lock` command to generate a _Dependency Lock File_ that is compatible with all relevant operating systems and CPU architectures.

### Recommended matrix

Per [Recommended Provider Binary Operating Systems and Architectures](https://developer.hashicorp.com/terraform/registry/providers/os-arch):

```shell
terraform providers lock \
    -platform=darwin_amd64 \
    -platform=darwin_arm64 \
    -platform=linux_amd64 \
    -platform=linux_arm \
    -platform=linux_arm64 \
    -platform=windows_amd64 \
    ;
```

### Extended matrix

```shell
terraform providers lock \
    -platform=darwin_amd64 \
    -platform=darwin_arm64 \
    -platform=freebsd_386 \
    -platform=freebsd_amd64 \
    -platform=linux_386 \
    -platform=linux_amd64 \
    -platform=linux_arm \
    -platform=linux_arm64 \
    -platform=windows_386 \
    -platform=windows_amd64 \
    ;
```

### Complete matrix

This is the complete list of supported operating systems and architectures for this specific provider:

```shell
terraform providers lock \
    -platform=darwin_amd64 \
    -platform=darwin_arm64 \
    -platform=freebsd_386 \
    -platform=freebsd_amd64 \
    -platform=freebsd_arm \
    -platform=freebsd_arm64 \
    -platform=linux_386 \
    -platform=linux_amd64 \
    -platform=linux_arm \
    -platform=linux_arm64 \
    -platform=netbsd_386 \
    -platform=netbsd_amd64 \
    -platform=openbsd_386 \
    -platform=openbsd_amd64 \
    -platform=windows_386 \
    -platform=windows_amd64 \
    -platform=windows_arm \
    -platform=windows_arm64 \
    ;
```

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->
