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

package corefuncprovider

import (
	"context"
	"fmt"
	"os"
	"regexp"
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
	_ datasource.DataSource              = &envEnsureDataSource{}
	_ datasource.DataSourceWithConfigure = &envEnsureDataSource{}
)

// envEnsureDataSource is the data source implementation.
type (
	envEnsureDataSource struct{}

	// envEnsureDataSourceModel maps the data source schema data.
	envEnsureDataSourceModel struct {
		Name    types.String `tfsdk:"name"`
		Pattern types.String `tfsdk:"pattern"`
		Value   types.String `tfsdk:"value"`
	}
)

// EnvEnsureDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func EnvEnsureDataSource() datasource.DataSource { // lint:allow_return_interface
	return &envEnsureDataSource{}
}

// Metadata returns the data source type name.
func (d *envEnsureDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_env_ensure"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending EnvEnsure DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *envEnsureDataSource) Schema( // lint:no_dupe
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Ensures that a given environment variable is set to a non-empty value.
        If the environment variable is unset or if it is set to an empty string,
        it will trigger a Terraform-level error.

        -> Not every Terraform provider checks to ensure that the environment variables it
        requires are properly set before performing work, leading to late-stage errors.
        This will force an error to occur early in the execution if the environment
        variable is not set, or if its value doesn't match the expected patttern.

        Maps to the ` + linkPackage("EnvEnsure") + ` Go method, which can be used in
        ` + Terratest + `.
        `)),
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the environment variable to check.",
				Required:            true,
			},
			"pattern": schema.StringAttribute{
				MarkdownDescription: "A valid Go ([re2](https://github.com/google/re2/wiki/Syntax)) regular expression pattern.",
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the environment variable, if it exists.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending EnvEnsure DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *envEnsureDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending EnvEnsure DataSource Configure method.")
}

func (d *envEnsureDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure DataSource Create method.")

	var plan envEnsureDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending EnvEnsure DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *envEnsureDataSource) Read(
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure DataSource Read method.")

	var state envEnsureDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	err := corefunc.EnvEnsure(
		state.Name.ValueString(),
		regexp.MustCompile(state.Pattern.ValueString()),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Problem with Environment Variable",
			err.Error(),
		)

		return
	}

	// Slightly differ from the Go function, which returns true if the
	// environment variable is defined. In this case, we want to return the
	// value of the environment variable.
	state.Value = types.StringValue(
		os.Getenv(
			state.Name.ValueString(),
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending EnvEnsure DataSource Read method.")
}
