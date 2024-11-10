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
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &strBase64GunzipDataSource{}
	_ datasource.DataSourceWithConfigure = &strBase64GunzipDataSource{}
)

// strBase64GunzipDataSource is the data source implementation.
type (
	strBase64GunzipDataSource struct{}

	// strBase64GunzipDataSourceModel maps the data source schema data.
	strBase64GunzipDataSourceModel struct {
		GzippedBase64 types.String `tfsdk:"gzipped_base64"`
		Value         types.String `tfsdk:"value"`
	}
)

// StrBase64GunzipDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrBase64GunzipDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strBase64GunzipDataSource{}
}

// Metadata returns the data source type name.
func (d *strBase64GunzipDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_base64_gunzip"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending StrBase64Gunzip DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strBase64GunzipDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Base64Gunzip is a function that decodes a Base64-encoded string, then
		decompresses the result with gzip. Supports both padded and non-padded Base64
		strings.

		-> This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
		This also provides a 1:1 implementation that can be used with Terratest or other
		Go code.

		Maps to the ` + linkPackage("Base64Gunzip") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"gzipped_base64": schema.StringAttribute{
				MarkdownDescription: "A string of gzipped then Base64-encoded data.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The Base64-decoded, then un-gzipped data.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrBase64Gunzip DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strBase64GunzipDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrBase64Gunzip DataSource Configure method.")
}

func (d *strBase64GunzipDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip DataSource Create method.")

	var plan strBase64GunzipDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrBase64Gunzip DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strBase64GunzipDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrBase64Gunzip DataSource Read method.")

	var state strBase64GunzipDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	gunzipped, err := corefunc.Base64Gunzip(
		state.GzippedBase64.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Problem with Base64 decoding and Gzip decompression",
			err.Error(),
		)

		return
	}

	state.Value = types.StringValue(gunzipped)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrBase64Gunzip DataSource Read method.")
}
