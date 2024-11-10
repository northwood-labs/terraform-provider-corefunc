data "corefunc_net_cidr_contains" "cidr" {
  containing_cidr      = "192.168.2.0/20"
  contained_ip_or_cidr = "192.168.2.1"
}

#=> true
