data "corefunc_runtime_cpuarch" "arch" {}

data "corefunc_runtime_goroot" "goroot" {}

data "corefunc_runtime_numcpus" "count" {}

data "corefunc_runtime_os" "os" {}

output "runtime_cpuarch_ds" {
  description = "This returns the CPU architecture of the provider."
  value       = data.corefunc_runtime_cpuarch.arch.value
}

output "runtime_goroot_ds" {
  description = "This returns the GOROOT directory, if it exists."
  value       = data.corefunc_runtime_goroot.goroot.value
}

output "runtime_numcpus_ds" {
  description = "This returns the number of logical CPU cores."
  value       = data.corefunc_runtime_numcpus.count.value
}

output "runtime_os_ds" {
  description = "This returns the operating system of the provider."
  value       = data.corefunc_runtime_os.os.value
}
