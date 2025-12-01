data "corefunc_hash_hmac_base64sha256" "hash_hmac_base64sha256" {
  input = "hello world"
  key   = "secretkey"
}

#=> rmzSYF1iIxZWTR92v8DAT4nZ+vsU9Fs+GMKj4ove8p0=
