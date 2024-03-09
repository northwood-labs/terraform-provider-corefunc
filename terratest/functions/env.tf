output "env_ensure_fn" {
  description = "This returns the GOROOT environment variable."
  value       = provider::corefunc::env_ensure("GOROOT")
}
