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
var _ function.Function = &strBase64GunzipFunction{}

type (
	// strBase64GunzipFunction is the function implementation.
	strBase64GunzipFunction struct{}
)

// StrBase64GunzipFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrBase64GunzipFunction() function.Function { // lint:allow_return_interface
	return &strBase64GunzipFunction{}
}

func (f *strBase64GunzipFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip Function Metadata method.")

	resp.Name = "str_base64_gunzip"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending StrBase64Gunzip Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strBase64GunzipFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Base64Gunzip is a function that decodes a Base64-encoded string, then decompresses the result with gzip.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Base64Gunzip is a function that decodes a Base64-encoded string, then
		decompresses the result with gzip. Supports both padded and non-padded Base64
		strings.

		Uses the "standard" Base64 alphabet as defined in
		[RFC 4648 §4](https://datatracker.ietf.org/doc/html/rfc4648#section-4).

		~> There is a data limit of 10 MiB (10485760 bytes) for the decompressed data. This
		is to avoid "decompression bomb" vulnerabilities.

		-> This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
		This also provides a 1:1 implementation that can be used with Terratest or other
		Go code.

		Maps to the ` + linkPackage("Base64Gunzip") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "gzipped_base64",
				MarkdownDescription: "A string of gzipped then Base64-encoded data.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrBase64Gunzip Function Definition method.")
}

func (f *strBase64GunzipFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip Function Run method.")

	var gzippedBase64 string
	err := req.Arguments.Get(ctx, &gzippedBase64)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value, e := corefunc.Base64Gunzip(gzippedBase64)
	if e != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(e.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrBase64Gunzip Function Run method.")
}
