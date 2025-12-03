output "time_parse1_fn" {
  value = provider::corefunc::time_parse("2006-01-02T15:04:05Z")
}

output "time_parse2_fn" {
  value = provider::corefunc::time_parse("Mon, 02 Jan 2006 15:04:05 GMT")
}

output "time_parse3_fn" {
  value = provider::corefunc::time_parse("Monday, 02-Jan-2006 15:04:05 MST")
}
