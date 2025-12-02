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
	_ datasource.DataSource              = &strContainsDataSource{}
	_ datasource.DataSourceWithConfigure = &strContainsDataSource{}
)

// strContainsDataSource is the data source implementation.
type (
	strContainsDataSource struct{}

	// strContainsDataSourceModel maps the data source schema data.
	strContainsDataSourceModel struct {
		Input  types.String `tfsdk:"input"`
		Substr types.String `tfsdk:"substr"`
		Value  types.Bool   `tfsdk:"value"`
	}
)

// StrContainsDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrContainsDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strContainsDataSource{}
}

// Metadata returns the data source type name.
func (d *strContainsDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrContains DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_contains"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending StrContains DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strContainsDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrContains DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Checks if a string contains a given substring.

		-> This is identical to the built-in ` + "`strcontains`" + ` function
		which was added in Terraform 1.3. This provides a 1:1 implementation
		that can be used with Terratest or other Go code, as well as with
		OpenTofu and Terraform going all the way back to v1.0.

		Maps to the ` + linkPackage("StrContains") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string to check.",
				Required:            true,
			},
			"substr": schema.StringAttribute{
				MarkdownDescription: "The substring to look for.",
				Required:            true,
			},
			"value": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the string contains the specified substring. " +
					"A value of `true` indicates that it does. A value of `false` indicates that it does not.",
				Computed: true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrContains DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strContainsDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrContains DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrContains DataSource Configure method.")
}

func (d *strContainsDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrContains DataSource Create method.")

	var plan strContainsDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrContains DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strContainsDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrContains DataSource Read method.")

	var state strContainsDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value := corefunc.StrContains(
		state.Input.ValueString(),
		state.Substr.ValueString(),
	)

	state.Value = types.BoolValue(value)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrContains DataSource Read method.")
}
