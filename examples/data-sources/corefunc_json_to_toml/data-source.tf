# Read a JSON file…
data "corefunc_json_to_toml" "json_to_toml" {
  json = file("${path.module}/hello.json")
}

# …and write a TOML file.
resource "local_file" "toml_file" {
  content  = data.corefunc_json_to_toml.json_to_toml.toml
  filename = "${path.module}/hello.toml"
}
