output "str" {
  value = provider::corefunc::str_base64_gunzip(
    "H4sIAAAAAAAA/8pIrVTISK3UUShPVS9KVSjJSFXIzc/LTk0tBgQAAP//qz+dmhoAAAA"
  )
}

#=> hey hey, we're the monkees
