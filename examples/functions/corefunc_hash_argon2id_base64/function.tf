output "hash_argon2id_base64" {
  value = provider::corefunc::hash_argon2id_base64("hello world", "somesaltvalue")
}

#=> ZguOJZ3zSW6cQpvr2X611YqIzUaFHjDB3QO5RIwGbqQ=
