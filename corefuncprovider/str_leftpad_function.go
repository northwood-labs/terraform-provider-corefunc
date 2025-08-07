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
var _ function.Function = &strLeftpadFunction{}

type (
	// strLeftpadFunction is the function implementation.
	strLeftpadFunction struct{}
)

// StrLeftpadFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrLeftpadFunction() function.Function { // lint:allow_return_interface
	return &strLeftpadFunction{}
}

func (f *strLeftpadFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad Function Metadata method.")

	resp.Name = "str_leftpad"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending StrLeftpad Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strLeftpadFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Pads a string with additional characters on the left.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Pads a string with additional characters on the left.

		Maps to the ` + linkPackage("StrLeftPad") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "str",
				MarkdownDescription: "The string to pad with padding characters.",
			},
			function.Int64Parameter{
				Name:                "pad_width",
				MarkdownDescription: "The max number of padding characters to pad the string with.",
			},
		},
		VariadicParameter: function.StringParameter{
			Name: "pad_char",
			MarkdownDescription: "The padding character to use. Only supports a single byte. If more than one " +
				"byte is provided, only the first byte will be used. The default value is a space character.",
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrLeftpad Function Definition method.")
}

func (f *strLeftpadFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrLeftpad Function Run method.")

	var (
		str      string
		value    string
		padWidth int64
		padChar  []string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &str, &padWidth, &padChar))
	if resp.Error != nil {
		return
	}

	if len(padChar) > 0 && padChar[0] != "" {
		b := []byte(padChar[0])
		value = corefunc.StrLeftPad(str, int(padWidth), b[0])
	} else {
		value = corefunc.StrLeftPad(str, int(padWidth))
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrLeftpad Function Run method.")
}
