output "md5_base64" {
  value = provider::corefunc::hash_md5_base64("hello world")
}

#=> XrY7u+Ae7tCTyyK7j1rNww==
