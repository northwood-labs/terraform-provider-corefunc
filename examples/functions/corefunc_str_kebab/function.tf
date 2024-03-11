output "test_with_number" {
  value = provider::corefunc::str_kebab("test with number -123.456")
}

#=> test-with-number-123-456
