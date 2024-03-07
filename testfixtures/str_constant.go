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

// StrConstantTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrConstantTestTable = map[string]struct { // lint:no_dupe
	Input    string
	Expected string
}{
	"blank": {
		Input:    "",
		Expected: "",
	},
	"a": {
		Input:    "a",
		Expected: "A",
	},
	"aA": {
		Input:    "aA",
		Expected: "A_A",
	},
	"aAa": {
		Input:    "aAa",
		Expected: "A_AA",
	},
	"AaA": {
		Input:    "AaA",
		Expected: "AA_A",
	},
	"test_from_snake": {
		Input:    "test_from_snake",
		Expected: "TEST_FROM_SNAKE",
	},
	"TestFromCamel": {
		Input:    "TestFromCamel",
		Expected: "TEST_FROM_CAMEL",
	},
	"testFromLowerCamel": {
		Input:    "testFromLowerCamel",
		Expected: "TEST_FROM_LOWER_CAMEL",
	},
	"test_with_number_123": {
		Input:    "test_with_number_123",
		Expected: "TEST_WITH_NUMBER_123",
	},
	"test with number -123": {
		Input:    "test with number -123",
		Expected: "TEST_WITH_NUMBER_123",
	},
	"test with number -123.456": {
		Input:    "test with number -123.456",
		Expected: "TEST_WITH_NUMBER_123_456",
	},
	"test with number -123.456e-2": {
		Input:    "test with number -123.456e-2",
		Expected: "TEST_WITH_NUMBER_123_456E_2",
	},
	"test initialisms ram tcp ttl ascii": {
		Input:    "test initialisms ram tcp ttl ascii",
		Expected: "TEST_INITIALISMS_RAM_TCP_TTL_ASCII",
	},
	"v3.2.2": {
		Input:    "v3.2.2",
		Expected: "V3_2_2",
	},
	"This is [an] {example}${id32}.": {
		Input:    "This is [an] {example}$${id32}.",
		Expected: "THIS_IS_AN_EXAMPLE_ID_32",
	},
}
