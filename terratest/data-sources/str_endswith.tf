data "corefunc_str_endswith" "str" {
  input  = "Hello world!"
  suffix = "world!"
}

output "str_endswith_ds" {
  value = data.corefunc_str_endswith.str.value
}
