# AWS_PAGER=""

data "corefunc_env_ensure" "aws_pager" {
  name = "AWS_PAGER"
}

output "aws_pager_value" {
  value = data.corefunc_env_ensure.aws_pager.value
}

#=> [Error] Problem with Environment Variable: environment variable
#=>         AWS_PAGER is not defined
