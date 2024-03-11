output "str_truncate_fn" {
  description = "This is a truncation output."
  value = provider::corefunc::str_truncate_label(
    64,
    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
    "K8S Pods Not Running Deployment Check",
  )
}
