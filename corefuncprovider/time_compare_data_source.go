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
	_ datasource.DataSource              = &timeCompareDataSource{}
	_ datasource.DataSourceWithConfigure = &timeCompareDataSource{}
)

// timeCompareDataSource is the data source implementation.
type (
	timeCompareDataSource struct{}

	// timeCompareDataSourceModel maps the data source schema data.
	timeCompareDataSourceModel struct {
		TimestampA types.String `tfsdk:"timestamp_a"`
		TimestampB types.String `tfsdk:"timestamp_b"`
		Value      types.Int64  `tfsdk:"value"`
	}
)

// TimeCompareDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func TimeCompareDataSource() datasource.DataSource { // lint:allow_return_interface
	return &timeCompareDataSource{}
}

// Metadata returns the data source type name.
func (d *timeCompareDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_time_compare"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending TimeCompare DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *timeCompareDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Compares two timestamps and returns a number which represents the
		ordering of those timestamps.

		When comparing timestamps, it takes into account the UTC offsets. For
		example, ` + code("06:00:00+0200") + ` and ` + code("04:00:00Z") + ` are
		the same moment after taking into account the ` + code("+0200") + `
		offset on the first timestamp.

		-> This is backwards compatible with the built-in ` + code("timecmp") + `
		function which was added in Terraform 1.3. While Terraform and OpenTofu
		both support the RFC3339 format for timestamps, this function supports a
		wider range of timestamp formats for compatibility with various systems.

		Maps to the ` + linkPackage("TimeCompare") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"timestamp_a": schema.StringAttribute{
				MarkdownDescription: "The first timestamp value to compare.",
				Required:            true,
			},
			"timestamp_b": schema.StringAttribute{
				MarkdownDescription: "The second timestamp value to compare.",
				Required:            true,
			},
			"value": schema.Int64Attribute{
				MarkdownDescription: "An integer indicating the comparison result. A value of `-1`" +
					" indicates that `timestamp_a` is less than `timestamp_b`, `0` indicates that" +
					" they are equal, and `1` indicates that `timestamp_a` is greater than `timestamp_b`.",
				Computed: true,
			},
		},
	}

	tflog.Debug(ctx, "Ending TimeCompare DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *timeCompareDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending TimeCompare DataSource Configure method.")
}

func (d *timeCompareDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare DataSource Create method.")

	var plan timeCompareDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TimeCompare DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *timeCompareDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare DataSource Read method.")

	var state timeCompareDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value, err := corefunc.TimeCompare(
		state.TimestampA.ValueString(),
		state.TimestampB.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Time comparison error",
			"An error occurred while comparing the time values: "+err.Error(),
		)

		return
	}

	state.Value = types.Int64Value(int64(value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TimeCompare DataSource Read method.")
}
