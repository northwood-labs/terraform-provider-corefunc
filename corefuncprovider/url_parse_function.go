// Copyright 2023-2025, Northwood Labs
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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
	cftypes "github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &urlParseFunction{}

type (
	// urlParseFunction is the function implementation.
	urlParseFunction struct{}

	parsedURL struct {
		Fragment         string `tfsdk:"fragment"`
		Hash             string `tfsdk:"hash"`
		Host             string `tfsdk:"host"`
		Hostname         string `tfsdk:"hostname"`
		NormalizedNofrag string `tfsdk:"normalized_nofrag"`
		Normalized       string `tfsdk:"normalized"`
		Password         string `tfsdk:"password"`
		Path             string `tfsdk:"path"`
		Port             string `tfsdk:"port"`
		Protocol         string `tfsdk:"protocol"`
		Query            string `tfsdk:"query"`
		Scheme           string `tfsdk:"scheme"`
		Search           string `tfsdk:"search"`
		Username         string `tfsdk:"username"`
		DecodedPort      int64  `tfsdk:"decoded_port"`
	}
)

// URLParseFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func URLParseFunction() function.Function { // lint:allow_return_interface
	return &urlParseFunction{}
}

func (f *urlParseFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting URLParse Function Metadata method.")

	resp.Name = "url_parse"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending URLParse Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *urlParseFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting URLParse Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "URLParse is a [WHATWG spec-compliant](https://url.spec.whatwg.org/#url-parsing) URL parser " +
			"returns a struct representing a parsed absolute URL.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		URLParse is a [WHATWG spec-compliant](https://url.spec.whatwg.org/#url-parsing) URL parser returns a struct
		representing a parsed absolute URL.

		Will not resolve relative URLs. The parser is up to date as of
		[2023-05-24](https://url.spec.whatwg.org/commit-snapshots/eee49fdf4f99d59f717cbeb0bce29fda930196d4/)
		and passes all relevant tests from
		[web-platform-tests](https://github.com/web-platform-tests/wpt/tree/master/url). Its API is similar to Chapter 6
		in [WHATWG URL Standard](https://url.spec.whatwg.org/#api).

		Maps to the ` + linkPackage("URLParse") + ` Go method, which can be used in ` + Terratest + `.
        `)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name: "url",
				MarkdownDescription: "The absolute URL to parse according to the " +
					"[WHATWG URL API](https://url.spec.whatwg.org/#api).",
			},
		},
		VariadicParameter: function.StringParameter{
			Name: "canonicalizer",
			MarkdownDescription: "The method by which the URL should be canonicalized. " +
				"Allowed values: `standard`, `google_safe_browsing`. The default value is `standard`.",
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"fragment":          types.StringType,
				"hash":              types.StringType,
				"host":              types.StringType,
				"hostname":          types.StringType,
				"normalized_nofrag": types.StringType,
				"normalized":        types.StringType,
				"password":          types.StringType,
				"path":              types.StringType,
				"port":              types.StringType,
				"protocol":          types.StringType,
				"query":             types.StringType,
				"scheme":            types.StringType,
				"search":            types.StringType,
				"username":          types.StringType,
				"decoded_port":      types.Int64Type,
			},
		},
	}

	tflog.Debug(ctx, "Ending URLParse Function Definition method.")
}

func (f *urlParseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting URLParse Function Run method.")

	var (
		url           string
		canonicalizer []string
	)

	resp.Error = req.Arguments.Get(ctx, &url, &canonicalizer)
	if resp.Error != nil {
		return
	}

	// Default
	opts := cftypes.Standard

	if len(canonicalizer) > 0 {
		switch strings.ToLower(canonicalizer[0]) {
		case "standard":
			opts = cftypes.Standard
		case "google_safe_browsing":
			opts = cftypes.GoogleSafeBrowsing
		default:
			resp.Error = function.ConcatFuncErrors(
				function.NewArgumentFuncError(1, "canonicalizer must be one of: `standard`, `google_safe_browsing`"),
			)

			return
		}
	}

	value, err := corefunc.URLParse(url, opts)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)

		return
	}

	result := parsedURL{
		Normalized:       value.Href(false),
		NormalizedNofrag: value.Href(true),
		Protocol:         value.Protocol(),
		Scheme:           value.Scheme(),
		Username:         value.Username(),
		Password:         value.Password(),
		Host:             value.Host(),
		Hostname:         value.Hostname(),
		Port:             value.Port(),
		Path:             value.Pathname(),
		Search:           value.Search(),
		Query:            value.Query(),
		Hash:             value.Hash(),
		Fragment:         value.Fragment(),
		DecodedPort:      int64(value.DecodedPort()),
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))

	tflog.Debug(ctx, "Ending URLParse Function Run method.")
}
