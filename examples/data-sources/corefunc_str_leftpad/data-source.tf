data "corefunc_str_leftpad" "leftpad" {
  string    = "abc"
  pad_width = 5
  pad_char  = "."
}

#=> ..abc
