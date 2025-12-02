output "str_contains" {
  value = provider::corefunc::str_contains("hello world", "ello")
}

#=> true
