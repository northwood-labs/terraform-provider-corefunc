data "corefunc_str_pascal" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> ThisIsAnExampleID32
