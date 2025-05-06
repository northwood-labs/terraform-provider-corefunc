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
	"runtime"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &runtimeGorootFunction{}

type (
	// runtimeGorootFunction is the function implementation.
	runtimeGorootFunction struct{}
)

// RuntimeGorootFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func RuntimeGorootFunction() function.Function { // lint:allow_return_interface
	return &runtimeGorootFunction{}
}

func (f *runtimeGorootFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot Function Metadata method.")

	resp.Name = "runtime_goroot"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending RuntimeGoroot Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *runtimeGorootFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeGoroot Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Returns the GOROOT path for the current system, if one exists.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the GOROOT path for the current system, if one exists.

		Maps to the [` + "`" + `runtime.GOROOT()` + "`" + `](https://pkg.go.dev/runtime#GOROOT)
		Go method, which can be used in ` + Terratest + `.
		`)),
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending RuntimeGoroot Function Definition method.")
}

func (f *runtimeGorootFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting RuntimeGoroot Function Run method.")

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, runtime.GOROOT()))

	tflog.Debug(ctx, "Ending RuntimeGoroot Function Run method.")
}
