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
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// homedirExpandCmd represents the homedir-expand command
var homedirExpandCmd = &cobra.Command{
	Use:   "homedir-expand [path]",
	Short: "Expands a file path that begins with a tilde (~) to the user's home directory.",
	Long: clihelpers.LongHelpText(`
	Expands a file path that begins with a tilde (~) to the user's home directory.
	`),
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("The first argument must be a file path to expand.")
			fmt.Println("Re-run with --help to see options.")
			os.Exit(1)
		}

		homedir, err := corefunc.HomedirExpand(args[0])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(homedir)
	},
}

func init() { // lint:allow_init
	rootCmd.AddCommand(homedirExpandCmd)
}
