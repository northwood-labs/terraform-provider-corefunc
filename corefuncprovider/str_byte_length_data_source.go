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
	_ datasource.DataSource              = &strByteLengthDataSource{}
	_ datasource.DataSourceWithConfigure = &strByteLengthDataSource{}
)

// strByteLengthDataSource is the data source implementation.
type (
	strByteLengthDataSource struct{}

	// strByteLengthDataSourceModel maps the data source schema data.
	strByteLengthDataSourceModel struct {
		String types.String `tfsdk:"string"`
		Value  types.Int64  `tfsdk:"value"`
	}
)

// StrByteLengthDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrByteLengthDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strByteLengthDataSource{}
}

// Metadata returns the data source type name.
func (d *strByteLengthDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrByteLength DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_byte_length"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending StrByteLength DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strByteLengthDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrByteLength DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Counts the number of bytes in a string.

		This is different from the length of a string in characters, as some
		characters may be represented by multiple bytes in UTF-8 encoding.

		Maps to the ` + linkPackage("StrByteLength") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				MarkdownDescription: "The input string to count bytes for.",
				Required:            true,
			},
			"value": schema.Int64Attribute{
				MarkdownDescription: "The number of bytes in the input string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrByteLength DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strByteLengthDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrByteLength DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrByteLength DataSource Configure method.")
}

func (d *strByteLengthDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrByteLength DataSource Create method.")

	var plan strByteLengthDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrByteLength DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strByteLengthDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrByteLength DataSource Read method.")

	var state strByteLengthDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value := corefunc.StrByteLength(
		state.String.ValueString(),
	)

	state.Value = types.Int64Value(int64(value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrByteLength DataSource Read method.")
}
