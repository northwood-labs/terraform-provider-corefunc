<!--
---
page_title: "corefunc_int_leftpad Data Source - corefunc"
subcategory: ""
description: |-
  Converts an integer to a string, and then pads it with zeroes on the left.
  Maps to the corefunc.IntLeftPad() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#IntLeftPad Go method, which can be used in Terratest https://terratest.gruntwork.io.
---
-->

# corefunc_int_leftpad (Data Source)

Converts an integer to a string, and then pads it with zeroes on the left.

Maps to the [`corefunc.IntLeftPad()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#IntLeftPad) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Example Usage

```terraform
data "corefunc_int_leftpad" "leftpad" {
  num       = 123
  pad_width = 5
}
#=> "00123"
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

* `num` (Number) The integer to pad with zeroes.
* `pad_width` (Number) The max number of zeroes to pad the integer with.

### Read-Only

* `id` (Number) Not used. Required by the [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework).
* `value` (String) The value of the string.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->