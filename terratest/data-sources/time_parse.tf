data "corefunc_time_parse" "time1" {
  input = "2006-01-02T15:04:05Z"
}

output "time_parse1_ds" {
  value = data.corefunc_time_parse.time1.value
}

data "corefunc_time_parse" "time2" {
  input = "Mon, 02 Jan 2006 15:04:05 GMT"
}

output "time_parse2_ds" {
  value = data.corefunc_time_parse.time2.value
}

data "corefunc_time_parse" "time3" {
  input = "Monday, 02-Jan-2006 15:04:05 MST"
}

output "time_parse3_ds" {
  value = data.corefunc_time_parse.time3.value
}
