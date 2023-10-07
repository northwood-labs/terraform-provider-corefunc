data "corefunc_str_constant" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> THIS_IS_AN_EXAMPLE_ID_32
