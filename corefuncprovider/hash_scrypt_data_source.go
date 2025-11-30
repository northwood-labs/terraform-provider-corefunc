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
	_ datasource.DataSource              = &hashScryptDataSource{}
	_ datasource.DataSourceWithConfigure = &hashScryptDataSource{}
)

// hashScryptDataSource is the data source implementation.
type (
	hashScryptDataSource struct{}

	// hashScryptDataSourceModel maps the data source schema data.
	hashScryptDataSourceModel struct {
		Input types.String `tfsdk:"input"`
		Salt  types.String `tfsdk:"salt"`
		Value types.String `tfsdk:"value"`
	}
)

// HashScryptDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func HashScryptDataSource() datasource.DataSource { // lint:allow_return_interface
	return &hashScryptDataSource{}
}

// Metadata returns the data source type name.
func (d *hashScryptDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashScrypt DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_hash_scrypt"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending HashScrypt DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *hashScryptDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting HashScrypt DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Generates the Scrypt hash of a string with its associated salt value.

		Maps to the ` + linkPackage("HashScrypt") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string to generate the Scrypt hash for.",
				Required:            true,
			},
			"salt": schema.StringAttribute{
				MarkdownDescription: "A random value to provide additional entropy in the calculation.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The result of the hashing function.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending HashScrypt DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *hashScryptDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting HashScrypt DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending HashScrypt DataSource Configure method.")
}

func (d *hashScryptDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting HashScrypt DataSource Create method.")

	var plan hashScryptDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HashScrypt DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *hashScryptDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting HashScrypt DataSource Read method.")

	var state hashScryptDataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value, err := corefunc.HashScrypt(
		state.Input.ValueString(),
		[]byte(state.Salt.ValueString()),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Problem with calculating the Scrypt hash.",
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

	tflog.Debug(ctx, "Ending HashScrypt DataSource Read method.")
}
