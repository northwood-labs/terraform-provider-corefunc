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
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func TestHomedir(t *testing.T) {
	for name, tc := range testfixtures.HomedirGetTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := Homedir()
			if err != nil {
				t.Error("Failed to lookup the home directory")
			}

			if output != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, output)
			}
		})
	}
}

func BenchmarkHomedir(b *testing.B) {
	b.ReportAllocs()

	for name := range testfixtures.HomedirGetTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, _ = Homedir() // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHomedirParallel(b *testing.B) {
	b.ReportAllocs()

	for name := range testfixtures.HomedirGetTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = Homedir() // lint:allow_unhandled
				}
			})
		})
	}
}
