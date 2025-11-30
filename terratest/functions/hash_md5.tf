output "hash_md5_fn" {
  value = provider::corefunc::hash_md5("hello world")
}

output "hash_md5_base64_fn" {
  value = provider::corefunc::hash_md5_base64("hello world")
}
