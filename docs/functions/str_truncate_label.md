---
page_title: "str_truncate_label function - corefunc"
subcategory: ""
description: |-
  Supports prepending a prefix to a label, while truncating them
  to meet the maximum length constraints. Useful when grouping labels with a
  kind of prefix. Both the prefix and the label will be truncated if necessary.
  Uses a “balancing” algorithm between the prefix and the label, so that each
  section is truncated as a factor of how much space it takes up in the merged
  string.
  -> The motivation for this is in working with monitoring systems such
  as New Relic and Datadog where there are hundreds of applications in a
  monitoring “prod” account, and also hundreds of applications in a monitoring
  “nonprod” account. This allows us to group lists of monitors together using a
  shared prefix, but also truncate them appropriately to fit length
  constraints for names.
  Maps to the corefunc.TruncateLabel() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#TruncateLabel Go method, which can be used in
  Terratest https://terratest.gruntwork.io.
---

# str_truncate_label (function)

Supports prepending a prefix to a label, while truncating them
to meet the maximum length constraints. Useful when grouping labels with a
kind of prefix. Both the prefix and the label will be truncated if necessary.

Uses a “balancing” algorithm between the prefix and the label, so that each
section is truncated as a factor of how much space it takes up in the merged
string.

-> The motivation for this is in working with monitoring systems such
as New Relic and Datadog where there are hundreds of applications in a
monitoring “prod” account, and also hundreds of applications in a monitoring
“nonprod” account. This allows us to group lists of monitors together using a
shared prefix, but also truncate them appropriately to fit length
constraints for names.

Maps to the [`corefunc.TruncateLabel()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#TruncateLabel) Go method, which can be used in
[Terratest](https://terratest.gruntwork.io).

## Signature

<!-- signature generated by tfplugindocs -->
```text
str_truncate_label(max_length number, prefix string, label string) string
```

## Example Usage

```terraform
output "label" {
  value = provider::corefunc::str_truncate_label(
    64,
    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
    "K8S Pods Not Running Deployment Check",
  )
}

#=> NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…
```

## Arguments

1. `max_length` (Number) The maximum allowed length of the combined label. Minimum value is `1`.
1. `prefix` (String) The prefix to prepend to the label.
1. `label` (String) The label itself.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->