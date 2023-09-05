# Terraform Provider: Core Functions

> [!NOTE]
> We intend to support OpenTF in addition to Terraform.

Utilities that should have been Terraform _core functions_.

While some of these _can_ be implemented in HCL, some of them begin to push up against the limits of Terraform and the HCL2 configuration language. We also perform testing using the [Terratest](https://terratest.gruntwork.io) framework on a regular basis. Exposing these functions as both a Go library as well as a Terraform provider enables us to use the same functionality in both our Terraform applies as well as while using a testing framework.

Since Terraform doesn't have the concept of user-defined functions, the next step to open up the possibilities is to write a custom Terraform Provider which has the functions built-in, using Terraform's existing support for inputs and outputs.

**This does not add new syntax or constructs to Terraform.** Instead it uses the _existing_ concepts around Providers, Resources, Data Sources, Variables, and Outputs to expose new custom-built functionality.

The goal of this provider is not to call any APIs, but to provide pre-built functions in the form of _Data Sources_.

## Usage Examples

See the `docs/` directory for user-facing documentation.

## Planned Functions

* [ ] corefunc_dump(any)
* [X] corefunc_env_ensure([]string)
* [ ] corefunc_list_first(arr)
* [ ] corefunc_list_last(arr)
* [ ] corefunc_list_pop(array)
* [ ] corefunc_list_push(arr, any)
* [ ] corefunc_list_shift(arr, any)
* [ ] corefunc_list_sort_icase(arr)
* [ ] corefunc_list_sort_version(arr)
* [ ] corefunc_list_unshift(arr)
* [ ] corefunc_println(string)
* [ ] corefunc_str_camel(string)
* [ ] corefunc_str_constant(string)
* [ ] corefunc_str_kebab(string)
* [ ] corefunc_str_leftpad(any, string, length)
* [ ] corefunc_str_recursive_replace(string, map[string]any)
* [ ] corefunc_str_resolve_markers(string)
* [ ] corefunc_str_snake(string)
* [X] corefunc_str_truncate_label(int, string, string)
* [ ] corefunc_str_word(string)

## Documentation

* [TF Registry Documentation](https://registry.terraform.io/providers/northwood-labs/corefunc/) (not published yet; see the `docs/` directory in the interim)
* [Go Package Documentation](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc)
