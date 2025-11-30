output "hash_sha3x512_base64" {
  value = provider::corefunc::hash_sha3x512_base64("hello world")
}

#=> hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg==
