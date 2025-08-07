data "corefunc_toml_to_json" "toml_to_json" {
  toml = <<EOT
abc = 123
EOT
}

output "toml_as_json" {
  description = "A TOML value converted to JSON format."
  value       = data.corefunc_toml_to_json.toml_to_json.json
}
