# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_camel
data "corefunc_str_camel" "str" {
  string = local.input_string
}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_constant
data "corefunc_str_constant" "str" {
  string = local.input_string
}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_kebab
data "corefunc_str_kebab" "str" {
  string = local.input_string
}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_pascal
data "corefunc_str_pascal" "str" {
  string = local.input_string
}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/str_snake
data "corefunc_str_snake" "str" {
  string = local.input_string
}
