output "test_with_number" {
  value = provider::corefunc::str_snake("test with number -123.456")
}

#=> test_with_number_123_456
