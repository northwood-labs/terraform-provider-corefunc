data "corefunc_hash_sha3x384" "sha3x384" {
  input = "hello world"
}

output "hash_sha3x384_ds" {
  description = "The sha3x384 hash of the input string."
  value       = data.corefunc_hash_sha3x384.sha3x384.value
}

data "corefunc_hash_sha3x384_base64" "sha3x384_base64" {
  input = "hello world"
}

output "hash_sha3x384_base64_ds" {
  description = "The sha3x384 hash of the input string."
  value       = data.corefunc_hash_sha3x384_base64.sha3x384_base64.value
}
