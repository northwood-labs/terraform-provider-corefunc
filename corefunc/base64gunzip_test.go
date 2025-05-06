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

package corefunc

import (
	"fmt"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleBase64Gunzip() {
	output, err := Base64Gunzip(
		"H4sIAAAAAAAA/wrJSFXIL0jNUyjOLy1KTlXIzEsrSiwuKSpNLiktSlVILFZIzk9JVSjJz8/RAwQAAP//HPstqCwAAAA",
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// Output:
	// The open source infrastructure as code tool.
}

func TestBase64Gunzip(t *testing.T) {
	for name, tc := range testfixtures.Base64GunzipTestTable {
		t.Run(name, func(t *testing.T) {
			got, err := Base64Gunzip(tc.Input)
			if err != nil && tc.ExpectError != nil {
				if !tc.ExpectError.MatchString(err.Error()) {
					t.Errorf("Base64Gunzip() error = %v, want match %v", err, tc.ExpectError)
				}

				return
			} else if err != nil {
				t.Errorf("Base64Gunzip() error = %v", err)

				return
			}

			if got != tc.Expected {
				t.Errorf("Base64Gunzip() = %v, want %v", got, tc.Expected)
			}
		})
	}
}

func BenchmarkBase64Gunzip(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.Base64GunzipTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, _ = Base64Gunzip(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkBase64GunzipParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.Base64GunzipTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = Base64Gunzip(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzBase64Gunzip(f *testing.F) {
	for _, tc := range testfixtures.Base64GunzipTestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(t *testing.T, input string) {
			_, _ = Base64Gunzip(input) // lint:allow_unhandled
		},
	)
}
