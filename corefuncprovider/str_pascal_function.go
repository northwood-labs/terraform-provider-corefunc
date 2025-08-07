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

package corefuncprovider // lint:no_dupe

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &strPascalFunction{}

type (
	// strPascalFunction is the function implementation.
	strPascalFunction struct{}
)

// StrPascalFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrPascalFunction() function.Function { // lint:allow_return_interface
	return &strPascalFunction{}
}

func (f *strPascalFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal Function Metadata method.")

	resp.Name = "str_pascal"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending StrPascal Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strPascalFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Converts a string to `PascalCase` removing any non-alphanumeric characters.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `PascalCase` + "`" + `, removing any non-alphanumeric characters.

		-> Some acronyms are maintained as uppercase. See
		[caps: pkg-variables](https://pkg.go.dev/github.com/chanced/caps#pkg-variables) for a complete list.

		Maps to the ` + linkPackage("StrPascal") + ` Go method, which can be used in ` + Terratest + `.
        `)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "str",
				MarkdownDescription: "The string to convert to `PascalCase`.",
			},
			function.BoolParameter{
				Name: "use_acronym_caps",
				MarkdownDescription: "Whether or not to keep acronyms as uppercase. A value of `true` means " +
					"that acronyms will be converted to uppercase. A value of `false` means that acronyms " +
					"will using typical casing.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrPascal Function Definition method.")
}

func (f *strPascalFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrPascal Function Run method.")

	var (
		str            string
		useAcronymCaps bool
	)

	err := req.Arguments.Get(ctx, &str, &useAcronymCaps)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.StrPascal(str, useAcronymCaps)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrPascal Function Run method.")
}
