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
var _ function.Function = &hashSha1Base64Function{}

type (
	// hashSha1Base64Function is the function implementation.
	hashSha1Base64Function struct{}
)

// HashSha1Base64Function is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func HashSha1Base64Function() function.Function { // lint:allow_return_interface
	return &hashSha1Base64Function{}
}

func (f *hashSha1Base64Function) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashSha1Base64 Function Metadata method.")

	resp.Name = "hash_sha1_base64"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending HashSha1Base64 Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *hashSha1Base64Function) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting HashSha1Base64 Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Generates the SHA-1 hash of a string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Generates the SHA-1 hash of a string. Returns a Base64 value instead of a hexadecimal string.

		Maps to the ` + linkPackage("Base64HashSHA1") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "The string to generate an SHA-1 hash for.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending HashSha1Base64 Function Definition method.")
}

func (f *hashSha1Base64Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting HashSha1Base64 Function Run method.")

	var input string

	err := req.Arguments.Get(ctx, &input)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.Base64HashSHA1(input)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending HashSha1Base64 Function Run method.")
}
