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
	_ datasource.DataSource              = &timeParseDataSource{}
	_ datasource.DataSourceWithConfigure = &timeParseDataSource{}
)

// timeParseDataSource is the data source implementation.
type (
	timeParseDataSource struct{}

	// timeParseDataSourceModel maps the data source schema data.
	timeParseDataSourceModel struct {
		Input types.String `tfsdk:"input"`
		Value types.Int64  `tfsdk:"value"`
	}
)

// TimeParseDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func TimeParseDataSource() datasource.DataSource { // lint:allow_return_interface
	return &timeParseDataSource{}
}

// Metadata returns the data source type name.
func (d *timeParseDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting TimeParse DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_time_parse"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending TimeParse DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *timeParseDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting TimeParse DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Parses a timestamp string and converts it into the number of seconds since Unix epoch.

		Supports a wide variety of formats: ISO8601, RFC822, RFC850, RFC1036,
		RFC1123, RFC3339, RFC5322, and similar.

		Maps to the ` + linkPackage("TimeParse") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string timestamp to parse",
				Required:            true,
			},
			"value": schema.Int64Attribute{
				MarkdownDescription: "The parsed timestamp as the number of seconds since Unix epoch.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending TimeParse DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *timeParseDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting TimeParse DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending TimeParse DataSource Configure method.")
}

func (d *timeParseDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting TimeParse DataSource Create method.")

	var plan timeParseDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TimeParse DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *timeParseDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting TimeParse DataSource Read method.")

	var state timeParseDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value, err := corefunc.TimeParse(
		state.Input.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error with parsing the timestamp",
			fmt.Sprintf(
				"An error was encountered parsing the time string %q: %s",
				state.Input.ValueString(),
				err.Error(),
			),
		)

		return
	}

	state.Value = types.Int64Value(value)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TimeParse DataSource Read method.")
}
