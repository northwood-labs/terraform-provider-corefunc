output "str_endswith" {
  value = provider::corefunc::str_endswith("hello world", "world")
}

#=> true
