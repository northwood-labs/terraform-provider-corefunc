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
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/testfixtures"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccHashArgon2idFunction(t *testing.T) {
	t.Parallel()

	if os.Getenv("CI") != "" {
		t.Skip("Skipping Argon2id tests in CI environment.")
	}

	funcName := traceFuncName()

	for name, tc := range testfixtures.HashArgon2idTestTable { // lint:no_dupe
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		buf := &bytes.Buffer{}
		tmpl := template.Must(
			template.New("hash_argon2id_function_fixture.tftpl").
				Funcs(FuncMap()).
				ParseFiles("hash_argon2id_function_fixture.tftpl"),
		)

		err := tmpl.Execute(buf, tc)
		if err != nil {
			log.Fatalln(err)
		}

		if os.Getenv("PROVIDER_DEBUG") != "" {
			fmt.Fprintln(os.Stderr, buf.String())
		}

		resource.UnitTest(t, resource.TestCase{
			TerraformVersionChecks: []tfversion.TerraformVersionCheck{
				tfversion.SkipBelow(tfversion.Version1_8_0),
			},
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: providerConfig + buf.String(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckOutput("hash_argon2id", tc.Expected),
					),
					// ExpectError: tc.expectedError,
				},
			},
		})
	}
}
