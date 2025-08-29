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

	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &intLeftpadDataSource{}
	_ datasource.DataSourceWithConfigure = &intLeftpadDataSource{}
)

// intLeftpadDataSource is the data source implementation.
type (
	intLeftpadDataSource struct{}

	// intLeftpadDataSourceModel maps the data source schema data.
	intLeftpadDataSourceModel struct {
		Value    types.String `tfsdk:"value"`
		Num      types.Int64  `tfsdk:"num"`
		PadWidth types.Int64  `tfsdk:"pad_width"`
	}
)

// IntLeftpadDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func IntLeftpadDataSource() datasource.DataSource { // lint:allow_return_interface
	return &intLeftpadDataSource{}
}

// Metadata returns the data source type name.
func (d *intLeftpadDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_int_leftpad"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending IntLeftpad DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *intLeftpadDataSource) Schema( // lint:no_dupe
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts an integer to a string, and then pads it with zeroes on the left.

		-> If the integer is NOT in base10 (decimal), it will be converted to base10 (decimal) _before_ being padded.

		Maps to the ` + linkPackage("IntLeftPad") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"num": schema.Int64Attribute{
				MarkdownDescription: "The integer to pad with zeroes.",
				Required:            true,
			},
			"pad_width": schema.Int64Attribute{
				MarkdownDescription: "The max number of zeroes to pad the integer with.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending IntLeftpad DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *intLeftpadDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending IntLeftpad DataSource Configure method.")
}

func (d *intLeftpadDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad DataSource Create method.")

	var plan intLeftpadDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending IntLeftpad DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *intLeftpadDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting IntLeftpad DataSource Read method.")

	var state intLeftpadDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		corefunc.IntLeftPad(
			state.Num.ValueInt64(),
			int(state.PadWidth.ValueInt64()),
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending IntLeftpad DataSource Read method.")
}
