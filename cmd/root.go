// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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

	"github.com/spf13/cobra"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/northwood-labs/terraform-provider-corefunc/corefuncprovider"
)

var (
	debugMode bool

	// rootCmd represents the base command when called without any subcommands.
	rootCmd = &cobra.Command{
		Use:   "terraform-provider-corefunc",
		Short: "Utilities that should have been Terraform core functions.",
		Long: `--------------------------------------------------------------------------------
terraform-provider-corefunc

Utilities that should have been Terraform core functions.

While some of these *can* be implemented in HCL, some of them begin to push up
against the limits of Terraform and the HCL2 configuration language. We also
perform testing using the Terratest <https://terratest.gruntwork.io> framework
on a regular basis. Exposing these functions as both a Go library as well as a
Terraform provider enables us to use the same functionality in both our
Terraform applies as well as while using a testing framework.

Since Terraform doesn't have the concept of user-defined functions, the next
step to open up the possibilities is to write a custom Terraform Provider which
has the functions built-in, using Terraform's existing support for inputs and
outputs.

**This does not add new syntax or constructs to Terraform.** Instead it uses the
existing concepts around Providers, Resources, Data Sources, Variables, and
Outputs to expose new custom-built functionality.

The goal of this provider is not to call any APIs, but to provide pre-built
functions in the form of Data Sources.
--------------------------------------------------------------------------------`,
		Run: func(cmd *cobra.Command, args []string) {
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
		"Run with support for Go debuggers like delve. https://flwd.dk/3k055Lh",
	)
}

// Execute configures the Cobra CLI app framework and executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
