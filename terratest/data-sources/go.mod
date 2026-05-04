module github.com/northwood-labs/terraform-provider-corefunc/v2/terratest

go 1.26

replace github.com/northwood-labs/terraform-provider-corefunc/v2 v2.2.1 => ../../

require (
	github.com/charmbracelet/lipgloss v1.1.1-0.20250404203927-76690c660834
	github.com/gruntwork-io/terratest v0.56.0
	github.com/hairyhenderson/go-which v0.2.2
	github.com/northwood-labs/cli-helpers v0.0.0-20260428015633-ec689f4e2063
	github.com/northwood-labs/terraform-provider-corefunc/v2 v2.2.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.11.1
)
