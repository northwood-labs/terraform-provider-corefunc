data "corefunc_str_byte_length" "bytes1" {
  string = "abcde"
}

#=> 5

data "corefunc_str_byte_length" "bytes2" {
  string = "♫"
}

#=> 1

data "corefunc_str_byte_length" "bytes3" {
  string = "界"
}

#=> 2

data "corefunc_str_byte_length" "bytes4" {
  string = "スター☆"
}

#=> 7
