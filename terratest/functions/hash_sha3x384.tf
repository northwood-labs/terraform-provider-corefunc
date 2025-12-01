output "hash_sha3x384_fn" {
  value = provider::corefunc::hash_sha3x384("hello world")
}

output "hash_sha3x384_base64_fn" {
  value = provider::corefunc::hash_sha3x384_base64("hello world")
}
