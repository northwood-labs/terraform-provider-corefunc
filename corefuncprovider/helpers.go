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

package corefuncprovider

import (
	"html/template"
	"runtime"
	"strings"
)

const (
	// Terratest displays a Markdown link to Terratest.
	Terratest = "[Terratest](https://terratest.gruntwork.io)"

	// TPF displays a Markdown link to the Terraform Plugin Framework.
	TPF = "[Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework)"
)

// FuncMap returns a template.FuncMap with helper functions for templates.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"contains": strings.Contains,
		// "contains": func(input, find string) bool {
		// 	return strings.Contains(input, find)
		// },
	}
}

func traceFuncName() string { // lint:allow_unused
	pc := make([]uintptr, 15)   // lint:allow_raw_number
	n := runtime.Callers(2, pc) // lint:allow_raw_number
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	out := strings.TrimPrefix(
		frame.Function,
		"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc.",
	)

	out = strings.TrimPrefix(
		out,
		"github.com/northwood-labs/terraform-provider-corefunc/v2/corefuncprovider.",
	)

	return out + "\n"
}

func linkPackage(functionName string) string {
	return "[`corefunc." +
		functionName +
		"()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc#" +
		functionName +
		")"
}
