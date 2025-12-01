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
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/hairyhenderson/go-which"
	"github.com/stretchr/testify/assert"

	clihelpers "github.com/northwood-labs/cli-helpers"
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc"
	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc/types"
)

var (
	err error

	inputStr     = "This is [an] {example}$${id32}."
	strToReplace = "This is a string for testing replacements. New Relic. Set-up."

	exampleCom       = "example.com"
	exampleQuery     = "http://u:p@example.com/foo?q=1"
	exampleQueryFrag = "http://u:p@example.com/foo?q=1#bar"

	// Both must be installed first.
	binaries = []string{
		"terraform",
		"tofu",
	}
)

func TestDataSources(t *testing.T) { // lint:allow_complexity
	tmpDir := "."

	for i := range binaries {
		binary := binaries[i]

		if _, err = os.Stat(which.Which(binary)); os.IsNotExist(err) {
			t.Fatalf("Binary %s must be installed first", binary)
		}

		setAndPrint(t, binary)

		// Apply
		test_structure.RunTestStage(t, "apply", func() {
			message("Applying Terraform configuration...")

			terraformOptions := &terraform.Options{
				TerraformBinary: which.Which(binary),
				TerraformDir:    tmpDir,
				NoColor:         true,
				Upgrade:         true,
				Reconfigure:     true,
				Parallelism:     1,
			}

			test_structure.SaveTerraformOptions(t, tmpDir, terraformOptions)
			terraform.InitAndApply(t, terraformOptions)
		})

		// String formatting
		test_structure.RunTestStage(t, "strings", func() {
			message("Testing string functions...")

			terraformOptions := test_structure.LoadTerraformOptions(t, tmpDir)

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
		})

		// Format shifting
		test_structure.RunTestStage(t, "formats", func() {
			message("Testing format shifting functions...")

			terraformOptions := test_structure.LoadTerraformOptions(t, tmpDir)

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
		})

		// Ported OpenTofu functions
		test_structure.RunTestStage(t, "ported", func() {
			message("Testing ported OpenTofu functions...")

			terraformOptions := test_structure.LoadTerraformOptions(t, tmpDir)

			var (
				isContained bool
				decodedURL  string
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

			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "net_cidr_contains_ds"),
				strconv.FormatBool(isContained),
			)
			assert.Equal(t, terraform.Output(t, terraformOptions, "str_base64_gunzip_ds"), unzipped)
			assert.Equal(t, terraform.Output(t, terraformOptions, "url_decode_ds"), decodedURL)

			// runtime
			assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_cpuarch_ds"), runtime.GOARCH)
			assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_numcpus_ds"), strconv.Itoa(runtime.NumCPU()))
			assert.Equal(t, terraform.Output(t, terraformOptions, "runtime_os_ds"), runtime.GOOS)

			// url_parse
			urlParse := terraform.OutputMap(t, terraformOptions, "url_parse_ds")
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
			urlParseGSB := terraform.OutputMap(t, terraformOptions, "url_parse_gsb_ds")
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
		})

		// Home directory
		test_structure.RunTestStage(t, "homedir", func() {
			message("Testing home directory functions...")

			terraformOptions := test_structure.LoadTerraformOptions(t, tmpDir)

			var (
				homedir     string
				homedirPath string
			)

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
		})

		// Hashing
		test_structure.RunTestStage(t, "hashing", func() {
			message("Testing hashing functions...")

			terraformOptions := test_structure.LoadTerraformOptions(t, tmpDir)

			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_md5_ds"), corefunc.HashMD5("hello world"))
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_md5_base64_ds"),
				corefunc.Base64HashMD5("hello world"),
			)
			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_sha1_ds"), corefunc.HashSHA1("hello world"))
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha1_base64_ds"),
				corefunc.Base64HashSHA1("hello world"),
			)
			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_sha256_ds"), corefunc.HashSHA256("hello world"))
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha256_base64_ds"),
				corefunc.Base64HashSHA256("hello world"),
			)
			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_sha384_ds"), corefunc.HashSHA384("hello world"))
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha384_base64_ds"),
				corefunc.Base64HashSHA384("hello world"),
			)
			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_sha512_ds"), corefunc.HashSHA512("hello world"))
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha512_base64_ds"),
				corefunc.Base64HashSHA512("hello world"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha3x256_ds"),
				corefunc.HashSHA3x256("hello world"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha3x256_base64_ds"),
				corefunc.Base64HashSHA3x256("hello world"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha3x384_ds"),
				corefunc.HashSHA3x384("hello world"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha3x384_base64_ds"),
				corefunc.Base64HashSHA3x384("hello world"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha3x512_ds"),
				corefunc.HashSHA3x512("hello world"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_sha3x512_base64_ds"),
				corefunc.Base64HashSHA3x512("hello world"),
			)

			// Hashes with salts (Argon2id)
			hash, err := corefunc.HashArgon2id("hello world", []byte("somesaltvalue"))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_argon2id_ds"), hash)

			hash, err = corefunc.Base64HashArgon2id("hello world", []byte("somesaltvalue"))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_argon2id_base64_ds"), hash)

			// Hashes with salts (Scrypt)
			hash, err = corefunc.HashScrypt("hello world", []byte("somesaltvalue"))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_scrypt_ds"), hash)

			hash, err = corefunc.Base64HashScrypt("hello world", []byte("somesaltvalue"))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, terraform.Output(t, terraformOptions, "hash_scrypt_base64_ds"), hash)

			// Hashes with salts (HMACSHA256)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_hmac_sha256_ds"),
				corefunc.HashHMACSHA256("hello world", "secretkey"),
			)
			assert.Equal(
				t,
				terraform.Output(t, terraformOptions, "hash_hmac_base64sha256_ds"),
				corefunc.Base64HashHMACSHA256("hello world", "secretkey"),
			)
		})

		// Destroy
		test_structure.RunTestStage(t, "destroy", func() {
			message("Destroying resources...")

			terraformOptions := test_structure.LoadTerraformOptions(t, tmpDir)
			terraform.Destroy(t, terraformOptions)
		})
	}
}

func setAndPrint(t *testing.T, binary string) {
	t.Helper()

	var (
		tfAccTerraformPath     = which.Which(binary)
		tfAccProviderNamespace = "northwood-labs"
		tfAccProviderHost      string
	)

	// Set the necessary values.
	t.Setenv("TF_ACC_TERRAFORM_PATH", which.Which(binary))
	t.Setenv("TF_ACC_PROVIDER_NAMESPACE", "northwood-labs")

	switch binary {
	case "terraform":
		t.Setenv("TF_ACC_PROVIDER_HOST", "registry.terraform.io")

		tfAccProviderHost = "registry.terraform.io"
	case "tofu":
		t.Setenv("TF_ACC_PROVIDER_HOST", "registry.opentofu.org")

		tfAccProviderHost = "registry.opentofu.org"
	}

	fmt.Println(
		clihelpers.LongHelpText(`
		Testing Provider Data Sources with ` + binary + `

		Binary ` + which.Which(binary) + ` is installed
		TF_ACC_TERRAFORM_PATH=` + tfAccTerraformPath + `
		TF_ACC_PROVIDER_NAMESPACE=` + tfAccProviderNamespace + `
		TF_ACC_PROVIDER_HOST=` + tfAccProviderHost),
	)
}

func message(msg string) {
	fmt.Fprintln(os.Stderr, lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFC67")).
		BorderForeground(lipgloss.Color("12")).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Margin(1, 0, 0, 0).
		Bold(true).
		Render(msg))
}
