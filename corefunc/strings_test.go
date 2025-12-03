// Copyright 2024-2025, Northwood Labs, LLC <license@northwood-labs.com>
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

	"github.com/northwood-labs/terraform-provider-corefunc/v2/testfixtures"
)

func ExampleStrStartsWith() {
	output := StrStartsWith("abcde", "abc")
	fmt.Println(output)

	output = StrStartsWith("abcde", "bcd")
	fmt.Println(output)

	// Output:
	// true
	// false
}

func ExampleStrEndsWith() {
	output := StrEndsWith("abcde", "cde")
	fmt.Println(output)

	output = StrEndsWith("abcde", "bcd")
	fmt.Println(output)

	// Output:
	// true
	// false
}

func ExampleStrContains() {
	output := StrContains("abcde", "bcd")
	fmt.Println(output)

	output = StrContains("abcde", "xyz")
	fmt.Println(output)

	// Output:
	// true
	// false
}

func TestStrStartsWith(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.StrStartsWithTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrStartsWith(tc.Input, tc.Prefix)

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}

func TestStrEndsWith(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.StrEndsWithTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrEndsWith(tc.Input, tc.Suffix)

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}

func TestStrContains(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.StrContainsTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrContains(tc.Input, tc.Substr)

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}

func BenchmarkStrStartsWith(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrStartsWithTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = StrStartsWith(tc.Input, tc.Prefix) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrEndsWith(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrEndsWithTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = StrEndsWith(tc.Input, tc.Suffix) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrContains(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrContainsTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = StrContains(tc.Input, tc.Substr) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrStartsWithParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrStartsWithTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrStartsWith(tc.Input, tc.Prefix) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkStrEndsWithParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrEndsWithTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrEndsWith(tc.Input, tc.Suffix) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkStrContainsParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrContainsTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrContains(tc.Input, tc.Substr) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzStrStartsWith(f *testing.F) {
	for _, tc := range testfixtures.StrStartsWithTestTable {
		f.Add(tc.Input, tc.Prefix)
	}

	f.Fuzz(
		func(_ *testing.T, in, prefix string) {
			_ = StrStartsWith(in, prefix) // lint:allow_unhandled
		},
	)
}

func FuzzStrEndsWith(f *testing.F) {
	for _, tc := range testfixtures.StrEndsWithTestTable {
		f.Add(tc.Input, tc.Suffix)
	}

	f.Fuzz(
		func(_ *testing.T, in, suffix string) {
			_ = StrEndsWith(in, suffix) // lint:allow_unhandled
		},
	)
}

func FuzzStrContains(f *testing.F) {
	for _, tc := range testfixtures.StrContainsTestTable {
		f.Add(tc.Input, tc.Substr)
	}

	f.Fuzz(
		func(_ *testing.T, in, substr string) {
			_ = StrContains(in, substr) // lint:allow_unhandled
		},
	)
}
