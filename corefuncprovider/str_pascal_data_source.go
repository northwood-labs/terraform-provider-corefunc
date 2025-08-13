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

	"github.com/chanced/caps"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &strPascalDataSource{}
	_ datasource.DataSourceWithConfigure = &strPascalDataSource{}
)

// strPascalDataSource is the data source implementation.
type (
	strPascalDataSource struct{}

	// strPascalDataSourceModel maps the data source schema data.
	strPascalDataSourceModel struct {
		String      types.String `tfsdk:"string"`
		Value       types.String `tfsdk:"value"`
		AcronymCaps types.Bool   `tfsdk:"acronym_caps"`
	}
)

// StrPascalDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrPascalDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strPascalDataSource{}
}

// Metadata returns the data source type name.
func (d *strPascalDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_pascal"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending StrPascal DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strPascalDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `PascalCase` + "`" + `, removing any non-alphanumeric characters.

		-> Some acronyms are maintained as uppercase. See
		[caps: pkg-variables](https://pkg.go.dev/github.com/chanced/caps#pkg-variables) for a complete list.

		Maps to the ` + linkPackage("StrPascal") + ` Go method, which can be used in ` + Terratest + `.
        `)),
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				MarkdownDescription: "The string to convert to `PascalCase`.",
				Required:            true,
			},
			"acronym_caps": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to keep acronyms as uppercase. A value of `true` means that acronyms " +
					"will be converted to uppercase. A value of `false` means that acronyms will using typical " +
					"casing. The default value is `false`.",
				Optional: true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrPascal DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strPascalDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrPascal DataSource Configure method.")
}

func (d *strPascalDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal DataSource Create method.")

	var plan strPascalDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrPascal DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strPascalDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrPascal DataSource Read method.")

	var state strPascalDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	// Default values
	opts := caps.Opts{}

	if !state.AcronymCaps.ValueBool() {
		opts = caps.WithReplaceStyleCamel()
	}

	state.Value = types.StringValue(
		caps.ToCamel(
			state.String.ValueString(),
			opts,
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrPascal DataSource Read method.")
}
