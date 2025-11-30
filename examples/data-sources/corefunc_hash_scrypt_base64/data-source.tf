data "corefunc_hash_scrypt_base64" "hash_scrypt_base64" {
  input = "hello world"
  salt  = "somesaltvalue"
}

#=> xY40ROoFBE25zbFAasFDiEs1EQHD4jyV4BDsmUyNZd8=
