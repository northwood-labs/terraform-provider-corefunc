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
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTruncateLabelDataSourceDefaultMaxLength64(t *testing.T) {
	buf := &bytes.Buffer{}
	tmpl := template.Must(
		template.ParseFiles("truncate_label_data_source_fixture_default64.tftpl"),
	)

	err := tmpl.Execute(buf, nil)
	if err != nil {
		t.Error(err)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + buf.String(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.corefunc_str_truncate_label.truncated",
						"value",
						"NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…",
					),
				),
			},
		},
	})
}

func TestAccTruncateLabelDataSource(t *testing.T) {
	funcName := traceFuncName()

	for name, tc := range testfixtures.TruncateLabelTestTable { // lint:no_dupe
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		buf := &bytes.Buffer{}
		tmpl := template.Must(
			template.ParseFiles("truncate_label_data_source_fixture_maxlength.tftpl"),
		)

		// Minimum value for the provider is 1.
		if tc.MaxLength == 0 {
			tc.MaxLength += 1
		}

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
						resource.TestCheckResourceAttr(
							"data.corefunc_str_truncate_label.label",
							"value",
							func() string {
								if tc.Expected == "" {
									return "…"
								}

								return tc.Expected
							}(),
						),
					),
				},
			},
		})
	}
}
