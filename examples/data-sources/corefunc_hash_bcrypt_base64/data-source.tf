data "corefunc_hash_bcrypt_base64" "hash_bcrypt_base64" {
  input = "hello world"
  cost  = 10
}

#=> $2a$10$U09NRVNBTFRWTkFNRXUxYll4MHBHN3hPNmRYSnYxYkc4eTlxekU2akZ6Vzh1OUs1ZVY2Wnk=
