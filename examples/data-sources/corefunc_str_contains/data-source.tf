data "corefunc_str_contains" "str_contains" {
  input  = "hello world"
  substr = "ello"
}

#=> true
