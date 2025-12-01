output "hash_argon2id_fn" {
  value = provider::corefunc::hash_argon2id("hello world", "somesaltvalue")
}

output "hash_argon2id_base64_fn" {
  value = provider::corefunc::hash_argon2id_base64("hello world", "somesaltvalue")
}
