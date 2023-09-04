package corefuncprovider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTruncateLabelDataSourceDefaultMaxLength64(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				data "corefunc_str_truncate_label" "truncated" {
					prefix = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
					label  = "K8S Pods Not Running Deployment Check"
				}
				`,
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

		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: providerConfig + `
					data "corefunc_str_truncate_label" "label" {
						prefix = "` + tc.Prefix + `"
						label = "` + tc.Label + `"
						max_length   = ` +
						func() string {
							if tc.MaxLength == 0 {
								tc.MaxLength += 1
							}

							return fmt.Sprint(tc.MaxLength)
						}() + `
					}
					`,
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
