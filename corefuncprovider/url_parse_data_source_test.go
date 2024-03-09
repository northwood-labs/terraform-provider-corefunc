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

package corefuncprovider

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

const resourceAddr = "data.corefunc_url_parse.url"

func TestAccURLParseDataSource(t *testing.T) {
	funcName := traceFuncName()

	for name, tc := range testfixtures.URLParseTestTable { // lint:no_dupe
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		buf := &bytes.Buffer{}
		tmpl := template.Must(
			template.ParseFiles("url_parse_data_source_fixture.tftpl"),
		)

		err := tmpl.Execute(buf, tc)
		if err != nil {
			t.Error(err)
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
						resource.TestCheckResourceAttr(resourceAddr, "normalized", tc.Href),
						resource.TestCheckResourceAttr(resourceAddr, "protocol", tc.Protocol),
						resource.TestCheckResourceAttr(resourceAddr, "scheme", tc.Scheme),
						resource.TestCheckResourceAttr(resourceAddr, "username", tc.Username),
						resource.TestCheckResourceAttr(resourceAddr, "password", tc.Password),
						resource.TestCheckResourceAttr(resourceAddr, "host", tc.Host),
						resource.TestCheckResourceAttr(resourceAddr, "port", tc.Port),
						resource.TestCheckResourceAttr(resourceAddr, "path", tc.Path),
						resource.TestCheckResourceAttr(resourceAddr, "search", tc.Search),
						resource.TestCheckResourceAttr(resourceAddr, "query", tc.Query),
						resource.TestCheckResourceAttr(resourceAddr, "hash", tc.Hash),
						resource.TestCheckResourceAttr(resourceAddr, "fragment", tc.Fragment),
						resource.TestCheckResourceAttr(resourceAddr, "decoded_port", fmt.Sprint(tc.DecodedPort)),
					),
				},
			},
		})
	}
}
