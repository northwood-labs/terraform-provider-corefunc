data "corefunc_hash_sha384" "sha384" {
  input = "hello world"
}

output "hash_sha384_ds" {
  description = "The sha384 hash of the input string."
  value       = data.corefunc_hash_sha384.sha384.value
}

data "corefunc_hash_sha384_base64" "sha384_base64" {
  input = "hello world"
}

output "hash_sha384_base64_ds" {
  description = "The sha384 hash of the input string."
  value       = data.corefunc_hash_sha384_base64.sha384_base64.value
}
