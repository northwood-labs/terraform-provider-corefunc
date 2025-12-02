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

func ExampleStrStringContains() {
	output := StrStringContains("abcde", "bcd")
	fmt.Println(output)

	output = StrStringContains("abcde", "xyz")
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

func TestStrStringContains(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.StrContainsTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrStringContains(tc.Input, tc.Substr)

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}
