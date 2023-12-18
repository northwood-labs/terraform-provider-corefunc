# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_leftpad
data "corefunc_str_leftpad" "leftpad" {
  string    = "abc"
  pad_width = 5
  pad_char  = "."
}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/int_leftpad
data "corefunc_int_leftpad" "leftpad" {
  num       = 123
  pad_width = 5
}
