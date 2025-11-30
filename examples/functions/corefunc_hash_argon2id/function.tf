output "hash_argon2id" {
  value = provider::corefunc::hash_argon2id("hello world", "somesaltvalue")
}

#=> 660b8e259df3496e9c429bebd97eb5d58a88cd46851e30c1dd03b9448c066ea4
