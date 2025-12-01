data "corefunc_hash_hmac_sha256" "hash_hmac_sha256" {
  input = "hello world"
  salt  = "secretkey"
}

#=> ae6cd2605d622316564d1f76bfc0c04f89d9fafb14f45b3e18c2a3e28bdef29d
