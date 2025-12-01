output "hash_sha384_fn" {
  value = provider::corefunc::hash_sha384("hello world")
}

output "hash_sha384_base64_fn" {
  value = provider::corefunc::hash_sha384_base64("hello world")
}
