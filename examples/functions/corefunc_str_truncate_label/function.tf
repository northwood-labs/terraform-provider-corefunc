output "label" {
  value = provider::corefunc::str_truncate_label(
    64,
    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
    "K8S Pods Not Running Deployment Check",
  )
}

#=> NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…
