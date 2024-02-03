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

package corefuncprovider

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRuntimeNumcpusDataSource(t *testing.T) {
	funcName := traceFuncName()

	for name, tc := range testfixtures.RuntimeNumcpusTestTable { // lint:no_dupe
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		buf := &bytes.Buffer{}
		tmpl := template.Must(
			template.ParseFiles("runtime_numcpus_data_source_fixture.tftpl"),
		)

		err := tmpl.Execute(buf, tc)
		if err != nil {
			log.Fatalln(err)
		}

		if os.Getenv("PROVIDER_DEBUG") != "" {
			fmt.Fprintln(os.Stderr, buf.String())
		}

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: providerConfig + buf.String(),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"data.corefunc_runtime_numcpus.count",
							"value",
							fmt.Sprint(tc.Expected),
						),
					),
				},
			},
		})
	}
}
