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

package testfixtures // lint:no_dupe

import "github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"

// StrIterativeReplaceTestTable is used by both the standard Go tests and also
// the Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrIterativeReplaceTestTable = map[string]struct { // lint:no_dupe
	Input        string
	Expected     string
	Replacements []types.Replacement
}{
	"replacements": {
		Input: "This is a string for testing replacements. New Relic. Set-up.",
		Replacements: []types.Replacement{
			{Old: ".", New: ""},
			{Old: " ", New: "_"},
			{Old: "-", New: "_"},
			{Old: "New_Relic", New: "datadog"},
			{Old: "This", New: "this"},
			{Old: "Set_up", New: "setup"},
		},
		Expected: "this_is_a_string_for_testing_replacements_datadog_setup",
	},
}
