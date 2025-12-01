output "hash_sha3x512_fn" {
  value = provider::corefunc::hash_sha3x512("hello world")
}

output "hash_sha3x512_base64_fn" {
  value = provider::corefunc::hash_sha3x512_base64("hello world")
}
