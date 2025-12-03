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
	// TimeParseTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	TimeParseTestTable = map[string]struct {
		Input    string
		Expected int64
	}{
		//--------------------------------------------------------------------------
		// RFC3339/ISO8601-like

		"2006-01-02": {
			Input:    "2006-01-02",
			Expected: 1136160000, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05": {
			Input:    "2006-01-02T15:04:05",
			Expected: 1136214245, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05-07:00": {
			Input:    "2006-01-02T15:04:05-07:00",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05-0700": {
			Input:    "2006-01-02T15:04:05-0700",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05.9999999-07:00": {
			Input:    "2006-01-02T15:04:05.9999999-07:00",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05.9999999-0700": {
			Input:    "2006-01-02T15:04:05.9999999-0700",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05.9999999+07:00": {
			Input:    "2006-01-02T15:04:05.9999999+07:00",
			Expected: 1136189045, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05.9999999+0700": {
			Input:    "2006-01-02T15:04:05.9999999+0700",
			Expected: 1136189045, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05.9999999Z": {
			Input:    "2006-01-02T15:04:05.9999999Z",
			Expected: 1136214245, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05+07:00": {
			Input:    "2006-01-02T15:04:05+07:00",
			Expected: 1136189045, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05+0700": {
			Input:    "2006-01-02T15:04:05+0700",
			Expected: 1136189045, // lint:allow_raw_number
		},
		"2006-01-02T15:04:05Z": {
			Input:    "2006-01-02T15:04:05Z",
			Expected: 1136214245, // lint:allow_raw_number
		},

		//--------------------------------------------------------------------------
		// Older US military time formats
		// RFC822, RFC850, RFC1036, RFC1123, RFC2822, RFC5322, RFC7231

		// RFC5322
		"Mon, 02 Jan 2006 15:04:05 GMT": {
			Input:    "Mon, 02 Jan 2006 15:04:05 GMT",
			Expected: 1136214245, // lint:allow_raw_number
		},

		// Short day, 2-digit date, spaced, 2-digit year
		"Mon, 02 Jan 06 15:04:05 MST": {
			Input:    "Mon, 02 Jan 06 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Mon, 02 Jan 06 15:04:05.9999999 MST": {
			Input:    "Mon, 02 Jan 06 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Short day, 2-digit date, spaced, 4-digit year
		"Mon, 02 Jan 2006 15:04:05 MST": {
			Input:    "Mon, 02 Jan 2006 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Mon, 02 Jan 2006 15:04:05.9999999 MST": {
			Input:    "Mon, 02 Jan 2006 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Short day, 2-digit date, hyphenated, 2-digit year
		"Mon, 02-Jan-06 15:04:05 MST": {
			Input:    "Mon, 02-Jan-06 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Mon, 02-Jan-06 15:04:05.9999999 MST": {
			Input:    "Mon, 02-Jan-06 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Short day, 2-digit date, hyphenated, 4-digit year
		"Mon, 02-Jan-2006 15:04:05 MST": {
			Input:    "Mon, 02-Jan-2006 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Mon, 02-Jan-2006 15:04:05.9999999 MST": {
			Input:    "Mon, 02-Jan-2006 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Short day, 1-digit date, hyphenated, 2-digit year
		"Mon, 2-Jan-06 15:04:05 MST": {
			Input:    "Mon, 2-Jan-06 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Mon, 2-Jan-06 15:04:05.9999999 MST": {
			Input:    "Mon, 2-Jan-06 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Short day, 1-digit date, hyphenated, 4-digit year
		"Mon, 2-Jan-2006 15:04:05 MST": {
			Input:    "Mon, 2-Jan-2006 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Mon, 2-Jan-2006 15:04:05.9999999 MST": {
			Input:    "Mon, 2-Jan-2006 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Long day, 2-digit date, spaced, 2-digit year
		"Monday, 02 Jan 06 15:04:05 MST": {
			Input:    "Monday, 02 Jan 06 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Monday, 02 Jan 06 15:04:05.9999999 MST": {
			Input:    "Monday, 02 Jan 06 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Long day, 2-digit date, spaced, 4-digit year
		"Monday, 02 Jan 2006 15:04:05 MST": {
			Input:    "Monday, 02 Jan 2006 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Monday, 02 Jan 2006 15:04:05.9999999 MST": {
			Input:    "Monday, 02 Jan 2006 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Long day, 2-digit date, hyphenated, 2-digit year
		"Monday, 02-Jan-06 15:04:05 MST": {
			Input:    "Monday, 02-Jan-06 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Monday, 02-Jan-06 15:04:05.9999999 MST": {
			Input:    "Monday, 02-Jan-06 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Long day, 2-digit date, hyphenated, 4-digit year
		"Monday, 02-Jan-2006 15:04:05 MST": {
			Input:    "Monday, 02-Jan-2006 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Monday, 02-Jan-2006 15:04:05.9999999 MST": {
			Input:    "Monday, 02-Jan-2006 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Long day, 1-digit date, hyphenated, 2-digit year
		"Monday, 2-Jan-06 15:04:05 MST": {
			Input:    "Monday, 2-Jan-06 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Monday, 2-Jan-06 15:04:05.9999999 MST": {
			Input:    "Monday, 2-Jan-06 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},

		// Long day, 1-digit date, hyphenated, 4-digit year
		"Monday, 2-Jan-2006 15:04:05 MST": {
			Input:    "Monday, 2-Jan-2006 15:04:05 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
		"Monday, 2-Jan-2006 15:04:05.9999999 MST": {
			Input:    "Monday, 2-Jan-2006 15:04:05.9999999 MST",
			Expected: 1136239445, // lint:allow_raw_number
		},
	}

	// TimeCompareStringTestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	TimeCompareStringTestTable = map[string]struct {
		InputA   string
		InputB   string
		Expected int // -1 if A < B, 0 if A == B, +1 if A > B
	}{
		"equal times": {
			InputA:   "2006-01-02T15:04:05Z",
			InputB:   "2006-01-02T15:04:05Z",
			Expected: 0,
		},
		"a before b": {
			InputA:   "2006-01-02T15:04:05Z",
			InputB:   "2007-01-02T15:04:05Z",
			Expected: -1,
		},
		"a after b": {
			InputA:   "2007-01-02T15:04:05Z",
			InputB:   "2006-01-02T15:04:05Z",
			Expected: 1,
		},
	}

	// TimeCompareMixedTestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	TimeCompareMixedTestTable = map[string]struct {
		InputA   string
		InputB   int64
		Expected int // -1 if A < B, 0 if A == B, +1 if A > B
	}{
		"equal times": {
			InputA:   "2006-01-02T15:04:05Z",
			InputB:   1136214245, // lint:allow_raw_number
			Expected: 0,
		},
		"a before b": {
			InputA:   "2006-01-02T15:04:05Z",
			InputB:   1167750245, // lint:allow_raw_number
			Expected: -1,
		},
		"a after b": {
			InputA:   "2007-01-02T15:04:05Z",
			InputB:   1136239445, // lint:allow_raw_number
			Expected: 1,
		},
	}
)
