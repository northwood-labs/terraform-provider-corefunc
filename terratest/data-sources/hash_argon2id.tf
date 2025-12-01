data "corefunc_hash_argon2id" "argon2id" {
  input = "hello world"
  salt  = "somesaltvalue"
}

output "hash_argon2id_ds" {
  description = "The argon2id hash of the input string."
  value       = data.corefunc_hash_argon2id.argon2id.value
}

data "corefunc_hash_argon2id_base64" "argon2id_base64" {
  input = "hello world"
  salt  = "somesaltvalue"
}

output "hash_argon2id_base64_ds" {
  description = "The argon2id hash of the input string."
  value       = data.corefunc_hash_argon2id_base64.argon2id_base64.value
}
