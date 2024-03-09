data "corefunc_env_ensure" "goroot" {
  name = "GOROOT"
}

output "env_ensure_ds" {
  description = "This returns the GOROOT environment variable."
  value       = data.corefunc_env_ensure.goroot.value
}
