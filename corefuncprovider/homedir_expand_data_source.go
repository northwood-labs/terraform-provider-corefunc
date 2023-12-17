// Copyright 2023, Ryan Parman
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
	_ datasource.DataSource              = &homedirExpandDataSource{}
	_ datasource.DataSourceWithConfigure = &homedirExpandDataSource{}
)

// homedirExpandDataSource is the data source implementation.
type (
	homedirExpandDataSource struct{}

	// homedirExpandDataSourceModel maps the data source schema data.
	homedirExpandDataSourceModel struct {
		Path  types.String `tfsdk:"path"`
		Value types.String `tfsdk:"value"`
		ID    types.Int64  `tfsdk:"id"`
	}
)

// HomedirExpandDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func HomedirExpandDataSource() datasource.DataSource { // lint:allow_return_interface
	return &homedirExpandDataSource{}
}

// Metadata returns the data source type name.
func (d *homedirExpandDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Info(ctx, "Starting HomedirExpand DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_homedir_expand"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Info(ctx, "Ending HomedirExpand DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *homedirExpandDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Info(ctx, "Starting HomedirExpand DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Replaces the ~ character in a path with the current user's home directory.

		Maps to the ` + linkPackage("Homedir") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Not used. Required by the " + TPF + ".",
				Computed:    true,
			},
			"path": schema.StringAttribute{
				Description: "The path to expand.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "The path with the home directory expanded.",
				Computed:    true,
			},
		},
	}

	tflog.Info(ctx, "Ending HomedirExpand DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *homedirExpandDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Info(ctx, "Starting HomedirExpand DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "Ending HomedirExpand DataSource Configure method.")
}

func (d *homedirExpandDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Info(ctx, "Starting HomedirExpand DataSource Create method.")

	var plan homedirExpandDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending HomedirExpand DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *homedirExpandDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Info(ctx, "Starting HomedirExpand DataSource Read method.")

	var state homedirExpandDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.ID = types.Int64Value(1)

	homedir, err := corefunc.HomedirExpand(state.Path.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to resolve the path",
			err.Error(),
		)

		return
	}

	state.Value = types.StringValue(homedir)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending HomedirExpand DataSource Read method.")
}
