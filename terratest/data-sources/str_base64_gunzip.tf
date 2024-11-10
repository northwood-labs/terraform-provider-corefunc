data "corefunc_str_base64_gunzip" "str" {
  gzipped_base64 = "H4sIAAAAAAAA/8pIrVTISK3UUShPVS9KVSjJSFXIzc/LTk0tBgQAAP//qz+dmhoAAAA"
}

output "str_base64_gunzip_ds" {
  description = "This is the base64-decoded and de-gzipped output."
  value       = data.corefunc_str_base64_gunzip.str.value
}
