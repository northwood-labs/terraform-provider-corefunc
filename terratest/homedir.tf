# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/homedir_get
data "corefunc_homedir_get" "home" {}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/homedir_expand
data "corefunc_homedir_expand" "home" {
  path = "~/.aws/"
}
