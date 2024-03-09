# AWS_PAGER=""

output "aws_pager" {
  value = provider::corefunc::env_ensure("AWS_PAGER")
}

#=> [Error] Problem with Environment Variable: environment variable
#=>         AWS_PAGER is not defined
