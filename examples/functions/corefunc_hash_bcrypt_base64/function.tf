output "hash_bcrypt_base64" {
  value = provider::corefunc::hash_bcrypt_base64("hello world", 10)
}

#=> $2a$10$U09NRVNBTFRWTkFNRXUxYll4MHBHN3hPNmRYSnYxYkc4eTlxekU2akZ6Vzh1OUs1ZVY2Wnk=
