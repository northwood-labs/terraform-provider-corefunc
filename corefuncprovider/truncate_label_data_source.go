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
		ID        types.Int64  `tfsdk:"id"`
		MaxLength types.Int64  `tfsdk:"max_length"`
		Label     types.String `tfsdk:"label"`
		Prefix    types.String `tfsdk:"prefix"`
		Value     types.String `tfsdk:"value"`
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
	tflog.Info(ctx, "Starting TruncateLabel DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_str_truncate_label"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Info(ctx, "Ending TruncateLabel DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *truncateLabelDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Info(ctx, "Starting TruncateLabel DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Supports prepending a prefix to a label, while truncating them
        to meet the maximum length constraints. Useful when grouping labels with a
        kind of prefix. Both the prefix and the label will be truncated if necessary.

        Uses a “balancing” algorithm between the prefix and the label, so that each
        section is truncated as a factor of how much space it takes up in the merged
        string.

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
			"id": schema.Int64Attribute{
				Description: "Not used. Required by the " + TPF + ".",
				Computed:    true,
			},
			"max_length": schema.Int64Attribute{
				Description: "The maximum allowed length of the combined label. " +
					"Minimum value is `1`. The default value is `64`.",
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"prefix": schema.StringAttribute{
				Description: "The prefix to prepend to the label.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, validatorMaxLength),
				},
			},
			"label": schema.StringAttribute{
				Description: "The label itself.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, validatorMaxLength),
				},
			},
			"value": schema.StringAttribute{
				Description: "The value of the truncated label.",
				Computed:    true,
			},
		},
	}

	tflog.Info(ctx, "Ending TruncateLabel DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *truncateLabelDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Info(ctx, "Starting TruncateLabel DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Info(ctx, "Ending TruncateLabel DataSource Configure method.")
}

func (d truncateLabelDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Info(ctx, "Starting TruncateLabel DataSource Create method.")

	var plan truncateLabelDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Ending TruncateLabel DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *truncateLabelDataSource) Read(
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Info(ctx, "Starting TruncateLabel DataSource Read method.")

	var state truncateLabelDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	state.ID = types.Int64Value(1)

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

	tflog.Info(ctx, "Ending TruncateLabel DataSource Read method.")
}
