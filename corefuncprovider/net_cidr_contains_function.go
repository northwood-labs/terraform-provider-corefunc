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
	_ function.MetadataRequest,
	resp *function.MetadataResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains Function Metadata method.")

	resp.Name = "net_cidr_contains"

	tflog.Debug(ctx, "resp.Name = "+resp.Name)

	tflog.Debug(ctx, "Ending NetCidrContains Function Metadata method.")
}

// Definition defines the parameters and return type for the function.
func (f *netCidrContainsFunction) Definition( // lint:no_dupe
	ctx context.Context,
	_ function.DefinitionRequest,
	resp *function.DefinitionResponse,
) {
	tflog.Debug(ctx, "Starting NetCidrContains Function Definition method.")

	resp.Definition = function.Definition{
		Summary: "CIDRContains checks to see if an IP address or CIDR block is contained within another CIDR block.",
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
		CIDRContains determines whether or not a given IP address, or an address prefix
		given in CIDR notation, is within a given IP network address prefix.

		Both arguments must belong to the same address family, either IPv4 or IPv6. A
		family mismatch will result in an error.

		~> This functionality is built into OpenTofu 1.8, but has not been implemented
		in Terraform (as of version 1.15).

		-> This port from OpenTofu provides a 1:1 implementation that can be used with
		Terratest or other Go code, as well as with OpenTofu and Terraform going all
		the way back to v1.0.

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
