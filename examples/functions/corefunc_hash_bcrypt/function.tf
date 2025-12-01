output "hash_bcrypt" {
  value = provider::corefunc::hash_bcrypt("hello world", 10)
}

#=> $2a$10$SOMESALTVALUEu1pYx0pG7xO6dXJv1bG8y9qzE6jFzW8u9K5eV6Zy
