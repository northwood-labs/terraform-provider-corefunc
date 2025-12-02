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
	_ datasource.DataSource              = &strStartswithDataSource{}
	_ datasource.DataSourceWithConfigure = &strStartswithDataSource{}
)

// strStartswithDataSource is the data source implementation.
type (
	strStartswithDataSource struct{}

	// strStartswithDataSourceModel maps the data source schema data.
	strStartswithDataSourceModel struct {
		Input  types.String `tfsdk:"input"`
		Prefix types.String `tfsdk:"prefix"`
		Value  types.Bool   `tfsdk:"value"`
	}
)

// StrStartswithDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrStartswithDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strStartswithDataSource{}
}

// Metadata returns the data source type name.
func (d *strStartswithDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_startswith"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending StrStartswith DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strStartswithDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Takes a string to check and a prefix string, and returns true if the
		string begins with that exact prefix.

		-> This is identical to the built-in ` + "`startswith`" + ` function
		which was added in Terraform 1.3. This provides a 1:1 implementation
		that can be used with Terratest or other Go code, as well as with
		OpenTofu and Terraform going all the way back to v1.0.

		Maps to the ` + linkPackage("StrStartswith") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string to check.",
				Required:            true,
			},
			"prefix": schema.StringAttribute{
				MarkdownDescription: "The prefix to look for.",
				Required:            true,
			},
			"value": schema.BoolAttribute{
				MarkdownDescription: "The resulting boolean value as a string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrStartswith DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strStartswithDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrStartswith DataSource Configure method.")
}

func (d *strStartswithDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith DataSource Create method.")

	var plan strStartswithDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrStartswith DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strStartswithDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrStartswith DataSource Read method.")

	var state strStartswithDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value := corefunc.StrStartsWith(
		state.Input.ValueString(),
		state.Prefix.ValueString(),
	)

	state.Value = types.BoolValue(value)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrStartswith DataSource Read method.")
}
