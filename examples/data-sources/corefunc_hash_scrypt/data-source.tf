data "corefunc_hash_scrypt" "hash_scrypt" {
  input = "hello world"
  salt  = "somesaltvalue"
}

#=> c58e3444ea05044db9cdb1406ac143884b351101c3e23c95e010ec994c8d65df
