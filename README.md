# Core Functions

[![Terraform Docs](https://img.shields.io/badge/Terraform-Docs-7B42BC?style=for-the-badge)](https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs)
[![OpenTofu Docs](https://img.shields.io/badge/OpenTofu-Docs-FEDA15?style=for-the-badge)](https://search.opentofu.org/provider/northwood-labs/corefunc/latest)
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

Core functions with identical implementations for [Terraform], [OpenTofu], [Terratest], and other software in the so-called _Terra/Fu_ ecosystem.

While some of these _can_ be implemented in HCL, some of them begin to push up against the limits of Terraform and the HCL2 configuration language. Exposing these functions as both a Go library as well as a Terraform/OpenTofu provider enables us to use the same functionality in both our Terraform/OpenTofu applies as well as while using a testing framework.

The goal of this provider is not to call any network APIs, but to provide pre-built functions in the form of _Data Sources_ or _Provider Functions_.

## Vote for features!

[View the list of issues](https://github.com/northwood-labs/terraform-provider-corefunc/issues?q=is%3Aissue+is%3Aopen+sort%3Areactions-%2B1-desc), and give a thumbs-up to the ones you'd like to see. This is how we prioritize the work.

## Compatibility testing

* We have automated testing that runs on every commit and every pull request.
* We intend for the Go libraries to work with all non-EOL versions of Go (i.e., current, current-1).
* Built using the [Terraform Plugin Framework][TPF], which speaks [Terraform Protocol v6][tfproto6].

| Testing type | Details            | Description                                                                    |
|--------------|--------------------|--------------------------------------------------------------------------------|
| integration  | Terraform 1.0–1.14 | Executes the provider with this release, pulling from `registry.terraform.io`. |
| integration  | OpenTofu 1.6–1.11  | Executes the provider with this release, pulling from `registry.opentofu.org`. |
| unit         | Go 1.24–1.25       | Tests using these versions.                                                    |
| mutation     | Go 1.24–1.25       | Tests using these versions.                                                    |
| fuzz         | Go 1.24–1.25       | Tests using these versions.                                                    |
| terratest    | Go 1.24–1.25       | Tests using these versions.                                                    |

## Usage Examples

See the `docs/` directory for user-facing documentation.

## Documentation

### Registries

* [registry.terraform.io](https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs)
* [search.opentofu.org](https://search.opentofu.org/provider/northwood-labs/corefunc/latest)
* [library.tf](https://library.tf/providers/northwood-labs/corefunc/latest)

### Go Package

If you are using this as a Go library, see the documentation at [pkg.go.dev](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc).

```bash
go get -u github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc
```

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

## Extras

We've also bundled some CLI commands that you might find as useful as we have.

### JSON → TOML

```bash
terraform-provider-corefunc json2toml file.json > file.toml
```

### TOML → JSON

```bash
terraform-provider-corefunc toml2json file.toml > file.json
```

## Ensuring quality

Northwood Labs takes testing seriously, and quality is a high priority for us. This software undergoes the following types of testing before every release.

### On every commit

* Unit testing of the Go library to ensure that the code behaves (and doesn't behave) as expected.
* Unit testing of the Go library’s documentation examples to ensure that the examples actually work.
* [Mutation testing](https://github.com/gtramontina/ooze) of the Go library to ensure that our tests are good at testing the code (as opposed to chasing code coverage).
* Acceptance testing of the provider (through both Terraform and OpenTofu) which uses [table-driven testing](https://go.dev/wiki/TableDrivenTests) to ensure that the inputs and expected outputs are correct.
* [Terratest](https://terratest.gruntwork.io) tests the provider using the same tools as one might use to test their own Terraform/OpenTofu modules.
* A _very_ strict set of linting rules using [golangci-lint](https://golangci-lint.run).

### On every release

* [Fuzzing](https://go.dev/doc/tutorial/fuzz) of the Go library which throws garbage at the code to ensure edge-cases don't slip through.
* We use the [Bash Automated Testing System](https://bats-core.readthedocs.io) together with [tfschema](https://github.com/minamijoyo/tfschema) to ensure that nothing weird or unexpected occurs with the [Terraform Plugin Protocol](https://developer.hashicorp.com/terraform/plugin/terraform-plugin-protocol) v6 between releases.

### On a regular cadence

Supply-chain testing and validation. Some of these are on every commit, while a couple run before every release.

* [Dependabot](https://github.com/dependabot) with auto-merging (by [GitHub](https://github.com)).
* [Dependency Review](https://github.com/marketplace/actions/dependency-review) (by [GitHub](https://github.com)).
* [Go Dependency Submission](https://github.com/marketplace/actions/go-dependency-submission) (by [GitHub](https://github.com)).
* [govulncheck](https://github.com/marketplace/actions/golang-govulncheck-action) (by [the Go team](https://go.dev)).
* [Harden-Runner](https://github.com/marketplace/actions/harden-runner) (by [StepSecurity](https://stepsecurity.io)).
* [OSSF Scorecard](https://github.com/marketplace/actions/ossf-scorecard-action) (by [OpenSSF](https://openssf.org)).
* [OSV Scanner](https://google.github.io/osv-scanner/github-action/) (by [Google](https://osv.dev)).
* [Trufflehog OSS](https://github.com/marketplace/actions/trufflehog-oss) (by [Truffle Security](https://trufflesecurity.com)).

[OpenTofu]: https://opentofu.org
[Terraform]: https://terraform.io
[Terratest]: https://terratest.gruntwork.io
[tfproto6]: https://developer.hashicorp.com/terraform/plugin/terraform-plugin-protocol#protocol-version-6
[TPF]: https://github.com/hashicorp/terraform-plugin-framework
