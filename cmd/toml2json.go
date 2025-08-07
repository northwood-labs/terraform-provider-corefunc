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

package cmd // lint:no_dupe

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// toml2jsonCmd represents the toml2json command
var toml2jsonCmd = &cobra.Command{
	Use:   "toml2json",
	Short: "Converts TOML to JSON.",
	Long: clihelpers.LongHelpText(`
	Converts TOML to JSON.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Re-run with --help to see options.")
			os.Exit(1)
		}

		fString, err := corefunc.File(args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		output, err := corefunc.TOMLtoJSON(fString)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(output)
	},
}

func init() { // lint:allow_init
	rootCmd.AddCommand(toml2jsonCmd)
}
