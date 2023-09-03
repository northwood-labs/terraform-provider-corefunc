package main

import (
	"context"
	"flag"
	"log"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name corefunc --rendered-provider-name "Core Functions" # lint:ignore_length

func main() {
	var debug bool

	// https://developer.hashicorp.com/terraform/plugin/debugging#debugger-based-debugging
	flag.BoolVar(&debug, "debug", false, "Run with support for Go debuggers like delve. https://flwd.dk/3k055Lh")
	// version?
	// deep version?
	flag.Parse()

	err := providerserver.Serve(context.Background(), corefunc.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/northwood-labs/corefunc",
		Debug:   debug,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
