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
	_ datasource.DataSource              = &hashBcryptDataSource{}
	_ datasource.DataSourceWithConfigure = &hashBcryptDataSource{}
)

// hashBcryptDataSource is the data source implementation.
type (
	hashBcryptDataSource struct{}

	// hashBcryptDataSourceModel maps the data source schema data.
	hashBcryptDataSourceModel struct {
		Input types.String `tfsdk:"input"`
		Value types.String `tfsdk:"value"`
		Cost  types.Int32  `tfsdk:"cost"`
	}
)

// HashBcryptDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func HashBcryptDataSource() datasource.DataSource { // lint:allow_return_interface
	return &hashBcryptDataSource{}
}

// Metadata returns the data source type name.
func (d *hashBcryptDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashBcrypt DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_hash_bcrypt"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending HashBcrypt DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *hashBcryptDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting HashBcrypt DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Generates the Bcrypt hash of a string with its associated key value.

		Maps to the ` + linkPackage("HashBcrypt") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string to generate the Bcrypt hash for.",
				Required:            true,
			},
			"cost": schema.Int32Attribute{
				MarkdownDescription: "How much effort to spend generating a value. Common values are 10-14. " +
					"The default value is 10.",
				Optional: true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The result of the hashing function.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending HashBcrypt DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *hashBcryptDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting HashBcrypt DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending HashBcrypt DataSource Configure method.")
}

func (d *hashBcryptDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting HashBcrypt DataSource Create method.")

	var plan hashBcryptDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HashBcrypt DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *hashBcryptDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting HashBcrypt DataSource Read method.")

	var state hashBcryptDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value, err := corefunc.HashBcrypt(
		state.Input.ValueString(),
		int(state.Cost.ValueInt32()),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Problem with calculating the Bcrypt hash.",
			err.Error(),
		)

		return
	}

	state.Value = types.StringValue(value)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HashBcrypt DataSource Read method.")
}
