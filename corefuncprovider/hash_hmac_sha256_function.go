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
var _ function.Function = &hashHmacSha256Function{}

type (
	// hashHmacSha256Function is the function implementation.
	hashHmacSha256Function struct{}
)

// HashHmacSha256Function is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func HashHmacSha256Function() function.Function { // lint:allow_return_interface
	return &hashHmacSha256Function{}
}

func (f *hashHmacSha256Function) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashHmacSha256 Function Metadata method.")

	resp.Name = "hash_hmac_sha256"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending HashHmacSha256 Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *hashHmacSha256Function) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting HashHmacSha256 Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Generates the HmacSha256 hash of a string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Generates the HmacSha256 hash of a string with its associated salt value.

		Maps to the ` + linkPackage("HashHmacSha256") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "The string to generate the HmacSha256 hash for.",
			},
			function.StringParameter{
				Name:                "key",
				MarkdownDescription: "A secret value to provide additional entropy in the calculation.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending HashHmacSha256 Function Definition method.")
}

func (f *hashHmacSha256Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting HashHmacSha256 Function Run method.")

	var (
		input string
		key   string
	)

	err := req.Arguments.Get(ctx, &input, &key)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value := corefunc.HashHMACSHA256(input, key)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending HashHmacSha256 Function Run method.")
}
