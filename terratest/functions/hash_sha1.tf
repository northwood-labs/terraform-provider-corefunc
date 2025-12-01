output "hash_sha1_fn" {
  value = provider::corefunc::hash_sha1("hello world")
}

output "hash_sha1_base64_fn" {
  value = provider::corefunc::hash_sha1_base64("hello world")
}
