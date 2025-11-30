data "corefunc_hash_argon2id_base64" "hash_argon2id_base64" {
  input = "hello world"
  salt  = "somesaltvalue"
}

#=> ZguOJZ3zSW6cQpvr2X611YqIzUaFHjDB3QO5RIwGbqQ=
