// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleStrIterativeReplace() {
	output := StrIterativeReplace(
		"This is a string for testing replacements. New Relic. Set-up.",
		[]types.Replacement{
			{Old: ".", New: ""},
			{Old: " ", New: "_"},
			{Old: "-", New: "_"},
			{Old: "New_Relic", New: "datadog"},
			{Old: "This", New: "this"},
			{Old: "Set_up", New: "setup"},
		},
	)

	fmt.Println(output)

	// Output:
	// this_is_a_string_for_testing_replacements_datadog_setup
}

func TestStrIterativeReplace(t *testing.T) {
	for name, tc := range testfixtures.StrIterativeReplaceTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrIterativeReplace(tc.Input, tc.Replacements)

			if output != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, output)
			}
		})
	}
}

func BenchmarkStrIterativeReplace(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrIterativeReplaceTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrIterativeReplace(tc.Input, tc.Replacements) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrIterativeReplaceParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrIterativeReplaceTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrIterativeReplace(tc.Input, tc.Replacements) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzStrIterativeReplace(f *testing.F) {
	for _, tc := range testfixtures.StrIterativeReplaceTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)

		for _, r := range tc.Replacements {
			f.Add(r.Old)
			f.Add(r.New)
		}
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrIterativeReplace(in, []types.Replacement{ // lint:allow_unhandled
				{Old: in, New: in},
			})
		},
	)
}
