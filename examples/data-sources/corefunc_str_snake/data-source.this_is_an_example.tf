data "corefunc_str_snake" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> this_is_an_example_id_32
