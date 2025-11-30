output "sha1_base64" {
  value = provider::corefunc::hash_sha1_base64("hello world")
}

#=> Kq5sNclPz7QV2+lfQIuc6R7oRu0=
