data "corefunc_json_to_toml" "json_to_toml" {
  json = <<EOT
{"abc": 123}
EOT
}

output "json_as_toml" {
  description = "A JSON value converted to TOML format."
  value       = data.corefunc_json_to_toml.json_to_toml.toml
}
