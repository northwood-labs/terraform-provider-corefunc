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
	_ datasource.DataSource              = &runtimeNumcpusDataSource{}
	_ datasource.DataSourceWithConfigure = &runtimeNumcpusDataSource{}
)

// runtimeNumcpusDataSource is the data source implementation.
type (
	runtimeNumcpusDataSource struct{}

	// runtimeNumcpusDataSourceModel maps the data source schema data.
	runtimeNumcpusDataSourceModel struct {
		Value types.Int64 `tfsdk:"value"`
	}
)

// RuntimeNumcpusDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func RuntimeNumcpusDataSource() datasource.DataSource { // lint:allow_return_interface
	return &runtimeNumcpusDataSource{}
}

// Metadata returns the data source type name.
func (d *runtimeNumcpusDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_runtime_numcpus"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending RuntimeNumcpus DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *runtimeNumcpusDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the number of CPU cores (logical CPUs) for the current system.

		-> This counts the number of CPU cores, which may differ from the number
		of physical CPUs.

		Maps to the [` + "`" + `runtime.NumCPU()` + "`" + `](https://pkg.go.dev/runtime#NumCPU)
		Go method, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"value": schema.Int64Attribute{
				Description: "The number of CPU cores (logical CPUs) for the current system.",
				Computed:    true,
			},
		},
	}

	tflog.Debug(ctx, "Ending RuntimeNumcpus DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *runtimeNumcpusDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeNumcpus DataSource Configure method.")
}

func (d *runtimeNumcpusDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus DataSource Create method.")

	var plan runtimeNumcpusDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeNumcpus DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *runtimeNumcpusDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus DataSource Read method.")

	var state runtimeNumcpusDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.Int64Value(
		int64(runtime.NumCPU()),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeNumcpus DataSource Read method.")
}
