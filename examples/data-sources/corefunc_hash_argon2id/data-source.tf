data "corefunc_hash_argon2id" "hash_argon2id" {
  input = "hello world"
  salt  = "somesaltvalue"
}

#=> 660b8e259df3496e9c429bebd97eb5d58a88cd46851e30c1dd03b9448c066ea4
