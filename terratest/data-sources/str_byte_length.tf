data "corefunc_str_byte_length" "bytes1" {
  string = "abcde"
}

output "str_byte_length1_ds" {
  value = data.corefunc_str_byte_length.bytes1.value
}

data "corefunc_str_byte_length" "bytes2" {
  string = "♫"
}

output "str_byte_length2_ds" {
  value = data.corefunc_str_byte_length.bytes2.value
}

data "corefunc_str_byte_length" "bytes3" {
  string = "界"
}

output "str_byte_length3_ds" {
  value = data.corefunc_str_byte_length.bytes3.value
}

data "corefunc_str_byte_length" "bytes4" {
  string = "スター☆"
}

output "str_byte_length4_ds" {
  value = data.corefunc_str_byte_length.bytes4.value
}
