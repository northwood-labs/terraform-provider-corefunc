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

// MapKeyIfTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var MapKeyIfTestTable = map[string]struct { // lint:no_dupe
	Condition bool
	Value     string
	Expected  string
}{
	"true boolean": {
		Condition: true,
		Value:     "true",
		Expected:  "{}",
	},
	// "false boolean": {
	// 	Condition: false,
	// 	Value:     "true",
	// 	Expected:  "{}",
	// },
	"true string": {
		Condition: true,
		Value:     `"yes"`,
		Expected:  "{}",
	},
	// "false string": {
	// 	Condition: false,
	// 	Value:     `"no"`,
	// 	Expected:  "{}",
	// },
	"true number": {
		Condition: true,
		Value:     "1.23",
		Expected:  "{}",
	},
	// "false number": {
	// 	Condition: false,
	// 	Value:     "1.23",
	// 	Expected:  "{}",
	// },
	"true list": {
		Condition: true,
		Value:     `[1, "a", true]`,
		Expected:  "{}",
	},
	// "false list": {
	// 	Condition: false,
	// 	Value:     `[1, "a", true]`,
	// 	Expected:  "{}",
	// },
	"true map": {
		Condition: true,
		Value:     `{ company="example" }`,
		Expected:  "{}",
	},
	// "false map": {
	// 	Condition: false,
	// 	Value:     `{ company="example" }`,
	// 	Expected:  "{}",
	// },
}
