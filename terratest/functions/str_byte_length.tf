output "str_byte_length1_fn" {
  value = provider::corefunc::str_byte_length("abcde")
}

output "str_byte_length2_fn" {
  value = provider::corefunc::str_byte_length("♫")
}

output "str_byte_length3_fn" {
  value = provider::corefunc::str_byte_length("界")
}

output "str_byte_length4_fn" {
  value = provider::corefunc::str_byte_length("スター☆")
}
