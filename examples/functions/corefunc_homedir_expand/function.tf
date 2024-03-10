output "homedir" {
  value = provider::corefunc::homedir_expand("~/.aws")
}
