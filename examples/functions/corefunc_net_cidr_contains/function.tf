output "cidr" {
  value = provider::corefunc::net_cidr_contains("192.168.2.0/20", "192.168.2.1")
}

#=> true
