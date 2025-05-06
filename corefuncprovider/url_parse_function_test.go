// Copyright 2023-2025, Northwood Labs
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

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccURLParseFunction(t *testing.T) {
	t.Parallel()

	funcName := traceFuncName()

	for name, tc := range testfixtures.URLParseTestTable { // lint:no_dupe
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		buf := &bytes.Buffer{}
		tmpl := template.Must(
			template.ParseFiles("url_parse_function_fixture.tftpl"),
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
				tfversion.SkipBelow(
					version.Must(version.NewVersion("1.7.999")),
				),
			},
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: providerConfig + buf.String(),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("normalized"),
							knownvalue.StringExact(tc.Href),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("protocol"),
							knownvalue.StringExact(tc.Protocol),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("scheme"),
							knownvalue.StringExact(tc.Scheme),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("username"),
							knownvalue.StringExact(tc.Username),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("password"),
							knownvalue.StringExact(tc.Password),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("host"),
							knownvalue.StringExact(tc.Host),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("port"),
							knownvalue.StringExact(tc.Port),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("path"),
							knownvalue.StringExact(tc.Path),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("search"),
							knownvalue.StringExact(tc.Search),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("query"),
							knownvalue.StringExact(tc.Query),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("hash"),
							knownvalue.StringExact(tc.Hash),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("fragment"),
							knownvalue.StringExact(tc.Fragment),
						),
						statecheck.ExpectKnownOutputValueAtPath(
							"url",
							tfjsonpath.New("decoded_port"),
							knownvalue.Int64Exact(int64(tc.DecodedPort)),
						),
					},
				},
			},
		})
	}
}
