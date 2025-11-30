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
var _ function.Function = &hashSha384Function{}

type (
	// hashSha384Function is the function implementation.
	hashSha384Function struct{}
)

// HashSha384Function is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func HashSha384Function() function.Function { // lint:allow_return_interface
	return &hashSha384Function{}
}

func (f *hashSha384Function) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 Function Metadata method.")

	resp.Name = "hash_sha384"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending HashSha384 Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *hashSha384Function) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Generates the SHA-2/384-bit hash of a string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Generates the SHA-2/384-bit hash of a string.

		Maps to the ` + linkPackage("HashSHA384") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "The string to generate a SHA-2/384-bit hash for.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending HashSha384 Function Definition method.")
}

func (f *hashSha384Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting HashSha384 Function Run method.")

	var input string

	err := req.Arguments.Get(ctx, &input)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.HashSHA384(input)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending HashSha384 Function Run method.")
}
