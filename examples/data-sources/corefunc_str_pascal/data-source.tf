data "corefunc_str_pascal" "test_initialisms_ip_html_eof_ascii_cpu" {
  string       = "test initialisms ip html eof ascii cpu"
  acronym_caps = true
}

#=> TestInitialismsIPHTMLEOFASCIICPU

#-----------------------------------------------------------------------

data "corefunc_str_pascal" "test_from_camel" {
  string = "TestFromCamel"
}

#=> TestFromCamel

#-----------------------------------------------------------------------

data "corefunc_str_pascal" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> ThisIsAnExampleID32

#-----------------------------------------------------------------------

data "corefunc_str_pascal" "test_with_number" {
  string = "test with number -123.456"
}

#=> TestWithNumber123456
