# Read a TOML file…
locals {
  toml_data = provider::corefunc::toml_to_json(
    file("${path.module}/hello.toml")
  )
}

# …and write a JSON file.
resource "local_file" "json_file" {
  content  = local.toml_data
  filename = "${path.module}/hello.json"
}
