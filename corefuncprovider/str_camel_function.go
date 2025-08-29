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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &strCamelFunction{}

type (
	// strCamelFunction is the function implementation.
	strCamelFunction struct{}
)

// StrCamelFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrCamelFunction() function.Function { // lint:allow_return_interface
	return &strCamelFunction{}
}

func (f *strCamelFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel Function Metadata method.")

	resp.Name = "str_camel"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending StrCamel Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strCamelFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel Function Definition method.")

	const desc = `, removing any non-alphanumeric characters.

		-> Some acronyms are maintained as uppercase. See
		[caps: pkg-variables](https://pkg.go.dev/github.com/chanced/caps#pkg-variables) for a complete list.

		Maps to the `

	resp.Definition = function.Definition{
		Summary: "Converts a string to `camelCase`, removing any non-alphanumeric characters.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `camelCase` + "`" + desc + linkPackage("StrCamel") +
			` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "str",
				MarkdownDescription: "The string to convert to `camelCase`.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrCamel Function Definition method.")
}

func (f *strCamelFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrCamel Function Run method.")

	var str string

	err := req.Arguments.Get(ctx, &str)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.StrCamel(str)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrCamel Function Run method.")
}
