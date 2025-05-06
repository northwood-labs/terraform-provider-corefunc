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
var _ function.Function = &runtimeNumcpusFunction{}

type (
	// runtimeNumcpusFunction is the function implementation.
	runtimeNumcpusFunction struct{}
)

// RuntimeNumcpusFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func RuntimeNumcpusFunction() function.Function { // lint:allow_return_interface
	return &runtimeNumcpusFunction{}
}

func (f *runtimeNumcpusFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus Function Metadata method.")

	resp.Name = "runtime_numcpus"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending RuntimeNumcpus Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *runtimeNumcpusFunction) Definition(
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Returns the number of CPU cores (logical CPUs) for the current system.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the number of CPU cores (logical CPUs) for the current system.

		-> This counts the number of CPU cores, which may differ from the number
		of physical CPUs.

		Maps to the [` + "`" + `runtime.NumCPU()` + "`" + `](https://pkg.go.dev/runtime#NumCPU)
		Go method, which can be used in ` + Terratest + `.
		`)),
		Return: function.Int64Return{},
	}

	tflog.Debug(ctx, "Ending RuntimeNumcpus Function Definition method.")
}

func (f *runtimeNumcpusFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting RuntimeNumcpus Function Run method.")

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, runtime.NumCPU()))

	tflog.Debug(ctx, "Ending RuntimeNumcpus Function Run method.")
}
