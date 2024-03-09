data "corefunc_homedir_get" "home" {}

data "corefunc_homedir_expand" "home" {
  path = "~/.aws/"
}

output "homedir_get_ds" {
  description = "This returns the home directory of the running machine."
  value       = data.corefunc_homedir_get.home.value
}

output "homedir_expand_ds" {
  description = "This returns the home directory of the running machine as part of the output."
  value       = data.corefunc_homedir_expand.home.value
}
