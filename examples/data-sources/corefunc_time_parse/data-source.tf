data "corefunc_time_parse" "time1" {
  input = "2006-01-02T15:04:05Z"
}

#=> 1136214245

data "corefunc_time_parse" "time2" {
  input = "Monday, 02-Jan-2006 15:04:05 MST"
}

#=> 1136214245

data "corefunc_time_parse" "time3" {
  input = "Mon, 02 Jan 2006 15:04:05 GMT"
}

#=> 1136214245
