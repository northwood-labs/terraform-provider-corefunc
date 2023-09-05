# AWS_VAULT="dev"

data "corefunc_env_ensure" "aws_vault" {
  name    = "AWS_VAULT"
  pattern = "(non)?prod$" # Must end with "prod" or "nonprod".
}

output "aws_vault_value" {
  value = data.corefunc_env_ensure.aws_vault.value
}

#=> [Error] Problem with Environment Variable: environment variable
#=>         AWS_VAULT does not match pattern (non)?prod$
