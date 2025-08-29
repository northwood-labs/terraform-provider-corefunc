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
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/testfixtures"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEnvEnsureDataSource(t *testing.T) {
	funcName := traceFuncName()

	for name, tc := range testfixtures.EnvEnsureTestTable { // lint:no_dupe
		var err error

		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		t.Setenv(tc.EnvVarName, tc.SetValue)

		buf := &bytes.Buffer{}
		tmpl := template.Must(
			template.ParseFiles("env_ensure_data_source_fixture.tftpl"),
		)

		err = tmpl.Execute(buf, tc)
		if err != nil {
			t.Error(err)
		}

		if os.Getenv("PROVIDER_DEBUG") != "" {
			fmt.Fprintln(os.Stderr, buf.String())
		}

		// We expect the error to be nil.
		if tc.ExpectedErr == nil {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: providerConfig + buf.String(),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.corefunc_env_ensure.env", "value", tc.SetValue),
						),
					},
				},
			})
		} else {
			// We DO expect an error.
			// https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      providerConfig + buf.String(),
						ExpectError: tc.ExpectedErrRE,
					},
				},
			})
		}
	}
}
