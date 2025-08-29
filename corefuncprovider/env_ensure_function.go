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
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &envEnsureFunction{}

type (
	// envEnsureFunction is the function implementation.
	envEnsureFunction struct{}
)

// EnvEnsureFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func EnvEnsureFunction() function.Function { // lint:allow_return_interface
	return &envEnsureFunction{}
}

func (f *envEnsureFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure Function Metadata method.")

	resp.Name = "env_ensure"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending EnvEnsure Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *envEnsureFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting EnvEnsure Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Ensures that a given environment variable is set to a non-empty value.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Ensures that a given environment variable is set to a non-empty value.
        If the environment variable is unset or if it is set to an empty string,
        it will trigger a Terraform-level error.

        -> Not every Terraform provider checks to ensure that the environment variables it
        requires are properly set before performing work, leading to late-stage errors.
        This will force an error to occur early in the execution if the environment
        variable is not set, or if its value doesn't match the expected patttern.

        Maps to the ` + linkPackage("EnvEnsure") + ` Go method, which can be used in
        ` + Terratest + `.
        `)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "name",
				MarkdownDescription: "The name of the environment variable to check.",
			},
		},
		VariadicParameter: function.StringParameter{
			Name: "pattern",
			MarkdownDescription: "A valid Go ([re2](https://github.com/google/re2/wiki/Syntax)) " +
				"regular expression pattern.",
			AllowNullValue: true,
		},
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending EnvEnsure Function Definition method.")
}

func (f *envEnsureFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting EnvEnsure Function Run method.")

	var (
		name    string
		pattern []string
	)

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &name, &pattern))
	if resp.Error != nil {
		return
	}

	if len(pattern) > 0 && pattern[0] != "" {
		rePattern, err := regexp.Compile(pattern[0])
		if err != nil {
			resp.Error = function.ConcatFuncErrors(
				function.NewArgumentFuncError(1, err.Error()),
			)

			return
		}

		err = corefunc.EnvEnsure(name, rePattern)
		if err != nil {
			resp.Error = function.ConcatFuncErrors(
				function.NewFuncError(err.Error()),
			)

			return
		}
	} else {
		err := corefunc.EnvEnsure(name)
		if err != nil {
			resp.Error = function.ConcatFuncErrors(
				function.NewArgumentFuncError(0, err.Error()),
			)

			return
		}
	}

	value := os.Getenv(name)

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending EnvEnsure Function Run method.")
}
