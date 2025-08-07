# Read a TOML file…
data "corefunc_toml_to_json" "toml_to_json" {
  toml = file("${path.module}/hello.toml")
}

# …and write a JSON file.
resource "local_file" "json_file" {
  content  = data.corefunc_toml_to_json.toml_to_json.json
  filename = "${path.module}/hello.json"
}
