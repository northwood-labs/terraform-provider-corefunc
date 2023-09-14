/*
See "corefunc" package for the Go library code.
*/
package main

import (
	"github.com/northwood-labs/terraform-provider-corefunc/cmd"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name corefunc --rendered-provider-name "Core Functions" # lint:ignore_length

func main() {
	cmd.Execute()
}
