// Copyright 2024-2025, Northwood Labs, LLC <license@northwood-labs.com>
// Copyright 2023-2025, Ryan Parman <rparman@northwood-labs.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"log"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefuncprovider"
)

var (
	debugMode bool

	fPattern string

	// rootCmd represents the base command when called without any subcommands.
	rootCmd = &cobra.Command{
		Use:   "terraform-provider-corefunc",
		Short: "Utilities that should have been Terraform core functions.",
		Long: clihelpers.LongHelpText(`
		# Core Functions

		Utilities that should have been Terraform core functions.

		These CLI commands are implemented in the same way as the Data Sources
		and Functions are, and are meant to help you understand how the provider
		will perform given the same inputs.

		## Repository

		* https://bit.ly/corefunc-github

		## Provider documentation

		* https://bit.ly/corefunc-terraform
		* https://bit.ly/corefunc-opentofu
		* https://bit.ly/corefunc-library-tf

		## Go library documentation

		* https://bit.ly/corefunc-godocs

		## Markdown documentation for LLMs

		* https://bit.ly/corefunc-docs-llm
		`),
		Run: func(_ *cobra.Command, _ []string) {
			err := providerserver.Serve(context.Background(), corefuncprovider.New, providerserver.ServeOpts{
				Address: "registry.terraform.io/northwood-labs/corefunc",
				Debug:   debugMode,
			})
			if err != nil {
				log.Fatal(err.Error())
			}
		},
	}
)

func init() { // lint:allow_init
	// https://developer.hashicorp.com/terraform/plugin/debugging#debugger-based-debugging
	rootCmd.Flags().BoolVarP(
		&debugMode,
		"debug",
		"d",
		false,
		"Run with support for Go debuggers like delve. https://bit.ly/terraform-debugger",
	)
}

// Execute configures the Cobra CLI app framework and executes the root command.
func Execute() {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}

// Root exposes the root command for tools like doc generators.
// https://cobra.dev/docs/how-to-guides/clis-for-llms/
func Root() *cobra.Command {
	return rootCmd
}
