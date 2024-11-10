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

package terratest

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/hairyhenderson/go-which"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc"
	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

const (
	errTerraformPath     = "failed to set TF_ACC_TERRAFORM_PATH"
	errProviderNamespace = "failed to set TF_ACC_PROVIDER_NAMESPACE"
	errProviderHost      = "failed to set TF_ACC_PROVIDER_HOST"
)

var (
	err error

	inputStr     = "This is [an] {example}$${id32}."
	strToReplace = "This is a string for testing replacements. New Relic. Set-up."

	exampleCom       = "example.com"
	exampleQuery     = "http://u:p@example.com/foo?q=1"
	exampleQueryFrag = "http://u:p@example.com/foo?q=1#bar"

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
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_camel_ds"), corefunc.StrCamel(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_constant_ds"), corefunc.StrConstant(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_kebab_ds"), corefunc.StrKebab(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_pascal_ds"), corefunc.StrPascal(inputStr, false))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_snake_ds"), corefunc.StrSnake(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "int_leftpad_ds"), corefunc.IntLeftPad(123, 5))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_leftpad_ds"), corefunc.StrLeftPad("abc", 5, '.'))
		assert.Equal(t, terraform.Output(t, terraformOptions, "env_ensure_ds"), os.Getenv("GOROOT"))
		assert.Equal(t,
			terraform.Output(t, terraformOptions, "str_iterative_replace_ds"),
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

		// Ported OpenTofu functions
		isContained := false
		decodedURL := ""

		isContained, err = corefunc.CIDRContains("192.168.2.0/20", "192.168.2.1")
		if err != nil {
			t.Fatal(err)
		}

		decodedURL, err = corefunc.URLDecode("mailto%3Aemail%3Fsubject%3Dthis%2Bis%2Bmy%2Bsubject")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "net_cidr_contains_ds"), fmt.Sprintf("%t", isContained))
		assert.Equal(t, terraform.Output(t, terraformOptions, "url_decode_ds"), decodedURL)

		// runtime
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_cpuarch_ds"), runtime.GOARCH)
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_goroot_ds"), runtime.GOROOT())
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_numcpus_ds"), fmt.Sprint(runtime.NumCPU()))
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_os_ds"), runtime.GOOS)

		// url_parse
		urlParse := terraform.OutputMap(t, terraformOptions, "url_parse_ds")
		assert.Equal(t, urlParse["decoded_port"], fmt.Sprint(80))
		assert.Equal(t, urlParse["fragment"], "bar")
		assert.Equal(t, urlParse["hash"], "#bar")
		assert.Equal(t, urlParse["host"], exampleCom)
		assert.Equal(t, urlParse["hostname"], exampleCom)
		assert.Equal(t, urlParse["normalized"], exampleQueryFrag)
		assert.Equal(t, urlParse["normalized_nofrag"], exampleQuery)
		assert.Equal(t, urlParse["password"], "p")
		assert.Equal(t, urlParse["path"], "/foo")
		assert.Equal(t, urlParse["port"], "")
		assert.Equal(t, urlParse["protocol"], "http:")
		assert.Equal(t, urlParse["query"], "q=1")
		assert.Equal(t, urlParse["scheme"], "http")
		assert.Equal(t, urlParse["search"], "?q=1")
		assert.Equal(t, urlParse["url"], "HTTP://u:p@example.com:80/foo?q=1#bar")
		assert.Equal(t, urlParse["username"], "u")

		// url_parse_gsb
		urlParseGSB := terraform.OutputMap(t, terraformOptions, "url_parse_gsb_ds")
		assert.Equal(t, urlParseGSB["decoded_port"], fmt.Sprint(80))
		assert.Equal(t, urlParseGSB["fragment"], "")
		assert.Equal(t, urlParseGSB["hash"], "")
		assert.Equal(t, urlParseGSB["host"], exampleCom)
		assert.Equal(t, urlParseGSB["hostname"], exampleCom)
		assert.Equal(t, urlParseGSB["normalized"], exampleQuery)
		assert.Equal(t, urlParseGSB["normalized_nofrag"], exampleQuery)
		assert.Equal(t, urlParseGSB["password"], "p")
		assert.Equal(t, urlParseGSB["path"], "/foo")
		assert.Equal(t, urlParseGSB["port"], "")
		assert.Equal(t, urlParseGSB["protocol"], "http:")
		assert.Equal(t, urlParseGSB["query"], "q=1")
		assert.Equal(t, urlParseGSB["scheme"], "http")
		assert.Equal(t, urlParseGSB["search"], "?q=1")
		assert.Equal(t, urlParseGSB["url"], "HTTP://u:p@example.com:80/foo?q=1#bar")
		assert.Equal(t, urlParseGSB["username"], "u")

		homedir := ""
		homedirPath := ""

		homedir, err = corefunc.Homedir()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_get_ds"), homedir)

		homedirPath, err = corefunc.HomedirExpand("~/.aws")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_expand_ds"), homedirPath)

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
		return errors.Wrap(err, errTerraformPath)
	}

	err = os.Setenv("TF_ACC_PROVIDER_NAMESPACE", "northwood-labs")
	if err != nil {
		return errors.Wrap(err, errProviderNamespace)
	}

	switch binary {
	case "terraform":
		err = os.Setenv("TF_ACC_PROVIDER_HOST", "registry.terraform.io")
		if err != nil {
			return errors.Wrap(err, errProviderHost)
		}
	case "tofu":
		err = os.Setenv("TF_ACC_PROVIDER_HOST", "registry.opentofu.org")
		if err != nil {
			return errors.Wrap(err, errProviderHost)
		}
	}

	yellow := lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
	green := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	magenta := lipgloss.NewStyle().Foreground(lipgloss.Color("13"))
	underline := lipgloss.NewStyle().Underline(true).Bold(true)

	fmt.Println(
		clihelpers.LongHelpText(`
		` + underline.Render("Testing Data Sources") + `

		Binary ` + yellow.Render(which.Which(binary)) + ` is installed
		TF_ACC_TERRAFORM_PATH=` + yellow.Render(os.Getenv("TF_ACC_TERRAFORM_PATH")) + `
		TF_ACC_PROVIDER_NAMESPACE=` + green.Render(os.Getenv("TF_ACC_PROVIDER_NAMESPACE")) + `
		TF_ACC_PROVIDER_HOST=` + magenta.Render(os.Getenv("TF_ACC_PROVIDER_HOST")) + `
		`),
	)

	return nil
}

func restoreEnv() error {
	// Restore the original values.
	err = os.Setenv("TF_ACC_TERRAFORM_PATH", origPath)
	if err != nil {
		return errors.Wrap(err, errTerraformPath)
	}

	err = os.Setenv("TF_ACC_PROVIDER_NAMESPACE", origNamespace)
	if err != nil {
		return errors.Wrap(err, errProviderNamespace)
	}

	err = os.Setenv("TF_ACC_PROVIDER_HOST", origHostname)
	if err != nil {
		return errors.Wrap(err, errProviderHost)
	}

	return nil
}
