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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
	cftypes "github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &strIterativeReplaceFunction{}

type (
	// strIterativeReplaceFunction is the function implementation.
	strIterativeReplaceFunction struct{}
)

// StrIterativeReplaceFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func StrIterativeReplaceFunction() function.Function { // lint:allow_return_interface
	return &strIterativeReplaceFunction{}
}

func (f *strIterativeReplaceFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace Function Metadata method.")

	resp.Name = "str_iterative_replace"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending StrIterativeReplace Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *strIterativeReplaceFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting StrIterativeReplace Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Performs a series of replacements against a string.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Performs a series of replacements against a string. Allows a Terraform
        module to accept a set of replacements from a user.

        Maps to the [` + "`" + `corefunc.StrIterativeReplace()` + "`" + `](URL)
        Go method, which can be used in ` + Terratest + `.
        `)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "str",
				MarkdownDescription: "The string upon which replacements should be applied.",
			},
			function.ListParameter{
				Name: "replacements",
				MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
                A list of maps. Each map has an ` + "`" + `old` + "`" + ` and ` + "`" + `new` + "`" +
					` key. ` + "`" + `old` + "`" + ` represents the existing string to be replaced, and ` + "`" +
					`new` + "`" + ` represents the replacement string.
                `)),
				// https://discuss.hashicorp.com/t/provider-functions-equivalent-to-schema-listnestedattribute/63585
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending StrIterativeReplace Function Definition method.")
}

func (f *strIterativeReplaceFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting StrIterativeReplace Function Run method.")

	var (
		str          string
		replacements []map[string]string
	)

	err := req.Arguments.Get(ctx, &str, &replacements)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	// Remap!
	repls := make([]cftypes.Replacement, len(replacements))

	for i := range replacements {
		r := replacements[i]
		repls[i] = cftypes.Replacement{
			Old: r["old"],
			New: r["new"],
		}
	}

	value := corefunc.StrIterativeReplace(
		str,
		repls,
	)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending StrIterativeReplace Function Run method.")
}
