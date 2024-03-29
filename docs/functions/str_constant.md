---
page_title: "str_constant function - corefunc"
subcategory: ""
description: |-
  Converts a string to CONSTANT_CASE, removing any non-alphanumeric characters.
  Also known as SCREAMING_SNAKE_CASE.
  Maps to the corefunc.StrConstant() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#StrConstant Go method, which can be used in Terratest https://terratest.gruntwork.io.
---

# str_constant (function)

Converts a string to `CONSTANT_CASE`, removing any non-alphanumeric characters.
Also known as `SCREAMING_SNAKE_CASE`.

Maps to the [`corefunc.StrConstant()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#StrConstant) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Signature

<!-- signature generated by tfplugindocs -->
```text
str_constant(str string) string
```

## Example Usage

```terraform
output "test_with_number" {
  value = provider::corefunc::str_constant("test with number -123.456")
}

#=> TEST_WITH_NUMBER_123_456
```

## Arguments

1. `str` (String) The string to convert to `CONSTANT_CASE`.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->
