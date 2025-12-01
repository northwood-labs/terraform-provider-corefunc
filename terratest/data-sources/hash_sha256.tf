data "corefunc_hash_sha256" "sha256" {
  input = "hello world"
}

output "hash_sha256_ds" {
  description = "The sha256 hash of the input string."
  value       = data.corefunc_hash_sha256.sha256.value
}

data "corefunc_hash_sha256_base64" "sha256_base64" {
  input = "hello world"
}

output "hash_sha256_base64_ds" {
  description = "The sha256 hash of the input string."
  value       = data.corefunc_hash_sha256_base64.sha256_base64.value
}
