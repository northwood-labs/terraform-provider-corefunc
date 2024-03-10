output "runtime_cpuarch_fn" {
  description = "This returns the CPU architecture of the provider."
  value       = provider::corefunc::runtime_cpuarch()
}

output "runtime_goroot_fn" {
  description = "This returns the GOROOT directory, if it exists."
  value       = provider::corefunc::runtime_goroot()
}

output "runtime_numcpus_fn" {
  description = "This returns the number of logical CPU cores."
  value       = provider::corefunc::runtime_numcpus()
}

output "runtime_os_fn" {
  description = "This returns the operating system of the provider."
  value       = provider::corefunc::runtime_os()
}
