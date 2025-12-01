data "corefunc_hash_sha512" "sha512" {
  input = "hello world"
}

output "hash_sha512_ds" {
  description = "The sha512 hash of the input string."
  value       = data.corefunc_hash_sha512.sha512.value
}

data "corefunc_hash_sha512_base64" "sha512_base64" {
  input = "hello world"
}

output "hash_sha512_base64_ds" {
  description = "The sha512 hash of the input string."
  value       = data.corefunc_hash_sha512_base64.sha512_base64.value
}
