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
var _ function.Function = &intLeftpadFunction{}

type (
	// intLeftpadFunction is the function implementation.
	intLeftpadFunction struct{}
)

// IntLeftpadFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func IntLeftpadFunction() function.Function { // lint:allow_return_interface
	return &intLeftpadFunction{}
}

func (f *intLeftpadFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad Function Metadata method.")

	resp.Name = "int_leftpad"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending IntLeftpad Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *intLeftpadFunction) Definition( // lint:no_dupe
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Converts an integer to a string, and then pads it with zeroes on the left.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts an integer to a string, and then pads it with zeroes on the left.

		-> If the integer is NOT in base10 (decimal), it will be converted to base10 (decimal) _before_ being padded.

		Maps to the ` + linkPackage("IntLeftPad") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.Int64Parameter{
				Name:                "num",
				MarkdownDescription: "The integer to pad with zeroes.",
			},
			function.Int64Parameter{
				Name:                "pad_width",
				MarkdownDescription: "The max number of zeroes to pad the integer with.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending IntLeftpad Function Definition method.")
}

func (f *intLeftpadFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting IntLeftpad Function Run method.")

	var (
		num      int64
		padWidth int64
	)

	err := req.Arguments.Get(ctx, &num, &padWidth)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.IntLeftPad(num, int(padWidth))

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending IntLeftpad Function Run method.")
}
