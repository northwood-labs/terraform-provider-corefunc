data "corefunc_hash_sha3x512" "sha3x512" {
  input = "hello world"
}

output "hash_sha3x512_ds" {
  description = "The sha3x512 hash of the input string."
  value       = data.corefunc_hash_sha3x512.sha3x512.value
}

data "corefunc_hash_sha3x512_base64" "sha3x512_base64" {
  input = "hello world"
}

output "hash_sha3x512_base64_ds" {
  description = "The sha3x512 hash of the input string."
  value       = data.corefunc_hash_sha3x512_base64.sha3x512_base64.value
}
