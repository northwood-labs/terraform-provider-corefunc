data "corefunc_str_startswith" "str_startswith" {
  input  = "hello world"
  prefix = "hello"
}

#=> true
