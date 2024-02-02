// Copyright 2023-2024, Ryan Parman
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
	_ datasource.DataSource              = &strKebabDataSource{}
	_ datasource.DataSourceWithConfigure = &strKebabDataSource{}
)

// strKebabDataSource is the data source implementation.
type (
	strKebabDataSource struct{}

	// strKebabDataSourceModel maps the data source schema data.
	strKebabDataSourceModel struct {
		String types.String `tfsdk:"string"`
		Value  types.String `tfsdk:"value"`
	}
)

// StrKebabDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrKebabDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strKebabDataSource{}
}

// Metadata returns the data source type name.
func (d *strKebabDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Info(ctx, "Starting StrKebab DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_kebab"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Info(ctx, "Ending StrKebab DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strKebabDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Info(ctx, "Starting StrKebab DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `kebab-case` + "`" + `, removing any non-alphanumeric characters.

		Maps to the ` + linkPackage("StrKebab") + ` Go method, which can be used in ` + Terratest + `.
        `)),
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				Description: "The string to convert to `kebab-case`.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "The value of the string.",
				Computed:    true,
			},
		},
	}

	tflog.Info(ctx, "Ending StrKebab DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strKebabDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Info(ctx, "Starting StrKebab DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "Ending StrKebab DataSource Configure method.")
}

func (d *strKebabDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Info(ctx, "Starting StrKebab DataSource Create method.")

	var plan strKebabDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending StrKebab DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strKebabDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Info(ctx, "Starting StrKebab DataSource Read method.")

	var state strKebabDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		corefunc.StrKebab(
			state.String.ValueString(),
			// opts,
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending StrKebab DataSource Read method.")
}
