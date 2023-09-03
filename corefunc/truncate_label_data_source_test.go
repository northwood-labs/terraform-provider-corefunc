package corefunc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTruncateLabelDataSourceDefaultMaxLength64(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				data "corefunc_truncate_label" "truncated" {
					prefix = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
					label  = "K8S Pods Not Running Deployment Check"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.corefunc_truncate_label.truncated",
						"value",
						"NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…",
					),
				),
			},
		},
	})
}

func TestAccTruncateLabelDataSourceTruncateTests(t *testing.T) {
	funcName := traceFuncName()

	for name, tc := range tests {
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
					data "corefunc_truncate_label" "label" {
						prefix = "` + tc.prefix + `"
						label = "` + tc.label + `"
						max_length   = ` +
						func() string {
							if tc.maxLength == 0 {
								tc.maxLength += 1
							}

							return fmt.Sprint(tc.maxLength)
						}() + `
					}
					`,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"data.corefunc_truncate_label.label",
							"value",
							func() string {
								if tc.expected == "" {
									return "…"
								}

								return tc.expected
							}(),
						),
					),
				},
			},
		})
	}
}
