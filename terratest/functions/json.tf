locals {
  json_str = <<EOT
{"abc": 123}
EOT
}

output "json_as_toml" {
  description = "A JSON value converted to TOML format."
  value       = provider::corefunc::json_to_toml(local.json_str)
}
