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
