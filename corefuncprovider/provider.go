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

package corefuncprovider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
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
	tflog.Debug(ctx, "Starting Provider Metadata method.")

	resp.TypeName = "corefunc"

	tflog.Debug(ctx, "Ending Provider Metadata method.")
}

// Schema defines the provider-level schema for configuration data.
func (p *coreFuncProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	tflog.Debug(ctx, "Starting Provider Schema method.")

	resp.Schema = schema.Schema{
		MarkdownDescription: strings.TrimSpace(dedent.Dedent(`
        Utilities that should have been Terraform/OpenTofu _core functions_.

        While some of these _can_ be implemented in HCL, some of them begin to
        push up against the limits of Terraform and the HCL2 configuration
        language. We also perform testing using the
        [Terratest](https://terratest.gruntwork.io) framework on a regular basis.
        Exposing these functions as both a Go library as well as a Terraform/OpenTofu
        provider enables us to use the same functionality in both our Terraform/OpenTofu
        applies as well as while using a testing framework.

        Since earlier versions of Terraform/OpenTofu didn't have the concept of
        user-defined functions, the next step to open up the possibilities was
        to write a custom Provider which has the functions built-in, using
        existing support for inputs and outputs.

        **This does not add new syntax or constructs.** Instead it uses the
        _existing_ concepts around Providers, Resources, Data Sources,
        Variables, Outputs, and Functions (1.8) to expose new custom-built
        functionality.

        The goal of this provider is not to call any APIs, but to provide
        pre-built functions in the form of _Data Sources_ or _Provider
        Functions_ (1.8).
        `)),
	}

	tflog.Debug(ctx, "Ending Provider Schema method.")
}

func (p *coreFuncProvider) Configure(
	ctx context.Context,
	_ provider.ConfigureRequest,
	_ *provider.ConfigureResponse,
) {
	tflog.Debug(ctx, "Starting Provider Configure method.")
	tflog.Debug(ctx, "Ending Provider Configure method.")
}

// DataSources defines the data sources implemented in the provider.
func (p *coreFuncProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	tflog.Debug(ctx, "Running Provider DataSources method.")

	return []func() datasource.DataSource{
		EnvEnsureDataSource,
		HomedirExpandDataSource,
		HomedirGetDataSource,
		IntLeftpadDataSource,
		NetCidrContainsDataSource,
		RuntimeCpuarchDataSource,
		RuntimeGorootDataSource,
		RuntimeNumcpusDataSource,
		RuntimeOsDataSource,
		StrBase64GunzipDataSource,
		StrCamelDataSource,
		StrConstantDataSource,
		StrIterativeReplaceDataSource,
		StrKebabDataSource,
		StrLeftpadDataSource,
		StrPascalDataSource,
		StrSnakeDataSource,
		TomlToJSONDataSource,
		TruncateLabelDataSource,
		URLDecodeDataSource,
		URLParseDataSource,
	}
}

// Functions defines the functions implemented in the provider.
func (p *coreFuncProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		EnvEnsureFunction,
		HomedirExpandFunction,
		HomedirGetFunction,
		IntLeftpadFunction,
		NetCidrContainsFunction,
		RuntimeCpuarchFunction,
		RuntimeGorootFunction,
		RuntimeNumcpusFunction,
		RuntimeOsFunction,
		StrBase64GunzipFunction,
		StrCamelFunction,
		StrConstantFunction,
		StrKebabFunction,
		StrIterativeReplaceFunction,
		StrLeftpadFunction,
		StrPascalFunction,
		StrSnakeFunction,
		StrTruncateLabelFunction,
		TomlToJSONFunction,
		URLDecodeFunction,
		URLParseFunction,
	}
}

// Resources defines the resources implemented in the provider.
func (p *coreFuncProvider) Resources(ctx context.Context) []func() resource.Resource {
	tflog.Debug(ctx, "Running Provider Resources method.")

	return []func() resource.Resource{}
}
