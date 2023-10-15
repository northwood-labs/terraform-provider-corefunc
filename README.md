# Terraform Provider: Core Functions

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

* [TF Registry Documentation](https://registry.terraform.io/providers/northwood-labs/corefunc/) (not published yet; see the `docs/` directory in the interim)
* [Go Package Documentation](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc)

## Third-Party Libraries

Don't reimplement things that already work well. This project leans on the following libraries:

* [chanced/caps](https://github.com/chanced/caps) — Handles the case manipulation.

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
 Go version: go1.21.1
 Git commit: 80ac4b1062bfeb81734e505f7fd977050bc4a3e9
 Dirty repo: true
 PGO:        default.pgo
 Build date: 2023-09-24T23:50:26Z
 OS/Arch:    darwin/arm64
 System:     macOS on Apple Silicon
 CPU Cores:  10

 DEPENDENCIES
 github.com/chanced/caps                                    v1.0.1
 github.com/fatih/color                                     v1.15.0
 github.com/golang/protobuf                                 v1.5.3
 github.com/gookit/color                                    v1.5.4
 github.com/hashicorp/go-hclog                              v1.5.0
 github.com/hashicorp/go-plugin                             v1.5.1
 github.com/hashicorp/go-uuid                               v1.0.3
 github.com/hashicorp/terraform-plugin-framework            v1.4.0
 github.com/hashicorp/terraform-plugin-framework-validators v0.12.0
 github.com/hashicorp/terraform-plugin-go                   v0.19.0
 [...snip...]
```
