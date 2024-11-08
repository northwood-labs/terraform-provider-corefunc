# Terraform/OpenTofu Provider: Core Functions

[![Terraform Docs](https://img.shields.io/badge/Terraform-Docs-7B42BC?style=for-the-badge)](https://registry.terraform.io/providers/northwood-labs/corefunc/)
[![Library.tf](https://img.shields.io/badge/Library.tf-Docs-B3DBF1?style=for-the-badge)](https://library.tf/providers/northwood-labs/corefunc/latest)
[![Go Docs](https://img.shields.io/badge/Go-Docs-blue?style=for-the-badge)](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc)
[![GitHub](https://img.shields.io/github/license/northwood-labs/terraform-provider-corefunc?style=for-the-badge)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/northwood-labs/terraform-provider-corefunc?style=for-the-badge)](https://goreportcard.com/report/github.com/northwood-labs/terraform-provider-corefunc)
[![Open Source Insights](https://img.shields.io/badge/Open_Source-Insights-000000?style=for-the-badge)](https://deps.dev/project/github/northwood-labs%2Fterraform-provider-corefunc)
[![GitHub issues](https://img.shields.io/github/issues/northwood-labs/terraform-provider-corefunc?style=for-the-badge)](https://github.com/northwood-labs/terraform-provider-corefunc/issues)
[![GitHub contributors](https://img.shields.io/github/contributors/northwood-labs/terraform-provider-corefunc?style=for-the-badge)](https://github.com/northwood-labs/terraform-provider-corefunc/graphs/contributors)
[![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/m/northwood-labs/terraform-provider-corefunc?style=for-the-badge)](https://github.com/northwood-labs/terraform-provider-corefunc/commits/main/)
[![GitHub all releases](https://img.shields.io/github/downloads/northwood-labs/terraform-provider-corefunc/total?style=for-the-badge)](https://github.com/northwood-labs/terraform-provider-corefunc/releases)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/northwood-labs/terraform-provider-corefunc/test.yml?style=for-the-badge&label=Tests)](https://github.com/northwood-labs/terraform-provider-corefunc/actions/workflows/test.yml)
[![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/northwood-labs/terraform-provider-corefunc/release.yml?style=for-the-badge&label=Release)](https://github.com/northwood-labs/terraform-provider-corefunc/actions/workflows/release.yml)

## Overview

Utilities that should have been Terraform/OpenTofu _core functions_.

While some of these _can_ be implemented in HCL, some of them begin to push up against the limits of Terraform and the HCL2 configuration language. We also perform testing using the [Terratest](https://terratest.gruntwork.io) framework on a regular basis. Exposing these functions as both a Go library as well as a Terraform/OpenTofu provider enables us to use the same functionality in both our Terraform/OpenTofu applies as well as while using a testing framework.

> [!NOTE]
> While it’s common knowledge that Terraform is great at standing up and managing Cloud infrastructure, it’s also good at running _anything with an API_. People regularly manage [code repositories], [DNS records], [feature flags], [identity and access management], [content delivery], [passwords], [monitoring], [alerts], [zero trust network access], [cryptographic signatures], and can even [order a pizza].
>
> This provider is more analogous to HashiCorp’s _utility_ providers such as [local], [external], and [archive].

Since earlier versions of Terraform/OpenTofu didn't have the concept of user-defined functions, the next step to open up the possibilities was to write a custom Provider which has the functions built-in, using existing support for inputs and outputs.

**This does not add new syntax or constructs.** Instead it uses the _existing_ concepts around Providers, Resources, Data Sources, Variables, Outputs, and Functions to expose new custom-built functionality.

The goal of this provider is not to call any APIs, but to provide pre-built functions in the form of _Data Sources_ or _Provider Functions_.

## Compatibility testing

* We have automated testing that runs on every commit and every pull request.
* We intend for the Go libraries to work with all non-EOL versions of Go (i.e., current, current-1).
* Built using the [Terraform Plugin Framework][TPF], which speaks [Terraform Protocol v6][tfproto6].

| Testing type | Details            | Description                                                                    |
|--------------|--------------------|--------------------------------------------------------------------------------|
| integration  | Terraform 1.0–1.10 | Executes the provider with this release, pulling from `registry.terraform.io`. |
| integration  | OpenTofu 1.6–1.9   | Executes the provider with this release, pulling from `registry.opentofu.org`. |
| unit         | Go 1.22–1.23       | Tests using these versions.                                                    |
| mutation     | Go 1.22–1.23       | Tests using these versions.                                                    |
| fuzz         | Go 1.22–1.23       | Tests using these versions.                                                    |
| terratest    | Go 1.22–1.23       | Tests using these versions.                                                    |

## Usage Examples

See the `docs/` directory for user-facing documentation.

## Documentation

### Terraform Registry

If you are using this as a Terraform provider, see the documentation at [registry.terraform.io](https://registry.terraform.io/providers/northwood-labs/corefunc/).

### Go Package

If you are using this as a Go library, see the documentation at [pkg.go.dev](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc).

## More Information

After the provider is installed, you can run `terraform-provider-corefunc` on the CLI.

* Install with either `terraform init` or `make build`.
* The Go binary path (discovered by running `./find-go-bin.sh`) is on your `$PATH`.

This will display the following text:

```bash
terraform-provider-corefunc
```

```plain
This binary is a plugin. These are not meant to be executed directly.
Please execute the program that consumes these plugins, which will
load any plugins automatically
```

However, by passing the `--help` flag, you can see the other options available, including a description of the software.

```bash
terraform-provider-corefunc --help
```

The provider has one primary sub-command: `version`. It includes long-form version information, including the build commit hash, build date, Go version, and external dependencies.

```bash
terraform-provider-corefunc version
```

```plain
BASIC
Version:    dev
Go version: go1.22.1
Git commit: 822858fc0c80b34eebf2a5ddd1b48684414d71b3
Build date: 2024-03-11T03:09:24Z
OS/Arch:    darwin/arm64
System:     macOS on Apple Silicon
CPU Cores:  10

DEPENDENCIES
github.com/bits-and-blooms/bitset                          v1.13.0
github.com/chanced/caps                                    v1.0.2
github.com/fatih/color                                     v1.16.0
github.com/golang/protobuf                                 v1.5.4
github.com/gookit/color                                    v1.5.4
github.com/hashicorp/go-hclog                              v1.6.2
github.com/hashicorp/go-plugin                             v1.6.0
github.com/hashicorp/go-uuid                               v1.0.3
github.com/hashicorp/terraform-plugin-framework            v1.6.1
github.com/hashicorp/terraform-plugin-framework-validators v0.12.0
[...snip...]
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
