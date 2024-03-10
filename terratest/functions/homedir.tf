output "homedir_get_fn" {
  description = "This returns the home directory of the running machine."
  value       = provider::corefunc::homedir_get()
}

output "homedir_expand_fn" {
  description = "This returns the home directory of the running machine as part of the output."
  value       = provider::corefunc::homedir_expand("~/.aws/")
}
