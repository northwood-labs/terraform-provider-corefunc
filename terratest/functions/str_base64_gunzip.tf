output "str_base64_gunzip_fn" {
  description = "This is the base64-decoded and de-gzipped output."
  value       = provider::corefunc::str_base64_gunzip("H4sIAAAAAAAA/8pIrVTISK3UUShPVS9KVSjJSFXIzc/LTk0tBgQAAP//qz+dmhoAAAA")
}
