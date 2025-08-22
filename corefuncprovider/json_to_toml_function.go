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
var _ function.Function = &jsonToTomlFunction{}

type (
	// jsonToTomlFunction is the function implementation.
	jsonToTomlFunction struct{}
)

// JSONToTomlFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func JSONToTomlFunction() function.Function { // lint:allow_return_interface
	return &jsonToTomlFunction{}
}

func (f *jsonToTomlFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml Function Metadata method.")

	resp.Name = "json_to_toml"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending JSONToToml Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *jsonToTomlFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Converts a JSON string into an equivalent TOML string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a JSON string into an equivalent TOML string.

		Maps to the ` + linkPackage("JSONtoTOML") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "toml",
				MarkdownDescription: "The JSON string to convert to TOML.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending JSONToToml Function Definition method.")
}

func (f *jsonToTomlFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting JSONToToml Function Run method.")

	var asJSON string

	resp.Error = req.Arguments.Get(ctx, &asJSON)
	if resp.Error != nil {
		return
	}

	tomlVal, err := corefunc.JSONtoTOML(asJSON)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, strings.TrimSpace(tomlVal)))

	tflog.Debug(ctx, "Ending JSONToToml Function Run method.")
}
