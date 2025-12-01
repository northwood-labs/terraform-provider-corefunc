output "hash_hmac_sha256" {
  value = provider::corefunc::hash_hmac_sha256("hello world", "secretkey")
}

#=> ae6cd2605d622316564d1f76bfc0c04f89d9fafb14f45b3e18c2a3e28bdef29d
