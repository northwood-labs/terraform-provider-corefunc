data "corefunc_str_leftpad" "leftpad" {
  string    = "abc"
  pad_width = 5
  pad_char  = "."
}

data "corefunc_int_leftpad" "leftpad" {
  num       = 123
  pad_width = 5
}

output "str_leftpad_ds" {
  description = "This is a string leftpad output."
  value       = data.corefunc_str_leftpad.leftpad.value
}

output "int_leftpad_ds" {
  description = "This is a numeric leftpad output."
  value       = data.corefunc_int_leftpad.leftpad.value
}
