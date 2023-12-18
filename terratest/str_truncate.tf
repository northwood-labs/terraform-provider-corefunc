# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_truncate_label
data "corefunc_str_truncate_label" "label" {
  prefix     = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
  label      = "K8S Pods Not Running Deployment Check"
  max_length = 64
}
