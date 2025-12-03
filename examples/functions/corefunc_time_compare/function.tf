lifecycle {
  postcondition {
    condition = provider::corefunc::time_compare(
      timestamp(),
      timeadd(self.expiration_timestamp, "-720h")
    ) < 0

    error_message = "Certificate will expire in less than 30 days."
  }
}

output "timecmp1" {
  value = provider::corefunc::time_compare("2006-01-02T15:04:05Z", "2007-01-02T15:04:05Z")
}
