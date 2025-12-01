output "hash_sha512_fn" {
  value = provider::corefunc::hash_sha512("hello world")
}

output "hash_sha512_base64_fn" {
  value = provider::corefunc::hash_sha512_base64("hello world")
}
