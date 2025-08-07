locals {
  toml_str = <<EOT
abc = 123
EOT
}

output "toml_as_json" {
  description = "A TOML value converted to JSON format."
  value       = provider::corefunc::toml_to_json(local.toml_str)
}
