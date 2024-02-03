<!--
---
page_title: "corefunc_runtime_numcpus Data Source - corefunc"
subcategory: ""
description: |-
  Returns the number of CPU cores (logical CPUs) for the current system.
  -> This counts the number of CPU cores, which may differ from the number
  of physical CPUs.
  Maps to the runtime.NumCPU() https://pkg.go.dev/runtime#NumCPU
  Go method, which can be used in Terratest https://terratest.gruntwork.io.
---
-->

# corefunc_runtime_numcpus (Data Source)

Returns the number of CPU cores (logical CPUs) for the current system.

-> This counts the number of CPU cores, which may differ from the number
of physical CPUs.

Maps to the [`runtime.NumCPU()`](https://pkg.go.dev/runtime#NumCPU)
Go method, which can be used in [Terratest](https://terratest.gruntwork.io).

## Example Usage

```terraform
data "corefunc_runtime_numcpus" "count" {}
#=> 12
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

* `value` (Number) The number of CPU cores (logical CPUs) for the current system.

<!-- Preview the provider docs with the Terraform registry provider docs preview tool: https://registry.terraform.io/tools/doc-preview -->