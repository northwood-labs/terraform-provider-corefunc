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
var _ function.Function = &urlDecodeFunction{}

type (
	// urlDecodeFunction is the function implementation.
	urlDecodeFunction struct{}
)

// URLDecodeFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func URLDecodeFunction() function.Function { // lint:allow_return_interface
	return &urlDecodeFunction{}
}

func (f *urlDecodeFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode Function Metadata method.")

	resp.Name = "url_decode"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending URLDecode Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *urlDecodeFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "URLDecode decodes a URL-encoded string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		URLDecode decodes a URL-encoded string.

		It can decode a wide range of characters, including those beyond the ASCII set.
		Non-ASCII characters are first interpreted as UTF-8 bytes, then percent-decoded
		byte-by-byte, ensuring correct decoding of multibyte characters.

		-> This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
		This also provides a 1:1 implementation that can be used with Terratest or other
		Go code.

		Maps to the ` + linkPackage("URLDecode") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "encoded_url",
				MarkdownDescription: "An encoded URL.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending URLDecode Function Definition method.")
}

func (f *urlDecodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting URLDecode Function Run method.")

	var encodedURL string

	err := req.Arguments.Get(ctx, &encodedURL)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value, e := corefunc.URLDecode(encodedURL)
	if e != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(e.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending URLDecode Function Run method.")
}
