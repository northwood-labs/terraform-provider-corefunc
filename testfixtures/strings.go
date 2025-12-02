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

var (
	// StrStartsWithTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	//
	// Ported from OpenTofu, forked from Terraform. This function is licensed as
	// MPL-2.0.
	StrStartsWithTestTable = map[string]struct {
		Input    string
		Prefix   string
		Expected bool
	}{
		"hello world": {
			Input:    "hello world",
			Prefix:   "hello",
			Expected: true,
		},
		"hey world": {
			Input:    "hey world",
			Prefix:   "hello",
			Expected: false,
		},
		"empty": {
			Input:    "",
			Prefix:   "",
			Expected: true,
		},
		"a": {
			Input:    "a",
			Prefix:   "",
			Expected: true,
		},
		"empty-a": {
			Input:    "",
			Prefix:   "a",
			Expected: false,
		},
		"unicode-prefix": {
			// Unicode combining characters edge-case: we match the prefix
			// in terms of unicode code units rather than grapheme clusters,
			// which is inconsistent with our string processing elsewhere but
			// would be a breaking change to fix that bug now.

			// "Man Shrugging" is encoded as "Person Shrugging" followed by
			// zero-width joiner and then the masculine gender presentation
			// modifier.
			Input: "\U0001f937\u200d\u2642",
			// Just the "Person Shrugging" character without any modifiers.
			Prefix:   "\U0001f937",
			Expected: true,
		},
	}

	// StrEndsWithTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	StrEndsWithTestTable = map[string]struct {
		Input    string
		Suffix   string
		Expected bool
	}{
		"hello world": {
			Input:    "hello world",
			Suffix:   "world",
			Expected: true,
		},
		"hey world": {
			Input:    "hey wood",
			Suffix:   "world",
			Expected: false,
		},
		"empty": {
			Input:    "",
			Suffix:   "",
			Expected: true,
		},
		"a": {
			Input:    "a",
			Suffix:   "",
			Expected: true,
		},
		"empty-a": {
			Input:    "",
			Suffix:   "a",
			Expected: false,
		},
		"unicode-prefix": {
			// Unicode combining characters edge-case: we match the prefix
			// in terms of unicode code units rather than grapheme clusters,
			// which is inconsistent with our string processing elsewhere but
			// would be a breaking change to fix that bug now.

			// "Man Shrugging" is encoded as "Person Shrugging" followed by
			// zero-width joiner and then the masculine gender presentation
			// modifier.
			Input: "\U0001f937\u200d\u2642",
			// Just the "masculine gender presentation modifier" character.
			Suffix:   "\u2642",
			Expected: true,
		},
	}

	// StrContainsTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	//
	// Ported from OpenTofu, forked from Terraform. This function is licensed as
	// MPL-2.0.
	StrContainsTestTable = map[string]struct { // lint:no_dupe
		Input, Substr string
		Expected      bool
	}{
		"hello-hel": {
			Input:    "hello",
			Substr:   "hel",
			Expected: true,
		},
		"hello-lo": {
			"hello",
			"lo",
			true,
		},
		"hello1-1": {
			"hello1",
			"1",
			true,
		},
		"hello1-heo": {
			"hello1",
			"heo",
			false,
		},
	}
)
