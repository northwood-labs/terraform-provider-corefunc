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
	"runtime"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

// Ensure the implementation satisfies the expected interfaces.
var _ function.Function = &runtimeCpuarchFunction{}

type (
	// runtimeCpuarchFunction is the function implementation.
	runtimeCpuarchFunction struct{}
)

// RuntimeCpuarchFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func RuntimeCpuarchFunction() function.Function { // lint:allow_return_interface
	return &runtimeCpuarchFunction{}
}

func (f *runtimeCpuarchFunction) Metadata(
	ctx context.Context,
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch Function Metadata method.")

	resp.Name = "runtime_cpuarch"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending RuntimeCpuarch Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *runtimeCpuarchFunction) Definition(
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "Returns the CPU architecture for which the provider was compiled.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		Returns the CPU architecture for which the provider was compiled.
		Generally, this represents the current system.

		-> If you are running the provider via Rosetta on an Apple Silicon
		machine, using QEMU, or similar CPU emulation software, this will return
		the architecture of the emulated CPU, not the host CPU.

		Maps to the [` + "`" + `runtime.GOARCH` + "`" + `](https://pkg.go.dev/runtime#GOARCH)
		Go property, which can be used in ` + Terratest + `.
		`)),
		Return: function.StringReturn{},
	}

	tflog.Debug(ctx, "Ending RuntimeCpuarch Function Definition method.")
}

func (f *runtimeCpuarchFunction) Run(ctx context.Context, _ function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting RuntimeCpuarch Function Run method.")

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, runtime.GOARCH))

	tflog.Debug(ctx, "Ending RuntimeCpuarch Function Run method.")
}
