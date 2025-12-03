output "str_startswith_fn" {
  value = provider::corefunc::str_startswith("Hello world!", "Hello")
}

output "str_endswith_fn" {
  value = provider::corefunc::str_endswith("Hello world!", "world!")
}

output "str_contains_fn" {
  value = provider::corefunc::str_contains("Hello world!", "ello")
}
