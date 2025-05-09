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

// StrSnakeTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrSnakeTestTable = map[string]struct { // lint:no_dupe
	Input    string
	Expected string
}{
	"blank": {
		Input:    "",
		Expected: "",
	},
	"a": {
		Input:    "a",
		Expected: "a",
	},
	"aA": {
		Input:    "aA",
		Expected: "a_a",
	},
	"aAa": {
		Input:    "aAa",
		Expected: "a_aa",
	},
	"AaA": {
		Input:    "AaA",
		Expected: "aa_a",
	},
	"test_from_snake": {
		Input:    "test_from_snake",
		Expected: "test_from_snake",
	},
	"TestFromCamel": {
		Input:    "TestFromCamel",
		Expected: "test_from_camel",
	},
	"testFromLowerCamel": {
		Input:    "testFromLowerCamel",
		Expected: "test_from_lower_camel",
	},
	"test_with_number_123": {
		Input:    "test_with_number_123",
		Expected: "test_with_number_123",
	},
	"test with number -123": {
		Input:    "test with number -123",
		Expected: "test_with_number_123",
	},
	"test with number -123.456": {
		Input:    "test with number -123.456",
		Expected: "test_with_number_123_456",
	},
	"test with number -123.456e-2": {
		Input:    "test with number -123.456e-2",
		Expected: "test_with_number_123_456e_2",
	},
	"test initialisms ram tcp ttl ascii": {
		Input:    "test initialisms ram tcp ttl ascii",
		Expected: "test_initialisms_ram_tcp_ttl_ascii",
	},
	"v3.2.2": {
		Input:    "v3.2.2",
		Expected: "v3_2_2",
	},
	"This is [an] {example}${id32}.": {
		Input:    "This is [an] {example}$${id32}.",
		Expected: "this_is_an_example_id_32",
	},
}
