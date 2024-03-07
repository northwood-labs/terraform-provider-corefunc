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

// StrPascalTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrPascalTestTable = map[string]struct {
	Input       string
	Expected    string
	AcronymCaps bool
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
		Expected: "AA",
	},
	"aAa": {
		Input:    "aAa",
		Expected: "AAa",
	},
	"AaA": {
		Input:    "AaA",
		Expected: "AaA",
	},
	"test_from_snake": {
		Input:    "test_from_snake",
		Expected: "TestFromSnake",
	},
	"TestFromCamel": {
		Input:    "TestFromCamel",
		Expected: "TestFromCamel",
	},
	"testFromLowerCamel": {
		Input:    "testFromLowerCamel",
		Expected: "TestFromLowerCamel",
	},
	"test_with_number_123": {
		Input:    "test_with_number_123",
		Expected: "TestWithNumber123",
	},
	"test with number -123": {
		Input:    "test with number -123",
		Expected: "TestWithNumber123",
	},
	"test with number -123.456": {
		Input:    "test with number -123.456",
		Expected: "TestWithNumber123456",
	},
	"test with number -123.456e-2": {
		Input:    "test with number -123.456e-2",
		Expected: "TestWithNumber123456e2",
	},
	"test initialisms ip html eof ascii cpu (AcronymCaps)": {
		Input:       "test initialisms ip html eof ascii cpu",
		AcronymCaps: true,
		Expected:    "TestInitialismsIPHTMLEOFASCIICPU",
	},
	"test initialisms ip html eof ascii cpu": {
		Input:    "test initialisms ip html eof ascii cpu",
		Expected: "TestInitialismsIpHtmlEofAsciiCpu",
	},
	"v3.2.2": {
		Input:    "v3.2.2",
		Expected: "V322",
	},
	"This is [an] {example}${id32}. (AcronymCaps)": {
		Input:       "This is [an] {example}$${id32}.",
		AcronymCaps: true,
		Expected:    "ThisIsAnExampleID32",
	},
	"This is [an] {example}${id32}.": {
		Input:    "This is [an] {example}$${id32}.",
		Expected: "ThisIsAnExampleId32",
	},
	"entity id": {
		Input:       "entity id",
		AcronymCaps: true,
		Expected:    "EntityID",
	},
}
