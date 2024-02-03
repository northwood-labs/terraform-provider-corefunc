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
	_ datasource.DataSource              = &runtimeCpuarchDataSource{}
	_ datasource.DataSourceWithConfigure = &runtimeCpuarchDataSource{}
)

// runtimeCpuarchDataSource is the data source implementation.
type (
	runtimeCpuarchDataSource struct{}

	// runtimeCpuarchDataSourceModel maps the data source schema data.
	runtimeCpuarchDataSourceModel struct {
		Value types.String `tfsdk:"value"`
	}
)

// RuntimeCpuarchDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func RuntimeCpuarchDataSource() datasource.DataSource { // lint:allow_return_interface
	return &runtimeCpuarchDataSource{}
}

// Metadata returns the data source type name.
func (d *runtimeCpuarchDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_runtime_cpuarch"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending RuntimeCpuarch DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *runtimeCpuarchDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the CPU architecture for which the provider was compiled.
		Generally, this represents the current system.

		-> If you are running the provider via Rosetta on an Apple Silicon
		machine, using QEMU, or similar CPU emulation software, this will return
		the architecture of the emulated CPU, not the host CPU.

		Maps to the [` + "`" + `runtime.GOARCH` + "`" + `](https://pkg.go.dev/runtime#GOARCH)
		Go property, which can be used in ` + Terratest + `.
		`)),
		Attributes: map[string]schema.Attribute{
			"value": schema.StringAttribute{
				Description: "Returns the CPU architecture for which the provider was compiled.",
				Computed:    true,
			},
		},
	}

	tflog.Debug(ctx, "Ending RuntimeCpuarch DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *runtimeCpuarchDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeCpuarch DataSource Configure method.")
}

func (d *runtimeCpuarchDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch DataSource Create method.")

	var plan runtimeCpuarchDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeCpuarch DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *runtimeCpuarchDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch DataSource Read method.")

	var state runtimeCpuarchDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		runtime.GOARCH,
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending RuntimeCpuarch DataSource Read method.")
}