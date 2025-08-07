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

package corefuncprovider

import (
	"context"
	"fmt"
	"strings"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

const (
	defaultMaxLength   = 64
	validatorMaxLength = 255
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &truncateLabelDataSource{}
	_ datasource.DataSourceWithConfigure = &truncateLabelDataSource{}
)

// truncateLabelDataSource is the data source implementation.
type (
	truncateLabelDataSource struct{}

	// truncateLabelDataSourceModel maps the data source schema data.
	truncateLabelDataSourceModel struct {
		Label     types.String `tfsdk:"label"`
		Prefix    types.String `tfsdk:"prefix"`
		Value     types.String `tfsdk:"value"`
		MaxLength types.Int64  `tfsdk:"max_length"`
	}
)

// TruncateLabelDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func TruncateLabelDataSource() datasource.DataSource { // lint:allow_return_interface
	return &truncateLabelDataSource{}
}

// Metadata returns the data source type name.
func (d *truncateLabelDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting TruncateLabel DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_truncate_label"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending TruncateLabel DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *truncateLabelDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting TruncateLabel DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Supports prepending a prefix to a label, while truncating them
        to meet the maximum length constraints. Useful when grouping labels with a
        kind of prefix. Both the prefix and the label will be truncated if necessary.

        Uses a “balancing” algorithm between the prefix and the label, so that each
        section is truncated as a factor of how much space it takes up in the merged
        string.

        **DEPRECATED:** This function is deprecated and will be removed in the 2.0.0
        release. It's a bit too specialized to be useful in a general-purpose library.

        -> The motivation for this is in working with monitoring systems such
        as New Relic and Datadog where there are hundreds of applications in a
        monitoring “prod” account, and also hundreds of applications in a monitoring
        “nonprod” account. This allows us to group lists of monitors together using a
        shared prefix, but also truncate them appropriately to fit length
        constraints for names.

        Maps to the ` + linkPackage("TruncateLabel") + ` Go method, which can be used in
        ` + Terratest + `.
        `)),
		Attributes: map[string]schema.Attribute{
			"max_length": schema.Int64Attribute{
				MarkdownDescription: "The maximum allowed length of the combined label. " +
					"Minimum value is `1`. The default value is `64`.",
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"prefix": schema.StringAttribute{
				MarkdownDescription: "The prefix to prepend to the label.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, validatorMaxLength),
				},
			},
			"label": schema.StringAttribute{
				MarkdownDescription: "The label itself.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, validatorMaxLength),
				},
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The value of the truncated label.",
				Computed:            true,
			},
		},
	}

	tflog.Debug(ctx, "Ending TruncateLabel DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *truncateLabelDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting TruncateLabel DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending TruncateLabel DataSource Configure method.")
}

func (d *truncateLabelDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting TruncateLabel DataSource Create method.")

	var plan truncateLabelDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TruncateLabel DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *truncateLabelDataSource) Read(
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting TruncateLabel DataSource Read method.")

	var state truncateLabelDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	// Default values
	if state.MaxLength.ValueInt64() == 0 {
		state.MaxLength = types.Int64Value(defaultMaxLength)
	}

	state.Value = types.StringValue(
		corefunc.TruncateLabel(
			state.MaxLength.ValueInt64(),
			state.Prefix.ValueString(),
			state.Label.ValueString(),
		),
	)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending TruncateLabel DataSource Read method.")
}
