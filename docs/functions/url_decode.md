---
page_title: "url_decode function - corefunc"
subcategory: ""
description: |-
  URLDecode decodes a URL-encoded string.
  -> This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
  This also provides a 1:1 implementation that can be used with Terratest or other
  Go code.
  Maps to the corefunc.URLDecode() https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#URLDecode Go method, which can be used in Terratest https://terratest.gruntwork.io.
---

# url_decode (function)

URLDecode decodes a URL-encoded string.

-> This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
This also provides a 1:1 implementation that can be used with Terratest or other
Go code.

Maps to the [`corefunc.URLDecode()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#URLDecode) Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Signature

<!-- signature generated by tfplugindocs -->
```text
url_decode(encoded_url string) string
```

## Example Usage

```terraform
output "url_decode" {
  value = provider::corefunc::url_decode("mailto%3Aemail%3Fsubject%3Dthis%2Bis%2Bmy%2Bsubject")
}

#=> mailto:email?subject=this+is+my+subject
```

## Arguments

1. `encoded_url` (String) An encoded URL.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->