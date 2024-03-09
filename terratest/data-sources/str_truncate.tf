data "corefunc_str_truncate_label" "label" {
  prefix     = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
  label      = "K8S Pods Not Running Deployment Check"
  max_length = 64
}

output "str_truncate_ds" {
  description = "This is a truncation output."
  value       = data.corefunc_str_truncate_label.label.value
}
