package corefuncprovider

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
	"text/template"

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

		buf := new(bytes.Buffer)
		tmpl := template.Must(
			template.ParseFiles("env_ensure_data_source_fixture.tftpl"),
		)

		err := tmpl.Execute(buf, tc)
		if err != nil {
			log.Fatalln(err)
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
						ExpectError: regexp.MustCompile("environment variable (\\w+) (is not defined|does not match pattern)"),
					},
				},
			})
		}
	}
}
