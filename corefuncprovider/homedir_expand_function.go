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

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &homedirExpandFunction{}

type (
	// homedirExpandFunction is the function implementation.
	homedirExpandFunction struct{}
)

// HomedirExpandFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func HomedirExpandFunction() function.Function { // lint:allow_return_interface
	return &homedirExpandFunction{}
}

func (f *homedirExpandFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HomedirExpand Function Metadata method.")

	resp.Name = "homedir_expand"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending HomedirExpand Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *homedirExpandFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting HomedirExpand Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Replaces the ~ character in a path with the current user's home directory.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Replaces the ~ character in a path with the current user's home directory.

		Maps to the ` + linkPackage("HomedirExpand") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "path",
				MarkdownDescription: "The path to expand.",
			},
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending HomedirExpand Function Definition method.")
}

func (f *homedirExpandFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting HomedirExpand Function Run method.")

	var path string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &path))
	if resp.Error != nil {
		return
	}

	value, err := corefunc.HomedirExpand(path)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending HomedirExpand Function Run method.")
}
