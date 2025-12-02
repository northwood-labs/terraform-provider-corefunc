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
var _ function.Function = &strStartswithFunction{}

type (
	// strStartswithFunction is the function implementation.
	strStartswithFunction struct{}
)

// StrStartswithFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrStartswithFunction() function.Function { // lint:allow_return_interface
	return &strStartswithFunction{}
}

func (f *strStartswithFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith Function Metadata method.")

	resp.Name = "str_startswith"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending StrStartswith Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strStartswithFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Takes a string to check and a prefix string, and returns true if the string begins" +
			" with that exact prefix.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Takes a string to check and a prefix string, and returns true if the
		string begins with that exact prefix.

		-> This is identical to the built-in ` + "`startswith`" + ` function
		which was added in Terraform 1.3. This provides a 1:1 implementation
		that can be used with Terratest or other Go code, as well as with
		OpenTofu and Terraform going all the way back to v1.0.

		Maps to the ` + linkPackage("StrStartswith") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "The string to check.",
			},
			function.StringParameter{
				Name:                "prefix",
				MarkdownDescription: "The prefix to check for.",
			},
		},
		Return: function.BoolReturn{},
	}

	tflog.Debug(ctx, "Ending StrStartswith Function Definition method.")
}

func (f *strStartswithFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrStartswith Function Run method.")

	var (
		input  string
		prefix string
	)

	err := req.Arguments.Get(ctx, &input, &prefix)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.StrStartsWith(input, prefix)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrStartswith Function Run method.")
}
