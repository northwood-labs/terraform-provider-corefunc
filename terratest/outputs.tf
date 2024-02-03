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
  description = "This returns the home directory of the running machine."
  value       = data.corefunc_homedir_get.home.value
}

output "homedir_expand" {
  description = "This returns the home directory of the running machine as part of the output."
  value       = data.corefunc_homedir_expand.home.value
}

# ------------------------------------------------------------------------------
# Environment

output "env_ensure" {
  description = "This returns the GOROOT environment variable."
  value       = data.corefunc_env_ensure.goroot.value
}

# ------------------------------------------------------------------------------
# Runtime

output "runtime_cpuarch" {
  description = "This returns the CPU architecture of the provider."
  value       = data.corefunc_runtime_cpuarch.arch.value
}

output "runtime_goroot" {
  description = "This returns the GOROOT directory, if it exists."
  value       = data.corefunc_runtime_goroot.goroot.value
}

output "runtime_numcpus" {
  description = "This returns the number of logical CPU cores."
  value       = data.corefunc_runtime_numcpus.count.value
}

output "runtime_os" {
  description = "This returns the operating system of the provider."
  value       = data.corefunc_runtime_os.os.value
}
