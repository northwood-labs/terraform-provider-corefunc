# Pass the prefix and the label, then limit to 64 characters

data "corefunc_str_truncate_label" "label" {
  prefix     = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
  label      = "K8S Pods Not Running Deployment Check"
  max_length = 64
}

# `value` is the read-only attribute containing the truncated label

output "value" {
  value = data.corefunc_str_truncate_label.label.value
}

# => NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…
