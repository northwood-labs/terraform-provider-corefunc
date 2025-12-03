output "time_parse1" {
  value = provider::corefunc::time_parse("2006-01-02T15:04:05Z")
}

#=> 1136214245

output "time_parse2" {
  value = provider::corefunc::time_parse("Monday, 02-Jan-2006 15:04:05 MST")
}

#=> 1136214245

output "time_parse3" {
  value = provider::corefunc::time_parse("Mon, 02 Jan 2006 15:04:05 GMT")
}

#=> 1136214245
