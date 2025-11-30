output "sha384" {
  value = provider::corefunc::hash_sha384("hello world")
}

#=> fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
