output "hash_hmac_sha256_fn" {
  value = provider::corefunc::hash_hmac_sha256("hello world", "secretkey")
}

output "hash_hmac_base64sha256_fn" {
  value = provider::corefunc::hash_hmac_base64sha256("hello world", "secretkey")
}
