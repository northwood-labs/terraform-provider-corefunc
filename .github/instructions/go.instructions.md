---
applyTo: "*.go"
---

Files with the `.go` extension are associated with the Go programming language. Optimize code for Go 1.23.

Any code that is generated should focus on performance and reducing memory allocations. Avoid writing code with errors.

Code which exists in the `./corefuncprovider` directory of this repository leverages the Terraform Plugin Framework, whose original source code exists at <https://github.com/hashicorp/terraform-plugin-framework> and whose documentation exists at <https://developer.hashicorp.com/terraform/plugin/framework>. Please learn from these sources in order to better understand how to write accurate and useful code.
