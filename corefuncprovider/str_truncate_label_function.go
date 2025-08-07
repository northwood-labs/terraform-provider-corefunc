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
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &strTruncateLabelFunction{}

type (
	// strTruncateLabelFunction is the function implementation.
	strTruncateLabelFunction struct{}
)

// StrTruncateLabelFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrTruncateLabelFunction() function.Function { // lint:allow_return_interface
	return &strTruncateLabelFunction{}
}

func (f *strTruncateLabelFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrTruncateLabel Function Metadata method.")

	resp.Name = "str_truncate_label"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending StrTruncateLabel Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strTruncateLabelFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrTruncateLabel Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Supports prepending a prefix to a label, while truncating them to meet the maximum " +
			"length constraints.",
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
		Parameters: []function.Parameter{
			function.Int64Parameter{
				Name:                "max_length",
				MarkdownDescription: "The maximum allowed length of the combined label. Minimum value is `1`.",
			},
			function.StringParameter{
				Name:                "prefix",
				MarkdownDescription: "The prefix to prepend to the label.",
			},
			function.StringParameter{
				Name:                "label",
				MarkdownDescription: "The label itself.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrTruncateLabel Function Definition method.")
}

func (f *strTruncateLabelFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrTruncateLabel Function Run method.")

	var (
		maxLength int64
		prefix    string
		label     string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &maxLength, &prefix, &label))
	if resp.Error != nil {
		return
	}

	if maxLength < 1 {
		resp.Error = function.ConcatFuncErrors(
			function.NewArgumentFuncError(0, "max_length must be at least 1"),
		)

		return
	}

	value := corefunc.TruncateLabel(maxLength, prefix, label)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrTruncateLabel Function Run method.")
}
