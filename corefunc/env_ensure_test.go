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

package corefunc

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
	"github.com/stretchr/testify/assert"
)

func ExampleEnvEnsure() {
	var err error

	// ---------------------------------------------------------
	// Set the environment variables.

	err = os.Setenv("MY_ENV_VAR_VALUE", "abcd1234")
	fmt.Println(err)

	err = os.Setenv("MY_ENV_VAR_EMPTY", "")
	fmt.Println(err)

	// ---------------------------------------------------------
	// Ensure the environment variables are set.

	err = EnvEnsure("MY_ENV_VAR_VALUE")
	fmt.Println(err)

	err = EnvEnsure("MY_ENV_VAR_EMPTY")
	fmt.Println(err)

	err = EnvEnsure("MY_ENV_VAR_NOT_SET")
	fmt.Println(err)
	// Output:
	// <nil>
	// <nil>
	// <nil>
	// environment variable MY_ENV_VAR_EMPTY is not defined
	// environment variable MY_ENV_VAR_NOT_SET is not defined
}

func ExampleEnvEnsure_pattern() {
	var err error

	// ---------------------------------------------------------
	// Set the environment variables.

	err = os.Setenv("AWS_VAULT", "dev")
	fmt.Println(err)

	// ---------------------------------------------------------
	// Ensure the environment variables are set.

	// This will not throw an error because it is defined, and there is no pattern to match.
	err = EnvEnsure("AWS_VAULT")
	fmt.Println(err)

	// This will throw an error because the pattern does not match.
	err = EnvEnsure("AWS_VAULT", regexp.MustCompile(`(non)?prod$`))
	fmt.Println(err)

	// Output:
	// <nil>
	// <nil>
	// environment variable AWS_VAULT does not match pattern (non)?prod$
}

func TestEnvEnsure(t *testing.T) {
	for name, tc := range testfixtures.EnvEnsureTestTable {
		t.Run(name, func(t *testing.T) {
			_ = os.Setenv(tc.EnvVarName, tc.SetValue) // lint:allow_unhandled
			err := EnvEnsure(tc.EnvVarName, tc.Pattern)

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
		_ = os.Setenv(tc.EnvVarName, tc.SetValue) // lint:allow_unhandled

		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = EnvEnsure(tc.EnvVarName) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkEnvEnsureParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.EnvEnsureTestTable {
		_ = os.Setenv(tc.EnvVarName, tc.SetValue) // lint:allow_unhandled

		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = EnvEnsure(tc.EnvVarName) // lint:allow_unhandled
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
			_ = EnvEnsure(envVarName) // lint:allow_unhandled
		},
	)
}
