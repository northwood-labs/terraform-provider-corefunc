# Read a JSON file…
locals {
  json_data = provider::corefunc::json_to_toml(
    file("${path.module}/hello.json")
  )
}

# …and write a TOML file.
resource "local_file" "toml_file" {
  content  = local.json_data
  filename = "${path.module}/hello.toml"
}
