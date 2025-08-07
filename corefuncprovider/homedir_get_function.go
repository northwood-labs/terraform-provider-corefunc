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
var _ function.Function = &homedirGetFunction{}

type (
	// homedirGetFunction is the function implementation.
	homedirGetFunction struct{}
)

// HomedirGetFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func HomedirGetFunction() function.Function { // lint:allow_return_interface
	return &homedirGetFunction{}
}

func (f *homedirGetFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet Function Metadata method.")

	resp.Name = "homedir_get"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending HomedirGet Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *homedirGetFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting HomedirGet Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Returns the path to the home directory of the current user.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the path to the home directory of the current user.

		Maps to the ` + linkPackage("Homedir") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending HomedirGet Function Definition method.")
}

func (f *homedirGetFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting HomedirGet Function Run method.")

	value, err := corefunc.Homedir()
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(err.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, value))

	tflog.Debug(ctx, "Ending HomedirGet Function Run method.")
}
