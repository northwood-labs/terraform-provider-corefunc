output "str_snake_fn" {
  description = "This is a snake_case output."
  value       = provider::corefunc::str_snake(local.input_string)
}
