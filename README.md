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

* [ ] utils_str_snake(string)
* [ ] utils_str_camel(string)
* [ ] utils_str_kebab(string)
* [ ] utils_str_constant(string)
* [ ] utils_str_word(string)
* [ ] utils_str_leftpad(any, string, length)
* [ ] utils_str_recursive_replace(string, map[string]any)
* [ ] utils_str_resolve_markers(string)
* [ ] utils_list_first(arr)
* [ ] utils_list_last(arr)
* [ ] utils_list_push(arr, any)
* [ ] utils_list_pop(array)
* [ ] utils_list_shift(arr, any)
* [ ] utils_list_unshift(arr)
* [ ] utils_list_sort_icase(arr)
* [ ] utils_list_sort_version(arr)
* [ ] utils_map_has_key(map)
* [ ] utils_env_get(string, string)
* [ ] utils_env_ensure([]string)
* [ ] utils_dump(any)
* [ ] utils_println(string)

## Documentation

* [TF Provider Documentation](#)
* [Go Package Documentation](#)
