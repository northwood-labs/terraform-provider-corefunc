// Copyright 2023-2024, Ryan Parman
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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"text/tabwriter"

	"github.com/gookit/color"
	"github.com/northwood-labs/golang-utils/archstring"
	"github.com/northwood-labs/golang-utils/exiterrorf"
	"github.com/spf13/cobra"
)

var (
	// Color text.
	colorHeader = color.New(color.FgWhite, color.BgBlue, color.OpBold)

	// Version represents the version of the software.
	Version = "dev"

	// Commit represents the git commit hash of the software.
	Commit = vcs("vcs.revision", "unknown")

	// BuildDate represents the date the software was built.
	BuildDate = vcs("vcs.time", "unknown")

	// Dirty represents whether or not the git repo was dirty when the software was built.
	Dirty = vcs("vcs.modified", "unknown")

	// PGOEnabled represents whether or not the build leveraged Profile-Guided Optimization (PGO).
	PGOEnabled = vcs("-pgo", "false")

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Long-form version information",
		Long: `Long-form version information, including the build commit hash, build date, Go
version, and external dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			colorHeader.Println(" BASIC ")

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

			fmt.Fprintf(w, " Version:\t%s\t\n", Version)
			fmt.Fprintf(w, " Go version:\t%s\t\n", runtime.Version())
			fmt.Fprintf(w, " Git commit:\t%s\t\n", Commit)
			if Dirty == "true" {
				fmt.Fprintf(w, " Dirty repo:\t%s\t\n", Dirty)
			}
			if !strings.Contains(PGOEnabled, "false") {
				fmt.Fprintf(w, " PGO:\t%s\t\n", filepath.Base(PGOEnabled))
			}
			fmt.Fprintf(w, " Build date:\t%s\t\n", BuildDate)
			fmt.Fprintf(w, " OS/Arch:\t%s/%s\t\n", runtime.GOOS, runtime.GOARCH)
			fmt.Fprintf(w, " System:\t%s\t\n", archstring.GetFriendlyName(runtime.GOOS, runtime.GOARCH))
			fmt.Fprintf(w, " CPU Cores:\t%d\t\n", runtime.NumCPU())

			err := w.Flush()
			if err != nil {
				exiterrorf.ExitErrorf(err)
			}

			fmt.Println("")

			//----------------------------------------------------------------------

			if buildInfo, ok := debug.ReadBuildInfo(); ok {
				colorHeader.Println(" DEPENDENCIES ")

				w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

				for i := range buildInfo.Deps {
					dependency := buildInfo.Deps[i]
					fmt.Fprintf(w, " %s\t%s\t\n", dependency.Path, dependency.Version)
				}
			}

			err = w.Flush()
			if err != nil {
				exiterrorf.ExitErrorf(err)
			}

			// if info, ok := debug.ReadBuildInfo(); ok {
			// 	for i := range info.Settings {
			// 		setting := info.Settings[i]

			// 		fmt.Printf("%s = %s\n", setting.Key, setting.Value)
			// 	}
			// }

			fmt.Println("")
		},
	}
)

func init() { // lint:allow_init
	rootCmd.AddCommand(versionCmd)
}

func vcs(key, fallback string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for i := range info.Settings {
			setting := info.Settings[i]

			if setting.Key == key {
				return setting.Value
			}
		}
	}

	return fallback
}
