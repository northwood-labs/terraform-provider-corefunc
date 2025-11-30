output "md5" {
  value = provider::corefunc::hash_md5("hello world")
}

#=> 5eb63bbbe01eeed093cb22bb8f5acdc3
