# AWS_DEFAULT_REGION="us-east-1"

output "aws_default_region" {
  value = provider::corefunc::env_ensure("AWS_DEFAULT_REGION")
}

#=> us-east-1
