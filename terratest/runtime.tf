# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/runtime_cpuarch
data "corefunc_runtime_cpuarch" "arch" {}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/runtime_goroot
data "corefunc_runtime_goroot" "goroot" {}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/runtime_numcpus
data "corefunc_runtime_numcpus" "count" {}

# https://registry.terraform.io/providers/northwood-labs/corefunc/latest/docs/data-sources/runtime_os
data "corefunc_runtime_os" "os" {}
