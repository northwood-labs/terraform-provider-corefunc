# ------------------------------------------------------------------------------
# String formatting

output "str_camel" {
  description = "This is a camelCase output."
  value       = data.corefunc_str_camel.str.value
}

output "str_constant" {
  description = "This is a CONSTANT_CASE output."
  value       = data.corefunc_str_constant.str.value
}

output "str_kebab" {
  description = "This is a kebab-case output."
  value       = data.corefunc_str_kebab.str.value
}

output "str_pascal" {
  description = "This is a PascalCase output."
  value       = data.corefunc_str_pascal.str.value
}

output "str_snake" {
  description = "This is a snake_case output."
  value       = data.corefunc_str_snake.str.value
}

# ------------------------------------------------------------------------------
# String truncation

output "str_truncate" {
  description = "This is a truncation output."
  value       = data.corefunc_str_truncate_label.label.value
}

# ------------------------------------------------------------------------------
# String replacement

output "str_iterative_replace" {
  description = "This is a replacement output."
  value       = data.corefunc_str_iterative_replace.replacements.value
}

# ------------------------------------------------------------------------------
# Leftpadding

output "str_leftpad" {
  description = "This is a string leftpad output."
  value       = data.corefunc_str_leftpad.leftpad.value
}

output "int_leftpad" {
  description = "This is a numeric leftpad output."
  value       = data.corefunc_int_leftpad.leftpad.value
}

# ------------------------------------------------------------------------------
# Homedir

output "homedir_get" {
  description = "This should return the home directory of the running machine."
  value       = data.corefunc_homedir_get.home.value
}

output "homedir_expand" {
  description = "This should return the home directory of the running machine as part of the output."
  value       = data.corefunc_homedir_expand.home.value
}

# ------------------------------------------------------------------------------
# Environment

output "env_ensure" {
  description = "This should return the GOROOT environment variable."
  value       = data.corefunc_env_ensure.goroot.value
}
