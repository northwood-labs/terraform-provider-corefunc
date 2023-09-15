// Copyright 2023, Ryan Parman
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

package testfixtures

// StrCamelTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrCamelTestTable = map[string]struct {
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
		Expected: "aA",
	},
	"aAa": {
		Input:    "aAa",
		Expected: "aAa",
	},
	"AaA": {
		Input:    "AaA",
		Expected: "aaA",
	},
	"test_from_snake": {
		Input:    "test_from_snake",
		Expected: "testFromSnake",
	},
	"TestFromCamel": {
		Input:    "TestFromCamel",
		Expected: "testFromCamel",
	},
	"testFromLowerCamel": {
		Input:    "testFromLowerCamel",
		Expected: "testFromLowerCamel",
	},
	"test_with_number_123": {
		Input:    "test_with_number_123",
		Expected: "testWithNumber123",
	},
	"test with number -123": {
		Input:    "test with number -123",
		Expected: "testWithNumber123",
	},
	"test with number -123.456": {
		Input:    "test with number -123.456",
		Expected: "testWithNumber123456",
	},
	"test with number -123.456e-2": {
		Input:    "test with number -123.456e-2",
		Expected: "testWithNumber123456e2",
	},
	"test initialisms ip html eof ascii cpu": {
		Input:    "test initialisms ip html eof ascii cpu",
		Expected: "testInitialismsIPHTMLEOFASCIICPU",
	},
	"v3.2.2": {
		Input:    "v3.2.2",
		Expected: "v322",
	},
	"This is [an] {example}${id32}.": {
		Input:    "This is [an] {example}$${id32}.",
		Expected: "thisIsAnExampleID32",
	},
	"entity id": {
		Input:    "entity id",
		Expected: "entityID",
	},
}
