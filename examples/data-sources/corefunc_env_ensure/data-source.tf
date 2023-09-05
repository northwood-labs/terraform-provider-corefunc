#-------------------------------------------------------------------------
# AWS_DEFAULT_REGION="us-east-1"

data "corefunc_env_ensure" "aws_default_region" {
  name = "AWS_DEFAULT_REGION"
}

# `aws_default_region_value` is the read-only attribute containing the
# value of the environment variable
output "aws_default_region_value" {
  value = data.corefunc_env_ensure.aws_default_region.value
}

#=> us-east-1

#-------------------------------------------------------------------------
# AWS_PAGER=""

data "corefunc_env_ensure" "aws_pager" {
  name = "AWS_PAGER"
}

output "aws_pager_value" {
  value = data.corefunc_env_ensure.aws_pager.value
}

#=> [Error] Problem with Environment Variable: environment variable
#=>         AWS_PAGER is not defined

#-------------------------------------------------------------------------
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
