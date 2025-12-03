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
	"os"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/testfixtures"
)

// func ExampleTimeParse() {
// 	output, err := TimeParse("2006-01-02T15:04:05Z")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(output)

// 	output, err = TimeParse("Monday, 02-Jan-2006 15:04:05 MST")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(output)

// 	// Output:
// 	// 1136214245
// 	// 1136239445
// }

func ExampleTimeCompare() {
	output, err := TimeCompare("2006-01-02T15:04:05Z", "2006-01-02T15:04:05-07:00")
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	output, err = TimeCompare("2006-01-02T15:04:05Z", 1136214245)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// Output:
	// -1
	// 0
}

func TestTimeParse(t *testing.T) { // lint:allow_complexity
	if os.Getenv("CI") != "" {
		t.Skip("Skipping Argon2id tests in CI environment.")
	}

	for name, tc := range testfixtures.TimeParseTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := TimeParse(tc.Input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}

func TestTimeStringCompare(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping Argon2id tests in CI environment.")
	}

	for name, tc := range testfixtures.TimeCompareStringTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := TimeCompare(tc.InputA, tc.InputB)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}

func TestTimeMixedCompare(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping Argon2id tests in CI environment.")
	}

	for name, tc := range testfixtures.TimeCompareMixedTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := TimeCompare(tc.InputA, tc.InputB)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}

func BenchmarkTimeParse(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TimeParseTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_, _ = TimeParse(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkTimeParseParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TimeParseTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = TimeParse(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkTimeCompare(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TimeCompareStringTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_, _ = TimeCompare(tc.InputA, tc.InputB) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkTimeCompareParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TimeCompareStringTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = TimeCompare(tc.InputA, tc.InputB) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkTimeCompareMixed(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TimeCompareMixedTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_, _ = TimeCompare(tc.InputA, tc.InputB) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkTimeCompareMixedParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TimeCompareMixedTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = TimeCompare(tc.InputA, tc.InputB) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzTimeParse(f *testing.F) {
	for _, tc := range testfixtures.TimeParseTestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_, _ = TimeParse(in) // lint:allow_unhandled
		},
	)
}

func FuzzTimeCompare(f *testing.F) {
	for _, tc := range testfixtures.TimeCompareStringTestTable {
		f.Add(tc.InputA, tc.InputB)
	}

	f.Fuzz(
		func(_ *testing.T, inA, inB string) {
			_, _ = TimeCompare(inA, inB) // lint:allow_unhandled
		},
	)
}

func FuzzTimeCompareMixed(f *testing.F) {
	for _, tc := range testfixtures.TimeCompareMixedTestTable {
		f.Add(tc.InputA, tc.InputB)
	}

	f.Fuzz(
		func(_ *testing.T, inA string, inB int64) {
			_, _ = TimeCompare(inA, inB) // lint:allow_unhandled
		},
	)
}
