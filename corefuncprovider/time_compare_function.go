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

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &timeCompareFunction{}

type (
	// timeCompareFunction is the function implementation.
	timeCompareFunction struct{}
)

// TimeCompareFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func TimeCompareFunction() function.Function { // lint:allow_return_interface
	return &timeCompareFunction{}
}

func (f *timeCompareFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare Function Metadata method.")

	resp.Name = "time_compare"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending TimeCompare Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *timeCompareFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting TimeCompare Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Compares two timestamps and returns a number which represents the ordering of those timestamps.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Compares two timestamps and returns a number which represents the
		ordering of those timestamps.

		When comparing timestamps, it takes into account the UTC offsets. For
		example, ` + code("06:00:00+0200") + ` and ` + code("04:00:00Z") + ` are
		the same moment after taking into account the ` + code("+0200") + `
		offset on the first timestamp.

		-> This is backwards compatible with the built-in ` + code("timecmp") + `
		function which was added in Terraform 1.3. While Terraform and OpenTofu
		both support the RFC3339 format for timestamps, this function supports a
		wider range of timestamp formats for compatibility with various systems.

		Maps to the ` + linkPackage("TimeCompare") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "timestamp_a",
				MarkdownDescription: "The first timestamp value to compare.",
			},
			function.StringParameter{
				Name:                "timestamp_b",
				MarkdownDescription: "The second timestamp value to compare.",
			},
		},
		Return: function.Int64Return{},
	}

	tflog.Debug(ctx, "Ending TimeCompare Function Definition method.")
}

func (f *timeCompareFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting TimeCompare Function Run method.")

	var (
		timestampA string
		timestampB string
	)

	err := req.Arguments.Get(ctx, &timestampA, &timestampB)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	value, e := corefunc.TimeCompare(timestampA, timestampB)
	if e != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(e.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending TimeCompare Function Run method.")
}
