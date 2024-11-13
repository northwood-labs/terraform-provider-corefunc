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

package testfixtures // lint:no_dupe

import (
	"regexp"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

// DeepMergeTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var DeepMergeTestTable = map[string]struct {
	Expected    any
	ExpectedErr *regexp.Regexp
	Input       []any
	Config      types.MergeConfig
}{
	"blank": {
		Input:       nil,
		Expected:    nil,
		Config:      types.MergeConfig{},
		ExpectedErr: &regexp.Regexp{},
	},
	"abcd": {
		Input:       []any{"a", "b", "c", "d"},
		Expected:    "d",
		Config:      types.MergeConfig{},
		ExpectedErr: &regexp.Regexp{},
	},
	"1234": {
		Input:       []any{1, 2, 3, 4},
		Expected:    4, // lint:allow_raw_number
		Config:      types.MergeConfig{},
		ExpectedErr: &regexp.Regexp{},
	},
	"1a2b": {
		Input:       []any{1, "a", 2, "b"},
		Expected:    nil,
		Config:      types.MergeConfig{},
		ExpectedErr: regexp.MustCompile("all elements must be of the same type"),
	},
}
