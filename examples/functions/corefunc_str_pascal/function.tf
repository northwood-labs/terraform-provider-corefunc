output "test_with_number" {
  value = provider::corefunc::str_pascal("test with number -123.456")
}

#=> TestWithNumber123456
