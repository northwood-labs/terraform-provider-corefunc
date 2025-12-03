output "time_compare1_fn" {
  value = provider::corefunc::time_compare("2006-01-02T15:04:05Z", "Mon, 02 Jan 2006 15:04:05 GMT")
}

output "time_compare2_fn" {
  value = provider::corefunc::time_compare("Monday, 02-Jan-2006 15:04:05 MST", "2007-01-02T15:04:05Z")
}

output "time_compare3_fn" {
  value = provider::corefunc::time_compare("2007-01-02T15:04:05Z", "Monday, 02-Jan-2006 15:04:05 MST")
}
