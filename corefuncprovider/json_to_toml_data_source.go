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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &jsonToTomlDataSource{}
	_ datasource.DataSourceWithConfigure = &jsonToTomlDataSource{}
)

// jsonToTomlDataSource is the data source implementation.
type (
	jsonToTomlDataSource struct{}

	// jsonToTomlDataSourceModel maps the data source schema data.
	jsonToTomlDataSourceModel struct {
		JSON types.String `tfsdk:"json"`
		TOML types.String `tfsdk:"toml"`
	}
)

// JSONToTomlDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func JSONToTomlDataSource() datasource.DataSource { // lint:allow_return_interface
	return &jsonToTomlDataSource{}
}

// Metadata returns the data source type name.
func (d *jsonToTomlDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_json_to_toml"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending JSONToToml DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *jsonToTomlDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a JSON string into an equivalent TOML string.

		Maps to the ` + linkPackage("JSONtoTOML") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"json": schema.StringAttribute{
				MarkdownDescription: "The JSON string to convert to TOML.",
				Required:            true,
			},
			"toml": schema.StringAttribute{
				MarkdownDescription: "The equivalent TOML value.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending JSONToToml DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *jsonToTomlDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending JSONToToml DataSource Configure method.")
}

func (d *jsonToTomlDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml DataSource Create method.")

	var plan jsonToTomlDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending JSONToToml DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *jsonToTomlDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting JSONToToml DataSource Read method.")

	var state jsonToTomlDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	tomlVal, err := corefunc.JSONtoTOML(state.JSON.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Converting JSON to TOML",
			fmt.Sprintf("Failed to convert JSON to TOML: %s", err),
		)

		return
	}

	state.TOML = types.StringValue(strings.TrimSpace(tomlVal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending JSONToToml DataSource Read method.")
}
