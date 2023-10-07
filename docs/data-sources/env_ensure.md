<!--
---
page_title: "corefunc_env_ensure Data Source - corefunc"
subcategory: ""
description: |-
  Ensures that a given environment variable is set to a non-empty value.
  If the environment variable is unset or if it is set to an empty string,
  it will trigger a Terraform-level error.
  -> Not every Terraform provider checks to ensure that the environment variables it
  requires are properly set before performing work, leading to late-stage errors.
  This will force an error to occur early in the execution if the environment
  variable is not set, or if its value doesn't match the expected patttern.
  Maps to the corefunc.EnvEnsure() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#EnvEnsure Go method, which can be used in
  Terratest https://terratest.gruntwork.io.
---
-->

# corefunc_env_ensure (Data Source)

Ensures that a given environment variable is set to a non-empty value.
If the environment variable is unset or if it is set to an empty string,
it will trigger a Terraform-level error.

-> Not every Terraform provider checks to ensure that the environment variables it
requires are properly set before performing work, leading to late-stage errors.
This will force an error to occur early in the execution if the environment
variable is not set, or if its value doesn't match the expected patttern.

Maps to the [`corefunc.EnvEnsure()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#EnvEnsure) Go method, which can be used in
[Terratest](https://terratest.gruntwork.io).

## Example Usage

### `AWS_DEFAULT_REGION` is defined (no error)

```terraform
# AWS_DEFAULT_REGION="us-east-1"

data "corefunc_env_ensure" "aws_default_region" {
  name = "AWS_DEFAULT_REGION"
}

# `aws_default_region_value` is the read-only attribute containing the
# value of the environment variable
output "aws_default_region_value" {
  value = data.corefunc_env_ensure.aws_default_region.value
}

#=> us-east-1
```

### `AWS_PAGER` is set to an empty string (error)

```terraform
# AWS_PAGER=""

data "corefunc_env_ensure" "aws_pager" {
  name = "AWS_PAGER"
}

output "aws_pager_value" {
  value = data.corefunc_env_ensure.aws_pager.value
}

#=> [Error] Problem with Environment Variable: environment variable
#=>         AWS_PAGER is not defined
```

### `AWS_VAULT` is defined, but does not match pattern (error)

```terraform
# AWS_VAULT="dev"

data "corefunc_env_ensure" "aws_vault" {
  name    = "AWS_VAULT"
  pattern = "(non)?prod$" # Must end with "prod" or "nonprod".
}

output "aws_vault_value" {
  value = data.corefunc_env_ensure.aws_vault.value
}

#=> [Error] Problem with Environment Variable: environment variable
#=>         AWS_VAULT does not match pattern (non)?prod$
```

### Example plan output

```bash
AWS_DEFAULT_REGION="us-east-1" AWS_VAULT="dev" terraform plan
```

```plain
data.corefunc_env_ensure.aws_vault: Reading...
data.corefunc_env_ensure.aws_pager: Reading...
data.corefunc_env_ensure.aws_default_region: Reading...
data.corefunc_env_ensure.aws_default_region: Read complete after 0s [name=AWS_DEFAULT_REGION]

Changes to Outputs:
  + aws_default_region_value = "us-east-1"

You can apply this plan to save these new output values to the Terraform state, without changing any real infrastructure.
╷
│ Error: Problem with Environment Variable
│
│   with data.corefunc_env_ensure.aws_pager,
│   on data-source-AWS_PAGER.tf line 4, in data "corefunc_env_ensure" "aws_pager":
│    4: data "corefunc_env_ensure" "aws_pager" {
│
│ environment variable AWS_PAGER is not defined
╵
╷
│ Error: Problem with Environment Variable
│
│   with data.corefunc_env_ensure.aws_vault,
│   on data-source-AWS_VAULT.tf line 4, in data "corefunc_env_ensure" "aws_vault":
│    4: data "corefunc_env_ensure" "aws_vault" {
│
│ environment variable AWS_VAULT does not match pattern (non)?prod$
╵
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

* `name` (String) The name of the environment variable to check.

### Optional

* `pattern` (String) A valid Go ([re2](https://github.com/google/re2/wiki/Syntax)) regular expression pattern.

### Read-Only

* `id` (Number) Not used. Required by the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework).
* `value` (String) The value of the environment variable, if it exists.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->