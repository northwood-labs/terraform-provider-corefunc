package cmd

import (
	clihelpers "github.com/northwood-labs/cli-helpers"
)

var versionCmd = clihelpers.VersionScreen()

func init() { // lint:allow_init
	rootCmd.AddCommand(versionCmd)
}
