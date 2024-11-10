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

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleCIDRContains() {
	output, err := CIDRContains("192.168.2.0/20", "192.168.2.1")
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// Output:
	// true
}

func TestCIDRContains(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.CIDRContainsTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := CIDRContains(tc.ContainerCidr, tc.ContainedIPOrCidr)

			// We expect an error.
			if err != nil && !tc.ExpectedErr {
				t.Errorf("Unexpected error: %v", err)
			}

			if output != tc.Expected {
				t.Errorf("Expected %v, got %v", tc.Expected, output)
			}
		})
	}
}
