output "test_with_number" {
  value = provider::corefunc::str_camel("test with number -123.456")
}

#=> testWithNumber123456
