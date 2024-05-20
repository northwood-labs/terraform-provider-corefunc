// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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
	_ datasource.DataSource              = &strLeftpadDataSource{}
	_ datasource.DataSourceWithConfigure = &strLeftpadDataSource{}
)

// strLeftpadDataSource is the data source implementation.
type (
	strLeftpadDataSource struct{}

	// strLeftpadDataSourceModel maps the data source schema data.
	strLeftpadDataSourceModel struct {
		PadChar  types.String `tfsdk:"pad_char"`
		String   types.String `tfsdk:"string"`
		Value    types.String `tfsdk:"value"`
		PadWidth types.Int64  `tfsdk:"pad_width"`
	}
)

// StrLeftpadDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrLeftpadDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strLeftpadDataSource{}
}

// Metadata returns the data source type name.
func (d *strLeftpadDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_leftpad"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending StrLeftpad DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strLeftpadDataSource) Schema( // lint:no_dupe
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Pads a string with additional characters on the left.

		Maps to the ` + linkPackage("StrLeftPad") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"pad_width": schema.Int64Attribute{
				MarkdownDescription: "The max number of padding characters to pad the string with.",
				Required:            true,
			},
			"string": schema.StringAttribute{
				MarkdownDescription: "The string to pad with padding characters.",
				Required:            true,
			},
			"pad_char": schema.StringAttribute{
				MarkdownDescription: "The padding character to use. Only supports a single byte. If more than one byte is provided, only the first byte will be used. The default value is a space character.", // lint:ignore_length
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrLeftpad DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strLeftpadDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrLeftpad DataSource Configure method.")
}

func (d *strLeftpadDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad DataSource Create method.")

	var plan strLeftpadDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrLeftpad DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strLeftpadDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrLeftpad DataSource Read method.")

	var state strLeftpadDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	// Grab the pad character, or use a space by default.
	padChar := state.PadChar.ValueString()
	if padChar == "" {
		state.Value = types.StringValue(
			corefunc.StrLeftPad(
				state.String.ValueString(),
				int(state.PadWidth.ValueInt64()),
			),
		)
	} else {
		padByte := []byte(padChar)[0]

		state.Value = types.StringValue(
			corefunc.StrLeftPad(
				state.String.ValueString(),
				int(state.PadWidth.ValueInt64()),
				padByte,
			),
		)
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrLeftpad DataSource Read method.")
}
