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
	"runtime"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &runtimeGorootDataSource{}
	_ datasource.DataSourceWithConfigure = &runtimeGorootDataSource{}
)

// runtimeGorootDataSource is the data source implementation.
type (
	runtimeGorootDataSource struct{}

	// runtimeGorootDataSourceModel maps the data source schema data.
	runtimeGorootDataSourceModel struct {
		Value types.String `tfsdk:"value"`
	}
)

// RuntimeGorootDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func RuntimeGorootDataSource() datasource.DataSource { // lint:allow_return_interface
	return &runtimeGorootDataSource{}
}

// Metadata returns the data source type name.
func (d *runtimeGorootDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_runtime_goroot"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending RuntimeGoroot DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *runtimeGorootDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the GOROOT path for the current system, if one exists.

		Maps to the [` + "`" + `runtime.GOROOT()` + "`" + `](https://pkg.go.dev/runtime#GOROOT)
		Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "Returns the GOROOT path for the current system, if one exists.",
				Computed:    true,
			},
		},
	}

	tflog.Debug(ctx, "Ending RuntimeGoroot DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *runtimeGorootDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeGoroot DataSource Configure method.")
}

func (d *runtimeGorootDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot DataSource Create method.")

	var plan runtimeGorootDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeGoroot DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *runtimeGorootDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot DataSource Read method.")

	var state runtimeGorootDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		runtime.GOROOT(),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeGoroot DataSource Read method.")
}
