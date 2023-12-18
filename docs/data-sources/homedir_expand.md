<!--
---
page_title: "corefunc_homedir_expand Data Source - corefunc"
subcategory: ""
description: |-
  Replaces the ~ character in a path with the current user's home directory.
  Maps to the [`corefunc.Homedir()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#Homedir) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).
---
-->

# corefunc_homedir_expand (Data Source)

Replaces the ~ character in a path with the current user's home directory.

Maps to the [`corefunc.Homedir()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#Homedir) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Example Usage

```terraform
data "corefunc_homedir_expand" "home" {
  path = "~/.aws/"
}
#=> /home/username/.aws/
#=> /Users/username/.aws/
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

* `path` (String) The path to expand.

### Read-Only

* `id` (Number) Not used. Required by the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework).
* `value` (String) The path with the home directory expanded.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->