output "hash_sha3x256_fn" {
  value = provider::corefunc::hash_sha3x256("hello world")
}

output "hash_sha3x256_base64_fn" {
  value = provider::corefunc::hash_sha3x256_base64("hello world")
}
