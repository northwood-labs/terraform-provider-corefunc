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
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &tomlToJSONDataSource{}
	_ datasource.DataSourceWithConfigure = &tomlToJSONDataSource{}
)

// tomlToJSONDataSource is the data source implementation.
type (
	tomlToJSONDataSource struct{}

	// tomlToJSONDataSourceModel maps the data source schema data.
	tomlToJSONDataSourceModel struct {
		TOML types.String `tfsdk:"toml"`
		JSON types.String `tfsdk:"json"`
	}
)

// TomlToJSONDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func TomlToJSONDataSource() datasource.DataSource { // lint:allow_return_interface
	return &tomlToJSONDataSource{}
}

// Metadata returns the data source type name.
func (d *tomlToJSONDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_toml_to_json"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending TomlToJSON DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *tomlToJSONDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a TOML string into an equivalent JSON string.

		Maps to the ` + linkPackage("TOMLtoJSON") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"toml": schema.StringAttribute{
				MarkdownDescription: "The TOML string to convert to JSON.",
				Required:            true,
			},
			"json": schema.StringAttribute{
				MarkdownDescription: "The equivalent JSON value.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending TomlToJSON DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *tomlToJSONDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending TomlToJSON DataSource Configure method.")
}

func (d *tomlToJSONDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON DataSource Create method.")

	var plan tomlToJSONDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TomlToJSON DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *tomlToJSONDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting TomlToJSON DataSource Read method.")

	var state tomlToJSONDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	jsonVal, err := corefunc.TOMLtoJSON(state.TOML.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Converting TOML to JSON",
			fmt.Sprintf("Failed to convert TOML to JSON: %s", err),
		)

		return
	}

	state.JSON = types.StringValue(strings.TrimSpace(jsonVal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TomlToJSON DataSource Read method.")
}
