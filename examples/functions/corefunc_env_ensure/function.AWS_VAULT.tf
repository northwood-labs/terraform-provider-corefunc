# AWS_VAULT="dev"

output "aws_vault" {
  value = provider::corefunc::env_ensure("AWS_VAULT", "(non)?prod$") # Must end with "prod" or "nonprod".
}

#=> [Error] Invalid value for "name" parameter:
#=>         environment variable AWS_VAULT does not match pattern (non)?prod$
