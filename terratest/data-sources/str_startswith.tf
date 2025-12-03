data "corefunc_str_startswith" "str" {
  input  = "Hello world!"
  prefix = "Hello"
}

output "str_startswith_ds" {
  value = data.corefunc_str_startswith.str.value
}

data "corefunc_str_endswith" "str" {
  input  = "Hello world!"
  suffix = "world!"
}

output "str_endswith_ds" {
  value = data.corefunc_str_endswith.str.value
}

data "corefunc_str_contains" "str" {
  input  = "Hello world!"
  substr = "ello"
}

output "str_contains_ds" {
  value = data.corefunc_str_contains.str.value
}
