data "corefunc_hash_sha3x256" "sha3x256" {
  input = "hello world"
}

output "hash_sha3x256_ds" {
  description = "The sha3x256 hash of the input string."
  value       = data.corefunc_hash_sha3x256.sha3x256.value
}

data "corefunc_hash_sha3x256_base64" "sha3x256_base64" {
  input = "hello world"
}

output "hash_sha3x256_base64_ds" {
  description = "The sha3x256 hash of the input string."
  value       = data.corefunc_hash_sha3x256_base64.sha3x256_base64.value
}
