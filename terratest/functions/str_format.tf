output "str_camel_fn" {
  description = "This is a camelCase output."
  value       = provider::corefunc::str_camel(local.input_string)
}

output "str_constant_fn" {
  description = "This is a CONSTANT_CASE output."
  value       = provider::corefunc::str_constant(local.input_string)
}

output "str_kebab_fn" {
  description = "This is a kebab-case output."
  value       = provider::corefunc::str_kebab(local.input_string)
}

output "str_pascal_fn" {
  description = "This is a PascalCase output."
  value       = provider::corefunc::str_pascal(local.input_string, false)
}

output "str_snake_fn" {
  description = "This is a snake_case output."
  value       = provider::corefunc::str_snake(local.input_string)
}
