output "hash_hmac_base64sha256" {
  value = provider::corefunc::hash_hmac_base64sha256("hello world", "secretkey")
}

#=> rmzSYF1iIxZWTR92v8DAT4nZ+vsU9Fs+GMKj4ove8p0=
