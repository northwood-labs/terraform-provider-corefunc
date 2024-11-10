// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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
var _ function.Function = &netCidrContainsFunction{}

type (
	// netCidrContainsFunction is the function implementation.
	netCidrContainsFunction struct{}
)

// NetCidrContainsFunction is a method that exposes its paired Go function as a
// Terraform Function.
// https://developer.hashicorp.com/terraform/plugin/framework/functions/implementation
func NetCidrContainsFunction() function.Function { // lint:allow_return_interface
	return &netCidrContainsFunction{}
}

func (f *netCidrContainsFunction) Metadata(
	ctx context.Context,
	req function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains Function Metadata method.")

	resp.Name = "net_cidr_contains"

	tflog.Debug(ctx, fmt.Sprintf("resp.Name = %s", resp.Name))

	tflog.Debug(ctx, "Ending NetCidrContains Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *netCidrContainsFunction) Definition( // lint:nodupe
	ctx context.Context,
	req function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "CIDRContains checks to see if an IP address or CIDR block is contained within another CIDR block.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		CIDRContains checks to see if an IP address or CIDR block is contained
		within another CIDR block.

		-> Ported from OpenTofu.

		Maps to the ` + linkPackage("CIDRContains") + ` Go method, which can be used in ` + Terratest + `.
		`)),
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "containing_cidr",
				MarkdownDescription: "A CIDR range to check as a containing range.",
			},
			function.StringParameter{
				Name:                "contained_ip_or_cidr",
				MarkdownDescription: "An IP address or CIDR range to check as a contained range.",
			},
		},
		Return: function.BoolReturn{},
	}

	tflog.Debug(ctx, "Ending NetCidrContains Function Definition method.")
}

func (f *netCidrContainsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	tflog.Debug(ctx, "Starting NetCidrContains Function Run method.")

	var (
		containerCidr     string
		containedIPOrCidr string
	)

	err := req.Arguments.Get(ctx, &containerCidr, &containedIPOrCidr)

	resp.Error = function.ConcatFuncErrors(err)
	if resp.Error != nil {
		return
	}

	isContained, e := corefunc.CIDRContains(containerCidr, containedIPOrCidr)
	if e != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(e.Error()),
		)

		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, isContained))

	tflog.Debug(ctx, "Ending NetCidrContains Function Run method.")
}
