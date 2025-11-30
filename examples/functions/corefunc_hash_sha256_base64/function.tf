output "sha256_base64" {
  value = provider::corefunc::hash_sha256_base64("hello world")
}

#=> uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=
