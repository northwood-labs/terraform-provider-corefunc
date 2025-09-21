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
	"regexp"

	"github.com/spf13/cobra"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// envEnsureCmd represents the env-ensure command
var envEnsureCmd = &cobra.Command{
	Use:   "env-ensure",
	Short: "Ensures that an environment variable is set.",
	Long: clihelpers.LongHelpText(`
	Ensures that an environment variable is set, optionally validating its value against a provided pattern.
	`),
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("The first argument must be an environment variable to validate.")
			fmt.Println("Re-run with --help to see options.")
			os.Exit(1)
		}

		err := corefunc.EnvEnsure(args[0], regexp.MustCompile(fPattern))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(os.Getenv(args[0]))
	},
}

func init() { // lint:allow_init
	envEnsureCmd.Flags().
		StringVarP(&fPattern, "pattern", "p", "", "A valid Go ([re2](https://github.com/google/re2/wiki/Syntax)) "+
			"regular expression pattern.")

	rootCmd.AddCommand(envEnsureCmd)
}
