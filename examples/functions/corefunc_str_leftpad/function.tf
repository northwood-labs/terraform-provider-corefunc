output "leftpad" {
  value = provider::corefunc::str_leftpad("foo", 5, ".")
}

#=> ..foo
