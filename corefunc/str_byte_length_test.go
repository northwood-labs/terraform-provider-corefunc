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

func ExampleStrByteLength() {
	output := StrByteLength("abcde")
	fmt.Println(output)

	output = StrByteLength("\x00")
	fmt.Println(output)

	output = StrByteLength("👁")
	fmt.Println(output)

	output = StrByteLength("■㈱の世界①")
	fmt.Println(output)

	// Output:
	// 5
	// 0
	// 1
	// 10
}

func TestStrByteLength(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.StrByteLengthTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrByteLength(tc.Input)

			if output != tc.Expected {
				t.Errorf("Expected %d, got %d", tc.Expected, output)
			}
		})
	}
}
