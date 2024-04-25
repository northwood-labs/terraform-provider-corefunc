---
page_title: "Core Functions Provider"
subcategory: ""
description: |-
  Utilities that should have been Terraform/OpenTofu core functions.
  While some of these can be implemented in HCL, some of them begin to
  push up against the limits of Terraform and the HCL2 configuration
  language. We also perform testing using the
  Terratest https://terratest.gruntwork.io framework on a regular basis.
  Exposing these functions as both a Go library as well as a Terraform/OpenTofu
  provider enables us to use the same functionality in both our Terraform/OpenTofu
  applies as well as while using a testing framework.
  Since earlier versions of Terraform/OpenTofu didn't have the concept of
  user-defined functions, the next step to open up the possibilities was
  to write a custom Provider which has the functions built-in, using
  existing support for inputs and outputs.
  This does not add new syntax or constructs. Instead it uses the
  existing concepts around Providers, Resources, Data Sources,
  Variables, Outputs, and Functions (1.8) to expose new custom-built
  functionality.
  The goal of this provider is not to call any APIs, but to provide
  pre-built functions in the form of Data Sources or Provider
  Functions (1.8).
---

# Core Functions Provider

Utilities that should have been Terraform/OpenTofu _core functions_.

While some of these _can_ be implemented in HCL, some of them begin to
push up against the limits of Terraform and the HCL2 configuration
language. We also perform testing using the
[Terratest](https://terratest.gruntwork.io) framework on a regular basis.
Exposing these functions as both a Go library as well as a Terraform/OpenTofu
provider enables us to use the same functionality in both our Terraform/OpenTofu
applies as well as while using a testing framework.

Since earlier versions of Terraform/OpenTofu didn't have the concept of
user-defined functions, the next step to open up the possibilities was
to write a custom Provider which has the functions built-in, using
existing support for inputs and outputs.

**This does not add new syntax or constructs.** Instead it uses the
_existing_ concepts around Providers, Resources, Data Sources,
Variables, Outputs, and Functions (1.8) to expose new custom-built
functionality.

The goal of this provider is not to call any APIs, but to provide
pre-built functions in the form of _Data Sources_ or _Provider
Functions_ (1.8).

~> While it’s common knowledge that Terraform is great at standing up and managing Cloud infrastructure, it’s also good at running _anything with an API_. People regularly manage [code repositories], [DNS records], [feature flags], [identity and access management], [content delivery], [passwords], [monitoring], [alerts], [zero trust network access], [cryptographic signatures], and can even [order a pizza]. This provider is more analogous to HashiCorp’s _utility_ providers such as [local], [external], and [archive].

## Compatibility matrix

Built using the [Terraform Plugin Framework][TPF], which speaks [Terraform Protocol v6][tfproto6].

| Testing type | Details           | Description                                                                    |
|--------------|-------------------|--------------------------------------------------------------------------------|
| integration  | Terraform 1.0–1.8 | Executes the provider with this release, pulling from `registry.terraform.io`. |
| integration  | OpenTofu 1.6–1.7  | Executes the provider with this release, pulling from `registry.opentofu.org`. |
| unit         | Go 1.21–1.22      | Tests using these versions.                                                    |
| mutation     | Go 1.21–1.22      | Tests using these versions.                                                    |
| fuzz         | Go 1.21–1.22      | Tests using these versions.                                                    |
| terratest    | Go 1.21–1.22      | Tests using these versions.                                                    |

## Setting-up the provider

### Data Sources

For users of Terraform 1.0 (and newer), all of the the _Data Source_ implementations are available. Their implementations (inputs, outputs) are consistent with the _Provider Function_ implementations, but exposed as _Data Sources_.

```terraform
terraform {
  required_version = "~> 1.0"

  required_providers {
    corefunc = {
      source  = "northwood-labs/corefunc"
      version = "~> 1.0"
    }
  }
}

provider "corefunc" {
  # There are no configuration options
}
```

### Provider Functions

For users of Terraform 1.8/OpenTofu 1.7 (and newer), all of the the _Data Source_ and _Provider Function_ implementations are available. Their implementations (inputs, outputs) are consistent with each other, and will always return the same outputs from the same imputs. _Provider Functions_ require version 1.4.0 (or later) of this provider.

```terraform
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
```

## Updating your lockfile

Running `terraform init` will download the provider and update the [_Dependency Lock File_](https://developer.hashicorp.com/terraform/language/files/dependency-lock) (`.terraform.lock.hcl`) for your _current_ OS and CPU architecture. If you have a team with multiple operating systems or multiple CPU architectures, the _Dependency Lock File_ will be incomplete, and other members on the team won't be able to use it.

In order to resolve this, you can use the `terraform providers lock` command to generate a _Dependency Lock File_ that is compatible with all relevant operating systems and CPU architectures.

~> **NOTE:** For OpenTofu users, the `terraform` command can be replaced with `tofu`.

### Recommended matrix

Per [Recommended Provider Binary Operating Systems and Architectures](https://developer.hashicorp.com/terraform/registry/providers/os-arch):

```shell
#!/usr/bin/env bash
terraform providers lock \
    -platform=darwin_amd64 \
    -platform=darwin_arm64 \
    -platform=linux_amd64 \
    -platform=linux_arm \
    -platform=linux_arm64 \
    -platform=windows_amd64 \
    ;
```

[alerts]: https://registry.terraform.io/providers/PagerDuty/pagerduty/latest
[archive]: https://registry.terraform.io/providers/hashicorp/archive/latest/docs
[code repositories]: https://registry.terraform.io/providers/integrations/github/latest/docs
[content delivery]: https://registry.terraform.io/providers/fastly/fastly/latest/docs
[cryptographic signatures]: https://registry.terraform.io/providers/chainguard-dev/cosign/latest/docs
[DNS records]: https://registry.terraform.io/providers/infobloxopen/infoblox/latest/docs
[external]: https://registry.terraform.io/providers/hashicorp/external/latest/docs
[feature flags]: https://registry.terraform.io/providers/launchdarkly/launchdarkly/latest/docs
[identity and access management]: https://registry.terraform.io/providers/okta/okta/latest/docs
[local]: https://registry.terraform.io/providers/hashicorp/local/latest/docs
[monitoring]: https://registry.terraform.io/providers/DataDog/datadog/latest
[order a pizza]: https://registry.terraform.io/providers/MNThomson/dominos/latest/docs
[passwords]: https://registry.terraform.io/providers/1Password/onepassword/latest/docs
[tfproto6]: https://developer.hashicorp.com/terraform/plugin/terraform-plugin-protocol#protocol-version-6
[TPF]: https://github.com/hashicorp/terraform-plugin-framework
[zero trust network access]: https://registry.terraform.io/providers/zscaler/zpa/latest/docs

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->
