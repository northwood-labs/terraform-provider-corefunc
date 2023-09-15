data "corefunc_str_snake" "v322" {
  string = "v3.2.2"
}

output "value" {
  value = data.corefunc_str_snake.v322.value
}

#=> v3_2_2

#-----------------------------------------------------------------------

data "corefunc_str_snake" "test_from_camel" {
  string = "TestFromCamel"
}

output "value_from_camel" {
  value = data.corefunc_str_snake.test_from_camel.value
}

#=> test_from_camel

#-----------------------------------------------------------------------

data "corefunc_str_snake" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

output "value_is_an_example" {
  value = data.corefunc_str_snake.this_is_an_example.value
}

#=> this_is_an_example_id_32

#-----------------------------------------------------------------------

data "corefunc_str_snake" "test_with_number" {
  string = "test with number -123.456"
}

output "value_with_number" {
  value = data.corefunc_str_snake.test_with_number.value
}

#=> test_with_number_123_456
