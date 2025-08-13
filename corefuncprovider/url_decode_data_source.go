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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &urlDecodeDataSource{}
	_ datasource.DataSourceWithConfigure = &urlDecodeDataSource{}
)

// urlDecodeDataSource is the data source implementation.
type (
	urlDecodeDataSource struct{}

	// urlDecodeDataSourceModel maps the data source schema data.
	urlDecodeDataSourceModel struct {
		EncodedURL types.String `tfsdk:"encoded_url"`
		Value      types.String `tfsdk:"value"`
	}
)

// URLDecodeDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func URLDecodeDataSource() datasource.DataSource { // lint:allow_return_interface
	return &urlDecodeDataSource{}
}

// Metadata returns the data source type name.
func (d *urlDecodeDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_url_decode"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending URLDecode DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *urlDecodeDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode DataSource Schema method.")

	resp.Schema = schema.Schema{
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
		Attributes: map[string]schema.Attribute{
			"encoded_url": schema.StringAttribute{
				MarkdownDescription: "An encoded URL.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The decoded URL.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending URLDecode DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *urlDecodeDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending URLDecode DataSource Configure method.")
}

func (d *urlDecodeDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode DataSource Create method.")

	var plan urlDecodeDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending URLDecode DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *urlDecodeDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting URLDecode DataSource Read method.")

	var state urlDecodeDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	url, err := corefunc.URLDecode(
		state.EncodedURL.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to decode the URL",
			err.Error(),
		)

		return
	}

	state.Value = types.StringValue(url)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending URLDecode DataSource Read method.")
}
