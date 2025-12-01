output "hash_scrypt_fn" {
  value = provider::corefunc::hash_scrypt("hello world", "somesaltvalue")
}

output "hash_scrypt_base64_fn" {
  value = provider::corefunc::hash_scrypt_base64("hello world", "somesaltvalue")
}
