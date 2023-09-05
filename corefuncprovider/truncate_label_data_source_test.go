package corefuncprovider

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
	"text/template"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTruncateLabelDataSourceDefaultMaxLength64(t *testing.T) {
	buf := new(bytes.Buffer)
	tmpl := template.Must(
		template.ParseFiles("truncate_label_data_source_fixture_default64.tftpl"),
	)

	err := tmpl.Execute(buf, nil)
	if err != nil {
		log.Fatalln(err)
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

	for name, tc := range testfixtures.TruncateLabelTestTable {
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		buf := new(bytes.Buffer)
		tmpl := template.Must(
			template.ParseFiles("truncate_label_data_source_fixture_maxlength.tftpl"),
		)

		// Minimum value for the provider is 1.
		if tc.MaxLength == 0 {
			tc.MaxLength += 1
		}

		err := tmpl.Execute(buf, tc)
		if err != nil {
			log.Fatalln(err)
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
