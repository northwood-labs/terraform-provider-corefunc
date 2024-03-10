# AWS_VAULT="dev"

output "aws_vault" {
  # Must end with "prod" or "nonprod".
  value = provider::corefunc::env_ensure("AWS_VAULT", "(non)?prod$")
}

#=> [Error] Invalid value for "name" parameter:
#=>         environment variable AWS_VAULT does not match pattern (non)?prod$
