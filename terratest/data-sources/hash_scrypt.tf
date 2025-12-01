data "corefunc_hash_scrypt" "scrypt" {
  input = "hello world"
  salt  = "somesaltvalue"
}

output "hash_scrypt_ds" {
  description = "The scrypt hash of the input string."
  value       = data.corefunc_hash_scrypt.scrypt.value
}

data "corefunc_hash_scrypt_base64" "scrypt_base64" {
  input = "hello world"
  salt  = "somesaltvalue"
}

output "hash_scrypt_base64_ds" {
  description = "The scrypt hash of the input string."
  value       = data.corefunc_hash_scrypt_base64.scrypt_base64.value
}
