data "corefunc_str_contains" "str" {
  input  = "Hello world!"
  substr = "ello"
}

output "str_contains_ds" {
  value = data.corefunc_str_contains.str.value
}
