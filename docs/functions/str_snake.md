<!--
---
page_title: "str_snake function - corefunc"
subcategory: ""
description: |-
  Converts a string to snake_case, removing any non-alphanumeric characters.
  Maps to the corefunc.StrSnake() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#StrSnake Go method, which can be used in Terratest https://terratest.gruntwork.io.
---
-->

# str_snake (function)

Converts a string to `snake_case`, removing any non-alphanumeric characters.

Maps to the [`corefunc.StrSnake()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#StrSnake) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Signature

<!-- signature generated by tfplugindocs -->
```text
str_snake(str string) string
```

## Example Usage

```terraform
output "test_with_number" {
  value = provider::corefunc::str_snake("test with number -123.456")
}

#=> test_with_number_123_456
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `str` (String) The string to convert to `snake_case`.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->