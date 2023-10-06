data "corefunc_str_kebab" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> this-is-an-example-id-32
