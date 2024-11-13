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
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleDeepMerge() {
}

func TestDeepMerge(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.DeepMergeTestTable {
		t.Run(name, func(t *testing.T) {
			var (
				output any
				err    error
			)

			output, err = DeepMerge(tc.Config, tc.Input)
			if err != nil {
				if tc.ExpectedErr.MatchString(err.Error()) {
					return
				} else {
					t.Fatalf("Unexpected error %#v, got %#v", tc.ExpectedErr, err)
				}
			}

			if output != tc.Expected {
				t.Errorf("Expected %#v, got %#v", tc.Expected, output)
			}
		})
	}
}
