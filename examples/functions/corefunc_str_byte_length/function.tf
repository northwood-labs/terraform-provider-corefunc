output "bytes1" {
  value = provider::corefunc::str_byte_length("abcde")
}

#=> 5

output "bytes2" {
  value = provider::corefunc::str_byte_length("♫")
}

#=> 1

output "bytes3" {
  value = provider::corefunc::str_byte_length("界")
}

#=> 2

output "bytes4" {
  value = provider::corefunc::str_byte_length("スター☆")
}

#=> 7
