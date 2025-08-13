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

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &tomlToJSONFunction{}

type (
	// tomlToJSONFunction is the function implementation.
	tomlToJSONFunction struct{}
)

// TomlToJSONFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func TomlToJSONFunction() function.Function { // lint:allow_return_interface
	return &tomlToJSONFunction{}
}

func (f *tomlToJSONFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON Function Metadata method.")

	resp.Name = "toml_to_json"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending TomlToJSON Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *tomlToJSONFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Converts a TOML string into an equivalent JSON string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a TOML string into an equivalent JSON string.

		Maps to the ` + linkPackage("TOMLtoJSON") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "toml",
				MarkdownDescription: "The TOML string to convert to JSON.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending TomlToJSON Function Definition method.")
}

func (f *tomlToJSONFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting TomlToJSON Function Run method.")

	var asTOML string

	resp.Error = req.Arguments.Get(ctx, &asTOML)
	if resp.Error != nil {
		return
	}

	jsonVal, err := corefunc.TOMLtoJSON(asTOML)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, jsonVal))

	tflog.Debug(ctx, "Ending TomlToJSON Function Run method.")
}
