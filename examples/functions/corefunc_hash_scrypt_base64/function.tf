output "hash_scrypt_base64" {
  value = provider::corefunc::hash_scrypt_base64("hello world", "somesaltvalue")
}

#=> xY40ROoFBE25zbFAasFDiEs1EQHD4jyV4BDsmUyNZd8=
