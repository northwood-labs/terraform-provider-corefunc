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

package terratest

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/hairyhenderson/go-which"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc/types"
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
		"terraform", // 1.8.0+
		"tofu",      // 1.7.0+
	}
)

func TestTerraform(t *testing.T) { // lint:allow_complexity
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
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_camel_fn"), corefunc.StrCamel(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_constant_fn"), corefunc.StrConstant(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_kebab_fn"), corefunc.StrKebab(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_pascal_fn"), corefunc.StrPascal(inputStr, false))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_snake_fn"), corefunc.StrSnake(inputStr))
		assert.Equal(t, terraform.Output(t, terraformOptions, "int_leftpad_fn"), corefunc.IntLeftPad(123, 5))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_leftpad_fn"), corefunc.StrLeftPad("abc", 5, '.'))
		assert.Equal(t, terraform.Output(t, terraformOptions, "env_ensure_fn"), os.Getenv("GOROOT"))
		assert.Equal(t,
			terraform.Output(t, terraformOptions, "str_iterative_replace_fn"),
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

		// Format shifting
		var (
			asJSON string
			asTOML string
		)

		asJSON, err = corefunc.TOMLtoJSON("abc = 123")
		if err != nil {
			t.Fatal(err)
		}

		asTOML, err = corefunc.JSONtoTOML(`{"abc": 123}`)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "toml_as_json"), strings.TrimSpace(asJSON))
		assert.Equal(t, terraform.Output(t, terraformOptions, "json_as_toml"), strings.TrimSpace(asTOML))

		// Ported OpenTofu functions
		var (
			decodedURL  string
			homedir     string
			homedirPath string
			isContained bool
			unzipped    string
		)

		isContained, err = corefunc.CIDRContains("192.168.2.0/20", "192.168.2.1")
		if err != nil {
			t.Fatal(err)
		}

		decodedURL, err = corefunc.URLDecode("mailto%3Aemail%3Fsubject%3Dthis%2Bis%2Bmy%2Bsubject")
		if err != nil {
			t.Fatal(err)
		}

		unzipped, err = corefunc.Base64Gunzip("H4sIAAAAAAAA/8pIrVTISK3UUShPVS9KVSjJSFXIzc/LTk0tBgQAAP//qz+dmhoAAAA")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "net_cidr_contains_fn"), strconv.FormatBool(isContained))
		assert.Equal(t, terraform.Output(t, terraformOptions, "str_base64_gunzip_fn"), unzipped)
		assert.Equal(t, terraform.Output(t, terraformOptions, "url_decode_fn"), decodedURL)

		// runtime
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_cpuarch_fn"), runtime.GOARCH)
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_numcpus_fn"), strconv.Itoa(runtime.NumCPU()))
		assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_os_fn"), runtime.GOOS)

		// url_parse
		urlParse := terraform.OutputMap(t, terraformOptions, "url_parse_fn")
		assert.Equal(t, urlParse["decoded_port"], strconv.Itoa(80))
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
		assert.Equal(t, urlParse["username"], "u")

		// url_parse_gsb
		urlParseGSB := terraform.OutputMap(t, terraformOptions, "url_parse_gsb_fn")
		assert.Equal(t, urlParseGSB["decoded_port"], strconv.Itoa(80))
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
		assert.Equal(t, urlParseGSB["username"], "u")

		homedir, err = corefunc.Homedir()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_get_fn"), homedir)

		homedirPath, err = corefunc.HomedirExpand("~/.aws")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, terraform.Output(t, terraformOptions, "homedir_expand_fn"), homedirPath)

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
		` + underline.Render("Testing Provider Functions") + `

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
