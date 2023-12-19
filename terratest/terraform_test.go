// Copyright 2023-2024, Ryan Parman
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

package terratest

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
	"github.com/stretchr/testify/assert"
)

var (
	inputStr     = "This is [an] {example}$${id32}."
	prefix       = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label        = "K8S Pods Not Running Deployment Check"
	strToReplace = "This is a string for testing replacements. New Relic. Set-up."
)

func TestTerraform(t *testing.T) {
	// https://golang.org/pkg/testing/#T.Parallel
	t.Parallel()

	// https://pkg.go.dev/github.com/gruntwork-io/terratest/modules/terraform#Options
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "./",

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	assert.Equal(t, terraform.Output(t, terraformOptions, "str_camel"), corefunc.StrCamel(inputStr))
	assert.Equal(t, terraform.Output(t, terraformOptions, "str_constant"), corefunc.StrConstant(inputStr))
	assert.Equal(t, terraform.Output(t, terraformOptions, "str_kebab"), corefunc.StrKebab(inputStr))
	assert.Equal(t, terraform.Output(t, terraformOptions, "str_pascal"), corefunc.StrPascal(inputStr, false))
	assert.Equal(t, terraform.Output(t, terraformOptions, "str_snake"), corefunc.StrSnake(inputStr))
	assert.Equal(t, terraform.Output(t, terraformOptions, "int_leftpad"), corefunc.IntLeftPad(123, 5))
	assert.Equal(t, terraform.Output(t, terraformOptions, "str_leftpad"), corefunc.StrLeftPad("abc", 5, '.'))
	assert.Equal(t, terraform.Output(t, terraformOptions, "env_ensure"), os.Getenv("GOROOT"))

	assert.Equal(
		t,
		terraform.Output(t, terraformOptions, "str_truncate"),
		corefunc.TruncateLabel(64, prefix, label),
	)

	assert.Equal(
		t,
		terraform.Output(t, terraformOptions, "str_iterative_replace"),
		corefunc.StrIterativeReplace(
			strToReplace,
			[]types.Replacement{
				{Old: ".", New: ""},
				{Old: " ", New: "_"},
				{Old: "-", New: "_"},
				{Old: "New_Relic", New: "datadog"},
				{Old: "This", New: "this"},
				{Old: "Set_up", New: "setup"},
			},
		),
	)

	homedir, err := corefunc.Homedir()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_get"), homedir)

	homedirPath, err := corefunc.HomedirExpand("~/.aws")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_expand"), homedirPath)
}
