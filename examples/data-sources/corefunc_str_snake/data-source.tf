data "corefunc_str_snake" "v322" {
  string = "v3.2.2"
}

#=> v3_2_2

#-----------------------------------------------------------------------

data "corefunc_str_snake" "test_from_camel" {
  string = "TestFromCamel"
}

#=> test_from_camel

#-----------------------------------------------------------------------

data "corefunc_str_snake" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> this_is_an_example_id_32

#-----------------------------------------------------------------------

data "corefunc_str_snake" "test_with_number" {
  string = "test with number -123.456"
}

#=> test_with_number_123_456
