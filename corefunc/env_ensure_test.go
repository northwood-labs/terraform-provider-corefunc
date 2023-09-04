package corefunc

import (
	"fmt"
	"os"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
	"github.com/stretchr/testify/assert"
)

func ExampleEnvEnsure() {
	_ = os.Setenv("MY_ENV_VAR_VALUE", "abcd1234")
	_ = os.Setenv("MY_ENV_VAR_EMPTY", "")

	err := EnvEnsure("MY_ENV_VAR_VALUE")
	fmt.Println(err)

	err = EnvEnsure("MY_ENV_VAR_EMPTY")
	fmt.Println(err)

	err = EnvEnsure("MY_ENV_VAR_NOT_SET")
	fmt.Println(err)
	// Output:
	// <nil>
	// environment variable MY_ENV_VAR_EMPTY is not defined
	// environment variable MY_ENV_VAR_NOT_SET is not defined
}

func TestEnvEnsure(t *testing.T) {
	for name, tc := range testfixtures.EnvEnsureTestTable {
		t.Run(name, func(t *testing.T) {
			_ = os.Setenv(tc.EnvVarName, tc.SetValue)
			err := EnvEnsure(tc.EnvVarName)

			if tc.ExpectedErr != nil {
				expectedErrorMsg := tc.ExpectedErr.Error()
				assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, err)
			}
		})
	}
}

func BenchmarkEnvEnsure(b *testing.B) {
	b.ReportAllocs()
	for name, tc := range testfixtures.EnvEnsureTestTable {
		_ = os.Setenv(tc.EnvVarName, tc.SetValue)

		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				EnvEnsure(tc.EnvVarName)
			}
		})
	}
}

func BenchmarkEnvEnsureParallel(b *testing.B) {
	b.ReportAllocs()
	for name, tc := range testfixtures.EnvEnsureTestTable {
		_ = os.Setenv(tc.EnvVarName, tc.SetValue)

		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					EnvEnsure(tc.EnvVarName)
				}
			})
		})
	}
}

func FuzzEnvEnsure(f *testing.F) {
	for _, tc := range testfixtures.EnvEnsureTestTable {
		f.Add(tc.EnvVarName)
	}

	f.Fuzz(
		func(t *testing.T, envVarName string) {
			_ = EnvEnsure(envVarName)
		},
	)
}
