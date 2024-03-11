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
	cfTypes "github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &strIterativeReplaceDataSource{}
	_ datasource.DataSourceWithConfigure = &strIterativeReplaceDataSource{}
)

type (
	// strIterativeReplaceDataSource is the data source implementation.
	strIterativeReplaceDataSource struct{}

	// strIterativeReplaceDataSourceModel maps the data source schema data.
	strIterativeReplaceDataSourceModel struct {
		String       types.String          `tfsdk:"string"`
		Value        types.String          `tfsdk:"value"`
		Replacements []cfTypes.Replacement `tfsdk:"replacements"`
	}
)

// StrIterativeReplaceDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func StrIterativeReplaceDataSource() datasource.DataSource { // lint:allow_return_interface
	return &strIterativeReplaceDataSource{}
}

// Metadata returns the data source type name.
func (d *strIterativeReplaceDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_iterative_replace"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending StrIterativeReplace DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *strIterativeReplaceDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Performs a series of replacements against a string. Allows a Terraform
        module to accept a set of replacements from a user.

        Maps to the [` + "`" + `corefunc.StrIterativeReplace()` + "`" + `](URL)
        Go method, which can be used in ` + Terratest + `.
        `)),
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				MarkdownDescription: "The string upon which replacements should be applied.",
				Required:            true,
			},
			"replacements": schema.ListNestedAttribute{
				MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
                A list of maps. Each map has an ` + "`" + `old` + "`" + ` and ` + "`" + `new` + "`" +
					` key. ` + "`" + `old` + "`" + ` represents the existing string to be replaced, and ` + "`" +
					`new` + "`" + ` represents the replacement string.
                `)),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"old": schema.StringAttribute{
							MarkdownDescription: "The substring to be replaced.",
							Required:            true,
						},
						"new": schema.StringAttribute{
							MarkdownDescription: "The replacement substring.",
							Required:            true,
						},
					},
				},
				Required: true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the string.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending StrIterativeReplace DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *strIterativeReplaceDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending StrIterativeReplace DataSource Configure method.")
}

func (d *strIterativeReplaceDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace DataSource Create method.")

	var plan strIterativeReplaceDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrIterativeReplace DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *strIterativeReplaceDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace DataSource Read method.")

	var state strIterativeReplaceDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.Value = types.StringValue(
		corefunc.StrIterativeReplace(
			state.String.ValueString(),
			state.Replacements,
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending StrIterativeReplace DataSource Read method.")
}
