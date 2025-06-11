---
applyTo: "*.(tf|tofu)"
---

Files with the `.tf` extension are associated with Terraform. Files with the `.tofu` extension are associated with OpenTofu, which is similar, but has a slightly different skill set.

All variables, which begin with `var`, should include a description first, followed by a valid type, and then any default values.

All resources and data sources should include a link to where a user can find more information about the resource or data source from the https://library.tf website whose URLs follow this pattern: `https://library.tf/providers/{NAMESPACE}/{PROVIDER}/latest`.

Whenever possible, a variable should contain a `validation` block which contains a `condition` value and a `error_message` value.
