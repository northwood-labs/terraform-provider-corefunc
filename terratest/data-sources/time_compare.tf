data "corefunc_time_compare" "compare1" {
  timestamp_a = "2006-01-02T15:04:05Z"
  timestamp_b = "Mon, 02 Jan 2006 15:04:05 GMT"
}

output "time_compare1_ds" {
  value = data.corefunc_time_compare.compare1.value
}

data "corefunc_time_compare" "compare2" {
  timestamp_a = "Monday, 02-Jan-2006 15:04:05 MST"
  timestamp_b = "2007-01-02T15:04:05Z"
}

output "time_compare2_ds" {
  value = data.corefunc_time_compare.compare2.value
}

data "corefunc_time_compare" "compare3" {
  timestamp_a = "2007-01-02T15:04:05Z"
  timestamp_b = "Monday, 02-Jan-2006 15:04:05 MST"
}

output "time_compare3_ds" {
  value = data.corefunc_time_compare.compare3.value
}
