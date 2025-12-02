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
var _ function.Function = &strContainsFunction{}

type (
	// strContainsFunction is the function implementation.
	strContainsFunction struct{}
)

// StrContainsFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrContainsFunction() function.Function { // lint:allow_return_interface
	return &strContainsFunction{}
}

func (f *strContainsFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrContains Function Metadata method.")

	resp.Name = "str_contains"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending StrContains Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strContainsFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrContains Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Checks if a string contains a given substring.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Checks if a string contains a given substring.

		-> This is identical to the built-in ` + "`strcontains`" + ` function
		which was added in Terraform 1.3. This provides a 1:1 implementation
		that can be used with Terratest or other Go code, as well as with
		OpenTofu and Terraform going all the way back to v1.0.

		Maps to the ` + linkPackage("StrContains") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "The string to check.",
			},
			function.StringParameter{
				Name:                "substr",
				MarkdownDescription: "The substring to look for.",
			},
		},
		Return: function.BoolReturn{},
	}

	tflog.Debug(ctx, "Ending StrContains Function Definition method.")
}

func (f *strContainsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrContains Function Run method.")

	var (
		input  string
		substr string
	)

	err := req.Arguments.Get(ctx, &input, &substr)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.StrContains(input, substr)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrContains Function Run method.")
}
