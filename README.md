# Terraform Provider: Core Functions

[![Terraform Docs](https://img.shields.io/badge/Terraform-Docs-7B42BC?style=for-the-badge)](https://registry.terraform.io/providers/northwood-labs/corefunc/)
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

> [!NOTE]
> We intend to support OpenTofu (née OpenTF) in addition to Terraform.

Utilities that should have been Terraform _core functions_.

While some of these _can_ be implemented in HCL, some of them begin to push up against the limits of Terraform and the HCL2 configuration language. We also perform testing using the [Terratest](https://terratest.gruntwork.io) framework on a regular basis. Exposing these functions as both a Go library as well as a Terraform provider enables us to use the same functionality in both our Terraform applies as well as while using a testing framework.

Since Terraform doesn't have the concept of user-defined functions, the next step to open up the possibilities is to write a custom Terraform Provider which has the functions built-in, using Terraform's existing support for inputs and outputs.

**This does not add new syntax or constructs to Terraform.** Instead it uses the _existing_ concepts around Providers, Resources, Data Sources, Variables, and Outputs to expose new custom-built functionality.

The goal of this provider is not to call any APIs, but to provide pre-built functions in the form of _Data Sources_.

## Usage Examples

See the `docs/` directory for user-facing documentation.

## Documentation

### Terraform Registry

If you are using this as a Terraform provider, see the documentation at [registry.terraform.io](https://registry.terraform.io/providers/northwood-labs/corefunc/).

### Go Package

If you are using this as a Go library, see the documentation at [pkg.go.dev](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc).

## Third-Party Libraries

Don't reimplement things that already work well. This project leans on the following libraries:

* [chanced/caps](https://github.com/chanced/caps) — Handles the case manipulation.
* [mitchellh/go-homedir](https://github.com/mitchellh/go-homedir) — Handles looking-up a user's home directory without CGO.

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
Go version: go1.21.4
Git commit: c7b6411da19a792762d4037d2e17f4877d1f4e2b
Dirty repo: true
Build date: 2023-11-21T08:10:30Z
OS/Arch:    darwin/arm64
System:     macOS on Apple Silicon
CPU Cores:  12

DEPENDENCIES
github.com/chanced/caps                                    v1.0.1
github.com/fatih/color                                     v1.14.1
github.com/golang/protobuf                                 v1.5.3
github.com/gookit/color                                    v1.5.4
github.com/hashicorp/go-hclog                              v1.5.0
github.com/hashicorp/go-plugin                             v1.5.2
github.com/hashicorp/go-uuid                               v1.0.3
github.com/hashicorp/terraform-plugin-framework            v1.4.2
github.com/hashicorp/terraform-plugin-framework-validators v0.12.0
github.com/hashicorp/terraform-plugin-go                   v0.19.1
[...snip...]
```
