data "corefunc_hash_sha1" "sha1" {
  input = "hello world"
}

output "hash_sha1_ds" {
  description = "The sha1 hash of the input string."
  value       = data.corefunc_hash_sha1.sha1.value
}

data "corefunc_hash_sha1_base64" "sha1_base64" {
  input = "hello world"
}

output "hash_sha1_base64_ds" {
  description = "The sha1 hash of the input string."
  value       = data.corefunc_hash_sha1_base64.sha1_base64.value
}
