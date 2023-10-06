data "corefunc_str_camel" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> thisIsAnExampleID32
