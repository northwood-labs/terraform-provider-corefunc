output "str_startswith" {
  value = provider::corefunc::str_startswith("hello world", "hello")
}

#=> true
