output "net_cidr_contains_fn" {
  description = "This is whether or not the IP address is contained within the CIDR range."
  value       = provider::corefunc::net_cidr_contains("192.168.2.0/20", "192.168.2.1")
}
