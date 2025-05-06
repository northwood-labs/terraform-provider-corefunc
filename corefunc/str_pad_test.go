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

func ExampleStrLeftPad() {
	output1 := StrLeftPad("foobar", 10)
	fmt.Println(output1)

	output2 := StrLeftPad("foobar", 10, '_')
	fmt.Println(output2)

	// Output:
	//     foobar
	// ____foobar
}

func TestStrLeftPad(t *testing.T) {
	for name, tc := range testfixtures.StrLeftPadTestTable {
		t.Run(name, func(t *testing.T) {
			output := ""

			var emptyByte byte

			if tc.PadChar == emptyByte {
				output = StrLeftPad(tc.Input, tc.PadWidth)
			} else {
				output = StrLeftPad(tc.Input, tc.PadWidth, tc.PadChar)
			}

			if output != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, output)
			}
		})
	}
}

func BenchmarkStrLeftPad(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrLeftPadTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrLeftPad(tc.Input, tc.PadWidth, tc.PadChar) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrLeftPadParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrLeftPadTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrLeftPad(tc.Input, tc.PadWidth, tc.PadChar) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzStrLeftPad(f *testing.F) {
	for _, tc := range testfixtures.StrLeftPadTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrLeftPad(in, len(in)) // lint:allow_unhandled
		},
	)
}
