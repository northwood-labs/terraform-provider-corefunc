# output "env_ensure_fn" {
#   description = "This returns the AWS_REGION environment variable."
#   value       = provider::corefunc::env_ensure("AWS_REGION", "^([a-z]{2})-([^-]+)-(\\d)$")
# }

output "env_ensure_fn" {
  description = "This returns the AWS_REGION environment variable."
  value       = provider::corefunc::env_ensure("AWS_VAULT", "(non)?prod$")
}
