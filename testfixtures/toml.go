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

package testfixtures // lint:no_dupe

import _ "embed"

var (
	//go:embed toml.example.json
	jsonContent string

	//go:embed toml.example2.json
	json2Content string

	//go:embed toml.example.toml
	tomlContent string

	//go:embed toml.example2.toml
	toml2Content string

	// TOMLtoJSONTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	TOMLtoJSONTestTable = map[string]struct { // lint:no_dupe
		Input       string
		Expected    string
		ExpectedErr bool
	}{
		"abc=123": {
			Input:       `abc = 123`,
			Expected:    `{"abc":123}`,
			ExpectedErr: false,
		},
		"toml.example.toml": {
			Input:       tomlContent,
			Expected:    jsonContent,
			ExpectedErr: false,
		},
		"toml.example2.toml": {
			Input:       toml2Content,
			Expected:    json2Content,
			ExpectedErr: false,
		},
	}

	// JSONtoTOMLTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	JSONtoTOMLTestTable = map[string]struct { // lint:no_dupe
		Input       string
		Expected    string
		ExpectedErr bool
	}{
		`{"abc":123}`: {
			Input:       `{"abc": 123}`,
			Expected:    `abc = 123`,
			ExpectedErr: false,
		},
		"toml.example.json": {
			Input:       jsonContent,
			Expected:    tomlContent,
			ExpectedErr: false,
		},
		"toml.example2.json": {
			Input:       json2Content,
			Expected:    toml2Content,
			ExpectedErr: false,
		},
	}
)
