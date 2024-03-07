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
	"strings"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/google/go-cmp/cmp"
)

func ExampleTruncateLabel() {
	length1 := TruncateLabel(1, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length1)

	length3 := TruncateLabel(3, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length3)

	length5 := TruncateLabel(5, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length5)

	length7 := TruncateLabel(7, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length7)

	length10 := TruncateLabel(10, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length10)

	length20 := TruncateLabel(20, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length20)

	length32 := TruncateLabel(32, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length32)

	length64 := TruncateLabel(64, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length64)

	length128 := TruncateLabel(128, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length128)

	length256 := TruncateLabel(256, "NW-TEST-FIRST", "Name of my monitor")
	fmt.Println(length256)

	// Output:
	// …
	// …
	// …: …
	// …: Nam…
	// N…: Name…
	// NW-TES…: Name of my…
	// NW-TEST-FIRS…: Name of my monit…
	// NW-TEST-FIRST: Name of my monitor
	// NW-TEST-FIRST: Name of my monitor
	// NW-TEST-FIRST: Name of my monitor
}

func TestTruncateLabel(t *testing.T) {
	for name, tc := range testfixtures.TruncateLabelTestTable {
		t.Run(name, func(t *testing.T) {
			actual := TruncateLabel(tc.MaxLength, tc.Prefix, tc.Label)
			diff := cmp.Diff(tc.Expected, actual)

			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkTruncateLabel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TruncateLabelTestTable {
		b.ResetTimer()
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TruncateLabel(tc.MaxLength, tc.Prefix, tc.Label)
			}
		})
	}
}

func BenchmarkTruncateLabelParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.TruncateLabelTestTable {
		b.ResetTimer()
		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					TruncateLabel(tc.MaxLength, tc.Prefix, tc.Label)
				}
			})
		})
	}
}

func FuzzTruncateLabel(f *testing.F) {
	for _, tc := range testfixtures.TruncateLabelTestTable {
		f.Add( // Use f.Add to provide a seed corpus
			fmt.Sprintf("%s:::::%s", tc.Prefix, tc.Label),
		)
	}

	f.Fuzz(
		func(t *testing.T, input string) {
			if !strings.Contains(input, ":::::") {
				t.Skip("skipping test; input missing ':::::' delimiter")
			}

			pieces := strings.Split(input, ":::::")

			_ = TruncateLabel(0, pieces[0], pieces[1])
			_ = TruncateLabel(3, pieces[0], pieces[1])
			_ = TruncateLabel(5, pieces[0], pieces[1])
			_ = TruncateLabel(7, pieces[0], pieces[1])
			_ = TruncateLabel(8, pieces[0], pieces[1])
			_ = TruncateLabel(10, pieces[0], pieces[1])
			_ = TruncateLabel(20, pieces[0], pieces[1])
			_ = TruncateLabel(30, pieces[0], pieces[1])
			_ = TruncateLabel(32, pieces[0], pieces[1])
			_ = TruncateLabel(33, pieces[0], pieces[1])
			_ = TruncateLabel(64, pieces[0], pieces[1])
			_ = TruncateLabel(80, pieces[0], pieces[1])
			_ = TruncateLabel(100, pieces[0], pieces[1])
			_ = TruncateLabel(120, pieces[0], pieces[1])
			_ = TruncateLabel(128, pieces[0], pieces[1])
		},
	)
}
