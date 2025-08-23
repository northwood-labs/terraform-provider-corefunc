---
title: "terraform-provider-corefunc"
slug: "terraform-provider-corefunc"
description: "CLI reference for terraform-provider-corefunc"
---

## terraform-provider-corefunc

Utilities that should have been Terraform core functions.

### Synopsis

terraform-provider-corefunc

Utilities that should have been Terraform core functions.

While some of these _can_ be implemented in HCL, some of them begin to push up
against the limits of Terraform and the HCL2 configuration language. We also
perform testing using the Terratest <https://terratest.gruntwork.io> framework
on a regular basis. Exposing these functions as both a Go library as well as a
Terraform provider enables us to use the same functionality in both our
Terraform applies as well as while using a testing framework.

Since Terraform doesn't have the concept of user-defined functions, the next
step to open up the possibilities is to write a custom Terraform Provider which
has the functions built-in, using Terraform's existing support for inputs and
outputs.

**This does not add new syntax or constructs to Terraform.** Instead it uses the
existing concepts around Providers, Resources, Data Sources, Variables, and
Outputs to expose new custom-built functionality.

The goal of this provider is not to call any APIs, but to provide pre-built
functions in the form of Data Sources.

```bash
terraform-provider-corefunc [flags]
```

### Options

```text
  -d, --debug   Run with support for Go debuggers like delve. https://flwd.dk/3k055Lh
  -h, --help    help for terraform-provider-corefunc
```

### See also

* [terraform-provider-corefunc json2toml](terraform-provider-corefunc_json2toml.md) - Converts JSON to TOML.
* [terraform-provider-corefunc toml2json](terraform-provider-corefunc_toml2json.md) - Converts TOML to JSON.
* [terraform-provider-corefunc version](terraform-provider-corefunc_version.md) - Long-form version information
