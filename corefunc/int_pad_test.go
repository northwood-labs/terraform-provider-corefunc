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
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleIntLeftPad() {
	output1 := IntLeftPad(255, 5)
	fmt.Println(output1)

	output2 := IntLeftPad(0b11111111, 5)
	fmt.Println(output2)

	output3 := IntLeftPad(0xFF, 5)
	fmt.Println(output3)

	// Output:
	// 00255
	// 00255
	// 00255
}

func TestIntLeftPad(t *testing.T) {
	for name, tc := range testfixtures.IntLeftPadTestTable {
		t.Run(name, func(t *testing.T) {
			output := IntLeftPad(tc.Input, tc.PadWidth)

			if output != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, output)
			}
		})
	}
}

func BenchmarkIntLeftPad(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.IntLeftPadTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = IntLeftPad(tc.Input, tc.PadWidth) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkIntLeftPadParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.IntLeftPadTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = IntLeftPad(tc.Input, tc.PadWidth) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzIntLeftPad(f *testing.F) {
	for _, tc := range testfixtures.IntLeftPadTestTable {
		f.Add(tc.Input)
		f.Add(int64(tc.PadWidth))
	}

	f.Fuzz(
		func(t *testing.T, in int64) {
			_ = IntLeftPad(in, int(in)) // lint:allow_unhandled
		},
	)
}
