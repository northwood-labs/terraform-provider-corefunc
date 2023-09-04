package corefuncprovider

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccEnvEnsureDataSource(t *testing.T) {
	funcName := traceFuncName()

	for name, tc := range testfixtures.EnvEnsureTestTable {
		fmt.Printf(
			"=== RUN   %s/%s\n",
			strings.TrimSpace(funcName),
			strings.TrimSpace(name),
		)

		_ = os.Setenv(tc.EnvVarName, tc.SetValue)

		// We expect the error to be nil.
		if tc.ExpectedErr == nil {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: providerConfig + `
					data "corefunc_env_ensure" "env" {
						name = "` + tc.EnvVarName + `"
					}
					`,
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
						Config: providerConfig + `
                    data "corefunc_env_ensure" "env" {
                        name = "` + tc.EnvVarName + `"
                    }
                    `,
						ExpectError: regexp.MustCompile("environment variable (\\w+) is not defined"),
					},
				},
			})
		}
	}
}
