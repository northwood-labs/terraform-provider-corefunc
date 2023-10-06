// Copyright 2023, Ryan Parman
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

package corefuncprovider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/lithammer/dedent"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &coreFuncProvider{}
)

type (
	// coreFuncProvider is the provider implementation.
	coreFuncProvider struct{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider { // lint:allow_return_interface
	return &coreFuncProvider{}
}

// Metadata returns the provider type name.
func (p *coreFuncProvider) Metadata(
	ctx context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	tflog.Info(ctx, "Starting Provider Metadata method.")

	resp.TypeName = "corefunc"

	tflog.Info(ctx, "Ending Provider Metadata method.")
}

// Schema defines the provider-level schema for configuration data.
func (p *coreFuncProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	tflog.Info(ctx, "Starting Provider Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Utilities that should have been Terraform _core functions_.

        While some of these _can_ be implemented in HCL, some of them begin to
        push up against the limits of Terraform and the HCL2 configuration
        language. We also perform testing using the
        [Terratest](https://terratest.gruntwork.io) framework on a regular basis.
        Exposing these functions as both a Go library as well as a Terraform
        provider enables us to use the same functionality in both our Terraform
        applies as well as while using a testing framework.

        Since Terraform doesn't have the concept of user-defined functions, the
        next step to open up the possibilities is to write a custom Terraform
        Provider which has the functions built-in, using Terraform's existing
        support for inputs and outputs.

        **This does not add new syntax or constructs to Terraform.** Instead it
        uses the _existing_ concepts around Providers, Resources, Data Sources,
        Variables, and Outputs to expose new custom-built functionality.

        The goal of this provider is not to call any APIs, but to provide
        pre-built functions in the form of _Data Sources_.
        `)),
	}

	tflog.Info(ctx, "Ending Provider Schema method.")
}

func (p *coreFuncProvider) Configure(
	ctx context.Context,
	_ provider.ConfigureRequest,
	_ *provider.ConfigureResponse,
) {
	tflog.Info(ctx, "Starting Provider Configure method.")
	tflog.Info(ctx, "Ending Provider Configure method.")
}

// DataSources defines the data sources implemented in the provider.
func (p *coreFuncProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	tflog.Info(ctx, "Running Provider DataSources method.")

	return []func() datasource.DataSource{
		EnvEnsureDataSource,
		StrCamelDataSource,
		StrKebabDataSource,
		StrPascalDataSource,
		StrSnakeDataSource,
		TruncateLabelDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *coreFuncProvider) Resources(ctx context.Context) []func() resource.Resource {
	tflog.Info(ctx, "Running Provider Resources method.")

	return []func() resource.Resource{}
}
