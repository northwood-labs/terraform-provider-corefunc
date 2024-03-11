output "str_leftpad_fn" {
  description = "This is a string leftpad output."
  value       = provider::corefunc::str_leftpad("abc", 5, ".")
}

output "int_leftpad_fn" {
  description = "This is a numeric leftpad output."
  value       = provider::corefunc::int_leftpad(123, 5)
}
