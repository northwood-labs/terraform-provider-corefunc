data "corefunc_str_camel" "str" {
  string = local.input_string
}

data "corefunc_str_constant" "str" {
  string = local.input_string
}

data "corefunc_str_kebab" "str" {
  string = local.input_string
}

data "corefunc_str_pascal" "str" {
  string = local.input_string
}

data "corefunc_str_snake" "str" {
  string = local.input_string
}

output "str_camel_ds" {
  description = "This is a camelCase output."
  value       = data.corefunc_str_camel.str.value
}

output "str_constant_ds" {
  description = "This is a CONSTANT_CASE output."
  value       = data.corefunc_str_constant.str.value
}

output "str_kebab_ds" {
  description = "This is a kebab-case output."
  value       = data.corefunc_str_kebab.str.value
}

output "str_pascal_ds" {
  description = "This is a PascalCase output."
  value       = data.corefunc_str_pascal.str.value
}

output "str_snake_ds" {
  description = "This is a snake_case output."
  value       = data.corefunc_str_snake.str.value
}
