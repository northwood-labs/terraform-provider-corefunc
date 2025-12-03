data "corefunc_str_startswith" "str" {
  input  = "Hello world!"
  prefix = "Hello"
}

output "str_startswith_ds" {
  value = data.corefunc_str_startswith.str.value
}
