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
	_ datasource.DataSource              = &strCamelDataSource{}
	_ datasource.DataSourceWithConfigure = &strCamelDataSource{}
)

// strCamelDataSource is the data source implementation.
type (
	strCamelDataSource struct{}

	// strCamelDataSourceModel maps the data source schema data.
	strCamelDataSourceModel struct {
		// AcronymCaps types.Bool   `tfsdk:"acronym_caps"`
		String types.String `tfsdk:"string"`
		Value  types.String `tfsdk:"value"`
	}
)

// StrCamelDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrCamelDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strCamelDataSource{}
}

// Metadata returns the data source type name.
func (d *strCamelDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_camel"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending StrCamel DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strCamelDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel DataSource Schema method.")

	const desc = `, removing any non-alphanumeric characters.

		-> Some acronyms are maintained as uppercase. See
		[caps: pkg-variables](https://pkg.go.dev/github.com/chanced/caps#pkg-variables) for a complete list.

		Maps to the `

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `camelCase` + "`" + desc + linkPackage("StrCamel") +
			` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				MarkdownDescription: "The string to convert to `camelCase`.",
				Required:            true,
			},
			// "acronym_caps": schema.BoolAttribute{
			// 	MarkdownDescription: "Whether or not to keep acronyms as uppercase. A value of `true` means that acronyms " +
			// 		"will be converted to uppercase. A value of `false` means that acronyms will using typical " +
			// 		"casing. The default value is `false`.",
			// 	Optional: true,
			// },
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrCamel DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strCamelDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrCamel DataSource Configure method.")
}

func (d *strCamelDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel DataSource Create method.")

	var plan strCamelDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrCamel DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strCamelDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrCamel DataSource Read method.")

	var state strCamelDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	// Default values
	// opts := caps.Opts{}
	// opts := caps.WithReplaceStyleCamel()

	// if state.AcronymCaps.ValueBool() { // lint:allow_commented
	// 	opts = caps.WithReplaceStyleScreaming()
	// }

	state.Value = types.StringValue(
		corefunc.StrCamel(
			state.String.ValueString(),
			// opts,
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrCamel DataSource Read method.")
}
