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

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &strSnakeDataSource{}
	_ datasource.DataSourceWithConfigure = &strSnakeDataSource{}
)

// strSnakeDataSource is the data source implementation.
type (
	strSnakeDataSource struct{}

	// strSnakeDataSourceModel maps the data source schema data.
	strSnakeDataSourceModel struct {
		String types.String `tfsdk:"string"`
		Value  types.String `tfsdk:"value"`
	}
)

// StrSnakeDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrSnakeDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strSnakeDataSource{}
}

// Metadata returns the data source type name.
func (d *strSnakeDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_snake"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending StrSnake DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strSnakeDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Converts a string to ` + "`" + `snake_case` + "`" + `, removing any non-alphanumeric characters.

		Maps to the ` + linkPackage("StrSnake") + ` Go method, which can be used in ` + Terratest + `.
        `)),
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				MarkdownDescription: "The string to convert to `snake_case`.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrSnake DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strSnakeDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrSnake DataSource Configure method.")
}

func (d *strSnakeDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake DataSource Create method.")

	var plan strSnakeDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrSnake DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strSnakeDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrSnake DataSource Read method.")

	var state strSnakeDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		corefunc.StrSnake(
			state.String.ValueString(),
			// opts,
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrSnake DataSource Read method.")
}
