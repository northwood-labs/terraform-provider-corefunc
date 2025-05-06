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
	_ datasource.DataSource              = &netCidrContainsDataSource{}
	_ datasource.DataSourceWithConfigure = &netCidrContainsDataSource{}
)

// netCidrContainsDataSource is the data source implementation.
type (
	netCidrContainsDataSource struct{}

	// netCidrContainsDataSourceModel maps the data source schema data.
	netCidrContainsDataSourceModel struct {
		ContainingCidr    types.String `tfsdk:"containing_cidr"`
		ContainedIPOrCidr types.String `tfsdk:"contained_ip_or_cidr"`
		Value             types.Bool   `tfsdk:"value"`
	}
)

// NetCidrContainsDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func NetCidrContainsDataSource() datasource.DataSource { // lint:allow_return_interface
	return &netCidrContainsDataSource{}
}

// Metadata returns the data source type name.
func (d *netCidrContainsDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_net_cidr_contains"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending NetCidrContains DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *netCidrContainsDataSource) Schema( // lint:no_dupe
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		CIDRContains determines whether or not a given IP address, or an address prefix
		given in CIDR notation, is within a given IP network address prefix.

		Both arguments must belong to the same address family, either IPv4 or IPv6. A
		family mismatch will result in an error.

		-> This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
		This also provides a 1:1 implementation that can be used with Terratest or other
		Go code.

		Maps to the ` + linkPackage("CIDRContains") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"containing_cidr": schema.StringAttribute{
				MarkdownDescription: "A CIDR range to check as a containing range.",
				Required:            true,
			},
			"contained_ip_or_cidr": schema.StringAttribute{
				MarkdownDescription: "An IP address or CIDR range to check as a contained range.",
				Required:            true,
			},
			"value": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the IP address or CIDR range is contained within the container CIDR range.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending NetCidrContains DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *netCidrContainsDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending NetCidrContains DataSource Configure method.")
}

func (d *netCidrContainsDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains DataSource Create method.")

	var plan netCidrContainsDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending NetCidrContains DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *netCidrContainsDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains DataSource Read method.")

	var state netCidrContainsDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	isContained, err := corefunc.CIDRContains(
		state.ContainingCidr.ValueString(),
		state.ContainedIPOrCidr.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Problem with CIDR",
			err.Error(),
		)

		return
	}

	state.Value = types.BoolValue(isContained)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending NetCidrContains DataSource Read method.")
}
