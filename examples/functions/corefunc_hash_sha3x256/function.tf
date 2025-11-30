output "hash_sha3x256" {
  value = provider::corefunc::hash_sha3x256("hello world")
}

#=> 644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938
