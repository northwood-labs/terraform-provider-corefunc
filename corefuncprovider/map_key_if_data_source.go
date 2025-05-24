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

package corefuncprovider // lint:no_dupe

import (
	"context"
	"fmt"
	"strings"

	"github.com/dcarbone/terraform-plugin-framework-utils/v3/conv"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/debug"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &mapKeyIfDataSource{}
	_ datasource.DataSourceWithConfigure = &mapKeyIfDataSource{}
)

// mapKeyIfDataSource is the data source implementation.
type (
	mapKeyIfDataSource struct{}

	// mapKeyIfDataSourceModel maps the data source schema data.
	mapKeyIfDataSourceModel struct {
		Condition types.Bool    `tfsdk:"condition"`
		Map       types.Dynamic `tfsdk:"map"`
		Key       types.String  `tfsdk:"key"`
		Value     types.Dynamic `tfsdk:"value"`
		MapCopy   types.Dynamic `tfsdk:"map_copy"`
	}
)

// MapKeyIfDataSource is a method that exposes its paired Go function as a
// Terraform Data Source.
func MapKeyIfDataSource() datasource.DataSource { // lint:allow_return_interface
	return &mapKeyIfDataSource{}
}

// Metadata returns the data source type name.
func (d *mapKeyIfDataSource) Metadata(
	ctx context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting MapKeyIf DataSource Metadata method.")

	resp.TypeName = req.ProviderTypeName + "_map_key_if"

	tflog.Debug(ctx, fmt.Sprintf("req.ProviderTypeName = %s", req.ProviderTypeName))
	tflog.Debug(ctx, fmt.Sprintf("resp.TypeName = %s", resp.TypeName))

	tflog.Debug(ctx, "Ending MapKeyIf DataSource Metadata method.")
}

// Schema defines the schema for the data source.
func (d *mapKeyIfDataSource) Schema(
	ctx context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	tflog.Debug(ctx, "Starting MapKeyIf DataSource Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		When the condition evaluates to ` + code("true") + `, the given key-value
		pair is added to the map and a copy is returned (the map is NOT updated
		in-place). When the condition is ` + code("false") + `, a copy of the
		original map is returned.

		~> A **map** is a set of key-value pairs where _all of the values are
		the **same type**_. An **object** is a set of key-value pairs where
		_values can be **different types**_. If you add a key-value pair to a
		map where the value is a _different type_ from the existing values, then
		the ` + code("map") + ` will be converted to an ` + code("object") + `
		automatically.

		For ` + Terratest + `, map keys should be updated normally.
		`)),
		Attributes: map[string]schema.Attribute{
			"condition": schema.BoolAttribute {
				MarkdownDescription: "The condition to evaluate. If the value is `true`, the key will be added to the map.",
				Required:    true,
			},
			"map": schema.DynamicAttribute {
				MarkdownDescription: "The single-level map/object variable to which the key should be added.",
				Required:    true,
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "The key to be added to the map/object.",
				Required:    true,
			},
			"value": schema.DynamicAttribute {
				MarkdownDescription: "The value to assign to the key.",
				Required:    true,
			},
			"map_copy": schema.DynamicAttribute {
				MarkdownDescription: "@TODO",
				Computed:    true,
			},
		},
	}

	tflog.Debug(ctx, "Ending MapKeyIf DataSource Schema method.")
}

// Configure adds the provider configured client to the data source.
func (d *mapKeyIfDataSource) Configure(
	ctx context.Context,
	req datasource.ConfigureRequest,
	_ *datasource.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting MapKeyIf DataSource Configure method.")

	if req.ProviderData == nil {
		return
	}

	tflog.Debug(ctx, "Ending MapKeyIf DataSource Configure method.")
}

func (d *mapKeyIfDataSource) Create(
	ctx context.Context,
	req resource.CreateRequest, // lint:allow_large_memory
	resp *resource.CreateResponse,
) {
	tflog.Debug(ctx, "Starting MapKeyIf DataSource Create method.")

	var plan mapKeyIfDataSourceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending MapKeyIf DataSource Create method.")
}

// Read refreshes the Terraform state with the latest data.
func (d *mapKeyIfDataSource) Read( // lint:no_dupe
	ctx context.Context,
	_ datasource.ReadRequest, // lint:allow_large_memory
	resp *datasource.ReadResponse,
) {
	tflog.Debug(ctx, "Starting MapKeyIf DataSource Read method.")

	var state mapKeyIfDataSourceModel
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	pp := debug.GetSpew()

	if !state.Condition.ValueBool() {
		state.MapCopy = state.Map
	} else {
		switch state.Map.UnderlyingValue().(type) {
		case types.Map:
			m := conv.ValueToMapType(state.Map.UnderlyingValue())
			goMap := m.Elements()

			pp.Dump(goMap)

			key := state.Key.ValueString()

			switch state.Value.UnderlyingValue().(type) {
			case types.Bool:
				goMap[key] = conv.ValueToBoolType(state.Value.UnderlyingValue())
			case types.Number:
				goMap[key] = conv.ValueToNumberType(state.Value.UnderlyingValue())
			case types.List:
				goMap[key] = conv.ValueToListType(state.Value.UnderlyingValue())
			case types.Map:
				goMap[key] = conv.ValueToMapType(state.Value.UnderlyingValue())
			case types.Object:
				goMap[key] = conv.ValueToObjectType(state.Value.UnderlyingValue())
			case types.Set:
				goMap[key] = conv.ValueToSetType(state.Value.UnderlyingValue())
			case types.String:
				goMap[key] = conv.ValueToStringType(state.Value.UnderlyingValue())
			case types.Tuple:
				goMap[key] = conv.ValueToListType(state.Value.UnderlyingValue())
			}

			pp.Dump(goMap)

			// t, diag := types.MapValueFrom(ctx, state.Map.ElementType(ctx), goMap)
			// resp.Diagnostics.Append(diag...)

			// state.MapCopy = t

			pp.Dump("MAP")
			pp.Dump(state)
		case types.Object:
			pp.Dump("OBJECT")
			pp.Dump(state)
		default:
			pp.Dump("DEFAULT")
			pp.Dump(state.Map.UnderlyingValue().Type(ctx))
			pp.Dump(state)
			// Throw error
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Ending MapKeyIf DataSource Read method.")
}
