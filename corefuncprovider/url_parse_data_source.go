// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
	cftypes "github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &urlParseDataSource{}
	_ datasource.DataSourceWithConfigure = &urlParseDataSource{}
)

// urlParseDataSource is the data source implementation.
type (
	urlParseDataSource struct{}

	// urlParseDataSourceModel maps the data source schema data.
	urlParseDataSourceModel struct {
		InputURL         types.String `tfsdk:"url"`
		Canonicalizer    types.String `tfsdk:"canonicalizer"`
		Normalized       types.String `tfsdk:"normalized"`
		NormalizedNoFrag types.String `tfsdk:"normalized_nofrag"`
		Protocol         types.String `tfsdk:"protocol"`
		Scheme           types.String `tfsdk:"scheme"`
		Username         types.String `tfsdk:"username"`
		Password         types.String `tfsdk:"password"`
		Hostname         types.String `tfsdk:"hostname"`
		Host             types.String `tfsdk:"host"`
		Port             types.String `tfsdk:"port"`
		Path             types.String `tfsdk:"path"`
		Search           types.String `tfsdk:"search"`
		Query            types.String `tfsdk:"query"`
		Hash             types.String `tfsdk:"hash"`
		Fragment         types.String `tfsdk:"fragment"`
		DecodedPort      types.Int64  `tfsdk:"decoded_port"`
	}
)

// URLParseDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func URLParseDataSource() datasource.DataSource { // lint:allow_return_interface
	return &urlParseDataSource{}
}

// Metadata returns the data source type name.
func (d *urlParseDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting URLParse DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_url_parse"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending URLParse DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *urlParseDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting URLParse DataSource Schema method.")

	resp.Schema = schema.Schema{
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
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Description: "The absolute URL to parse according to the [WHATWG URL API](https://url.spec.whatwg.org/#api).",
				Required:    true,
			},
			"canonicalizer": schema.StringAttribute{
				Description: "The method by which the URL should be canonicalized. " +
					"Allowed values: `standard`, `google_safe_browsing`. The default value is `standard`.",
				Optional: true,
			},
			"canonicalizer": schema.StringAttribute{
				MarkdownDescription: "The method by which the URL should be canonicalized. " +
					"Allowed values: `standard`, `google_safe_browsing`. The default value is `standard`.",
				Optional: true,
			},
			"normalized": schema.StringAttribute{
				Description: "The normalized form of the URL.",
				Computed:    true,
			},
			"normalized_nofrag": schema.StringAttribute{
				Description: "The normalized form of the URL, without the fragment.",
				Computed:    true,
			},
			"protocol": schema.StringAttribute{
				Description: "The protocol of the URL, including the final `:`.",
				Computed:    true,
			},
			"scheme": schema.StringAttribute{
				Description: "The scheme of the URL.",
				Computed:    true,
			},
			"username": schema.StringAttribute{
				Description: "The username that was passed to the URL as basic authentication.",
				Computed:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password that was passed to the URL as basic authentication.",
				Computed:    true,
			},
			"host": schema.StringAttribute{
				Description: "The host of the URL, including the port if web browsers typically show " +
					"the port in the URL bar.",
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Description: "The hostname of the URL.",
				Computed:    true,
			},
			"port": schema.StringAttribute{
				Description: "The port of the URL. May be blank if web browsers typically hide the port in the URL bar.",
				Computed:    true,
			},
			"decoded_port": schema.Int64Attribute{
				Description: "The port of the URL.",
				Computed:    true,
			},
			"path": schema.StringAttribute{
				Description: "The path of the URL.",
				Computed:    true,
			},
			"search": schema.StringAttribute{
				Description: "The query string of the URL, including the preceding `?`.",
				Computed:    true,
			},
			"query": schema.StringAttribute{
				Description: "The query string of the URL.",
				Computed:    true,
			},
			"hash": schema.StringAttribute{
				Description: "The fragment string of the URL, including the preceding `#`.",
				Computed:    true,
			},
			"fragment": schema.StringAttribute{
				Description: "The fragment string of the URL.",
				Computed:    true,
			},
		},
	}

	tflog.Debug(ctx, "Ending URLParse DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *urlParseDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting URLParse DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending URLParse DataSource Configure method.")
}

func (d *urlParseDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting URLParse DataSource Create method.")

	var plan urlParseDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending URLParse DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *urlParseDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting URLParse DataSource Read method.")

	var state urlParseDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	// Default
	opts := cftypes.Standard
	if strings.EqualFold(state.Canonicalizer.ValueString(), "google_safe_browsing") {
		// Switch to Google Safe Browsing
		opts = cftypes.GoogleSafeBrowsing
	}

	url, err := corefunc.URLParse(state.InputURL.ValueString(), opts)
	if err != nil {
		resp.Diagnostics.AddError("Could not parse the URL.", err.Error())
		return
	}

	state.Normalized = types.StringValue(url.Href(false))
	state.NormalizedNoFrag = types.StringValue(url.Href(true))
	state.Protocol = types.StringValue(url.Protocol())
	state.Scheme = types.StringValue(url.Scheme())
	state.Username = types.StringValue(url.Username())
	state.Password = types.StringValue(url.Password())
	state.Host = types.StringValue(url.Host())
	state.Hostname = types.StringValue(url.Hostname())
	state.Port = types.StringValue(url.Port())
	state.Path = types.StringValue(url.Pathname())
	state.Search = types.StringValue(url.Search())
	state.Query = types.StringValue(url.Query())
	state.Hash = types.StringValue(url.Hash())
	state.Fragment = types.StringValue(url.Fragment())
	state.DecodedPort = types.Int64Value(int64(url.DecodedPort()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending URLParse DataSource Read method.")
}
