output "hash_sha256_fn" {
  value = provider::corefunc::hash_sha256("hello world")
}

output "hash_sha256_base64_fn" {
  value = provider::corefunc::hash_sha256_base64("hello world")
}
