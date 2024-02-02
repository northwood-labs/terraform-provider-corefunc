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
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/hairyhenderson/go-which"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
	err error

	inputStr     = "This is [an] {example}$${id32}."
	prefix       = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label        = "K8S Pods Not Running Deployment Check"
	strToReplace = "This is a string for testing replacements. New Relic. Set-up."

	origPath      = os.Getenv("TF_ACC_TERRAFORM_PATH")
	origNamespace = os.Getenv("TF_ACC_PROVIDER_NAMESPACE")
	origHostname  = os.Getenv("TF_ACC_PROVIDER_HOST")

	// Both must be installed first.
	binaries = []string{
		"terraform",
		"tofu",
	}
)

func TestTerraform(t *testing.T) {
	// https://golang.org/pkg/testing/#T.Parallel
	t.Parallel()

	for i := range binaries {
		binary := binaries[i]

		if _, err = os.Stat(which.Which(binary)); os.IsNotExist(err) {
			t.Fatalf("Binary %s must be installed first", binary)
		}

		err = setAndPrint(binary)
		if err != nil {
			t.Fatal(err)
		}

		// https://pkg.go.dev/github.com/gruntwork-io/terratest/modules/terraform#Options
		terraformOptions := &terraform.Options{
			// The path to Terraform.
			TerraformBinary: which.Which(binary),

			// The path to where our Terraform code is located
			TerraformDir: "./",

			// Disable colors in Terraform commands so its easier to parse stdout/stderr
			NoColor: true,

			// terraform init -upgrade
			Upgrade: true,

			// terraform init -reconfigure
			Reconfigure: true,

			// Set the maximum number of parallelism; equivalent to `nproc`.
			Parallelism: runtime.NumCPU(),
		}

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

		homedir := ""
		homedirPath := ""

		homedir, err = corefunc.Homedir()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_get"), homedir)

		homedirPath, err = corefunc.HomedirExpand("~/.aws")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_expand"), homedirPath)

		// At the end of the test, run `terraform destroy` to clean up any resources that were created
		err = restoreEnv()
		if err != nil {
			t.Fatal(err)
		}

		terraform.Destroy(t, terraformOptions)
	}
}

func setAndPrint(binary string) error {
	// Set the necessary values.
	err = os.Setenv("TF_ACC_TERRAFORM_PATH", which.Which(binary))
	if err != nil {
		return errors.Wrap(err, "failed to set TF_ACC_TERRAFORM_PATH")
	}

	err = os.Setenv("TF_ACC_PROVIDER_NAMESPACE", "northwood-labs")
	if err != nil {
		return errors.Wrap(err, "failed to set TF_ACC_PROVIDER_NAMESPACE")
	}

	switch binary {
	case "terraform":
		err = os.Setenv("TF_ACC_PROVIDER_HOST", "registry.terraform.io")
		if err != nil {
			return errors.Wrap(err, "failed to set TF_ACC_PROVIDER_HOST")
		}
	case "tofu":
		err = os.Setenv("TF_ACC_PROVIDER_HOST", "registry.opentofu.org")
		if err != nil {
			return errors.Wrap(err, "failed to set TF_ACC_PROVIDER_HOST")
		}
	}

	fmt.Println("")
	fmt.Println("================================================================================")
	fmt.Printf("Binary %s is installed\n", which.Which(binary))
	fmt.Println("TF_ACC_TERRAFORM_PATH=" + os.Getenv("TF_ACC_TERRAFORM_PATH"))
	fmt.Println("TF_ACC_PROVIDER_NAMESPACE=" + os.Getenv("TF_ACC_PROVIDER_NAMESPACE"))
	fmt.Println("TF_ACC_PROVIDER_HOST=" + os.Getenv("TF_ACC_PROVIDER_HOST"))
	fmt.Println("================================================================================")
	fmt.Println("")

	return nil
}

func restoreEnv() error {
	// Restore the original values.
	err = os.Setenv("TF_ACC_TERRAFORM_PATH", origPath)
	if err != nil {
		return errors.Wrap(err, "failed to set TF_ACC_TERRAFORM_PATH")
	}

	err = os.Setenv("TF_ACC_PROVIDER_NAMESPACE", origNamespace)
	if err != nil {
		return errors.Wrap(err, "failed to set TF_ACC_PROVIDER_NAMESPACE")
	}

	err = os.Setenv("TF_ACC_PROVIDER_HOST", origHostname)
	if err != nil {
		return errors.Wrap(err, "failed to set TF_ACC_PROVIDER_HOST")
	}

	return nil
}
