data "corefunc_str_camel" "test_initialisms_ip_html_eof_ascii_cpu" {
  string = "test initialisms ip html eof ascii cpu"
}

#=> testInitialismsIPHTMLEOFASCIICPU

#-----------------------------------------------------------------------

data "corefunc_str_camel" "test_from_camel" {
  string = "TestFromCamel"
}

#=> testFromCamel

#-----------------------------------------------------------------------

data "corefunc_str_camel" "this_is_an_example" {
  string = "This is [an] {example}$${id32}."
}

#=> thisIsAnExampleID32

#-----------------------------------------------------------------------

data "corefunc_str_camel" "test_with_number" {
  string = "test with number -123.456"
}

#=> testWithNumber123456
