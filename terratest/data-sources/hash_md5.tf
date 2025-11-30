data "corefunc_hash_md5" "md5" {
  input = "hello world"
}

output "hash_md5_ds" {
  description = "The MD5 hash of the input string."
  value       = data.corefunc_hash_md5.md5.value
}

data "corefunc_hash_md5_base64" "md5_base64" {
  input = "hello world"
}

output "hash_md5_base64_ds" {
  description = "The MD5 hash of the input string."
  value       = data.corefunc_hash_md5_base64.md5_base64.value
}
