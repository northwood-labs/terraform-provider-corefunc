data "corefunc_time_compare" "timecmp1" {
  timestamp_a = timestamp()
  timestamp_b = timeadd(self.expiration_timestamp, "-720h")
}

lifecycle {
  postcondition {
    condition     = data.corefunc_time_compare.timecmp1.value < 0
    error_message = "Certificate will expire in less than 30 days."
  }
}
