---
page_title: "corefunc_homedir_get Data Source - corefunc"
subcategory: ""
description: |-
  Returns the path to the home directory of the current user.
  Maps to the corefunc.Homedir() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#Homedir Go method, which can be used in Terratest https://terratest.gruntwork.io.
---

# corefunc_homedir_get (Data Source)

Returns the path to the home directory of the current user.

Maps to the [`corefunc.Homedir()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#Homedir) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Example Usage

```terraform
data "corefunc_homedir_get" "home" {}

#=> /Users/me
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

* `value` (String) The value of the home directory.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->
