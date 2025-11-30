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
	_ datasource.DataSource              = &hashSha384DataSource{}
	_ datasource.DataSourceWithConfigure = &hashSha384DataSource{}
)

// hashSha384DataSource is the data source implementation.
type (
	hashSha384DataSource struct{}

	// hashSha384DataSourceModel maps the data source schema data.
	hashSha384DataSourceModel struct {
		Input types.String `tfsdk:"input"`
		Value types.String `tfsdk:"value"`
	}
)

// HashSha384DataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func HashSha384DataSource() datasource.DataSource { // lint:allow_return_interface
	return &hashSha384DataSource{}
}

// Metadata returns the data source type name.
func (d *hashSha384DataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_hash_sha384"

	tflog.Debug(ctx, "req.ProviderTypeName = "+req.ProviderTypeName)
	tflog.Debug(ctx, "resp.TypeName = "+resp.TypeName)

	tflog.Debug(ctx, "Ending HashSha384 DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *hashSha384DataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Generates the SHA-2/384-bit hash of a string.

		Maps to the ` + linkPackage("HashSHA384") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				MarkdownDescription: "The string to generate a SHA-2/384-bit hash for.",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The result of the hashing function.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending HashSha384 DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *hashSha384DataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending HashSha384 DataSource Configure method.")
}

func (d *hashSha384DataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 DataSource Create method.")

	var plan hashSha384DataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HashSha384 DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *hashSha384DataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting HashSha384 DataSource Read method.")

	var state hashSha384DataSourceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		corefunc.HashSHA384(
			state.Input.ValueString(),
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending HashSha384 DataSource Read method.")
}
