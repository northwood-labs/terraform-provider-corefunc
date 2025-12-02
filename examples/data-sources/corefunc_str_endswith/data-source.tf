data "corefunc_str_endswith" "str_endswith" {
  input  = "hello world"
  suffix = "world"
}

#=> true
