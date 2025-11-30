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
	_ datasource.DataSource              = &hashArgon2idBase64DataSource{}
	_ datasource.DataSourceWithConfigure = &hashArgon2idBase64DataSource{}
)

// hashArgon2idBase64DataSource is the data source implementation.
type (
	hashArgon2idBase64DataSource struct{}

	// hashArgon2idBase64DataSourceModel maps the data source schema data.
	hashArgon2idBase64DataSourceModel struct {
		Input types.String `tfsdk:"input"`
		Salt  types.String `tfsdk:"salt"`
		Value types.String `tfsdk:"value"`
	}
)

// HashArgon2idBase64DataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func HashArgon2idBase64DataSource() datasource.DataSource { // lint:allow_return_interface
	return &hashArgon2idBase64DataSource{}
}

// Metadata returns the data source type name.
func (d *hashArgon2idBase64DataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashArgon2idBase64 DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_hash_argon2id_base64"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending HashArgon2idBase64 DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *hashArgon2idBase64DataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting HashArgon2idBase64 DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Generates the Argon2id hash of a string with its associated salt value. Returns a Base64
		value instead of a hexadecimal string.

		Maps to the ` + linkPackage("Base64HashArgon2id") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string to generate the Argon2id hash for.",
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

	tflog.Debug(ctx, "Ending HashArgon2idBase64 DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *hashArgon2idBase64DataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting HashArgon2idBase64 DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending HashArgon2idBase64 DataSource Configure method.")
}

func (d *hashArgon2idBase64DataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting HashArgon2idBase64 DataSource Create method.")

	var plan hashArgon2idBase64DataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HashArgon2idBase64 DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *hashArgon2idBase64DataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting HashArgon2idBase64 DataSource Read method.")

	var state hashArgon2idBase64DataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	value, err := corefunc.Base64HashArgon2id(
		state.Input.ValueString(),
		[]byte(state.Salt.ValueString()),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Problem with calculating the Argon2id hash.",
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

	tflog.Debug(ctx, "Ending HashArgon2idBase64 DataSource Read method.")
}
