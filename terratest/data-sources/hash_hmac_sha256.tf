data "corefunc_hash_hmac_sha256" "hmac_sha256" {
  input = "hello world"
  key   = "secretkey"
}

output "hash_hmac_sha256_ds" {
  description = "The hmac_sha256 hash of the input string."
  value       = data.corefunc_hash_hmac_sha256.hmac_sha256.value
}

data "corefunc_hash_hmac_base64sha256" "hmac_base64sha256" {
  input = "hello world"
  key   = "secretkey"
}

output "hash_hmac_base64sha256_ds" {
  description = "The hmac_sha256 hash of the input string."
  value       = data.corefunc_hash_hmac_base64sha256.hmac_base64sha256.value
}
