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

var (
	// URLParseTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	URLParseTestTable = map[string]struct { // lint:no_dupe
		Canonicalizer string
		InputURL      string
		Href          string
		Protocol      string
		Scheme        string
		Username      string
		Password      string
		Hostname      string
		Host          string
		Port          string
		Path          string
		Search        string
		Query         string
		Hash          string
		Fragment      string
		DecodedPort   int
	}{
		// ---------------------------------------------------------------------
		// No canonicalizer specified.

		"DEFAULT: HTTP://u:p@example.com:80/foo?q=1#bar": {
			InputURL:    "HTTP://u:p@example.com:80/foo?q=1#bar",
			Href:        "http://u:p@example.com/foo?q=1#bar",
			Protocol:    "http:",
			Scheme:      "http",
			Username:    "u",
			Password:    "p",
			Hostname:    "example.com",
			Host:        "example.com",
			Port:        "",
			Path:        "/foo",
			Search:      "?q=1",
			Query:       "q=1",
			Hash:        "#bar",
			Fragment:    "bar",
			DecodedPort: 80, // lint:allow_raw_number
		},
		"DEFAULT: HTTP://u:p@example.com/foo?q=1#bar": {
			InputURL:    "HTTP://u:p@example.com/foo?q=1#bar",
			Href:        "http://u:p@example.com/foo?q=1#bar",
			Protocol:    "http:",
			Scheme:      "http",
			Username:    "u",
			Password:    "p",
			Hostname:    "example.com",
			Host:        "example.com",
			Port:        "",
			Path:        "/foo",
			Search:      "?q=1",
			Query:       "q=1",
			Hash:        "#bar",
			Fragment:    "bar",
			DecodedPort: 80, // lint:allow_raw_number
		},
		"DEFAULT: HTTP://u:p@example.com:8080/foo?q=1#bar": {
			InputURL:    "HTTP://u:p@example.com:8080/foo?q=1#bar",
			Href:        "http://u:p@example.com:8080/foo?q=1#bar",
			Protocol:    "http:",
			Scheme:      "http",
			Username:    "u",
			Password:    "p",
			Hostname:    "example.com",
			Host:        "example.com:8080",
			Port:        "8080",
			Path:        "/foo",
			Search:      "?q=1",
			Query:       "q=1",
			Hash:        "#bar",
			Fragment:    "bar",
			DecodedPort: 8080, // lint:allow_raw_number
		},
		"DEFAULT: HTTPs://example.com": {
			InputURL:    "HTTPs://example.com",
			Href:        "https://example.com/",
			Protocol:    "https:",
			Scheme:      "https",
			Username:    "",
			Password:    "",
			Hostname:    "example.com",
			Host:        "example.com",
			Port:        "",
			Path:        "/",
			Search:      "",
			Query:       "",
			Hash:        "",
			Fragment:    "",
			DecodedPort: 443, // lint:allow_raw_number
		},

		// ---------------------------------------------------------------------
		// Standard canonicalizer specified.

		"STANDARD: HTTP://u:p@example.com:80/foo?q=1#bar": {
			Canonicalizer: "standard",
			InputURL:      "HTTP://u:p@example.com:80/foo?q=1#bar",
			Href:          "http://u:p@example.com/foo?q=1#bar",
			Protocol:      "http:",
			Scheme:        "http",
			Username:      "u",
			Password:      "p",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/foo",
			Search:        "?q=1",
			Query:         "q=1",
			Hash:          "#bar",
			Fragment:      "bar",
			DecodedPort:   80, // lint:allow_raw_number
		},
		"STANDARD: HTTP://u:p@example.com/foo?q=1#bar": {
			Canonicalizer: "standard",
			InputURL:      "HTTP://u:p@example.com/foo?q=1#bar",
			Href:          "http://u:p@example.com/foo?q=1#bar",
			Protocol:      "http:",
			Scheme:        "http",
			Username:      "u",
			Password:      "p",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/foo",
			Search:        "?q=1",
			Query:         "q=1",
			Hash:          "#bar",
			Fragment:      "bar",
			DecodedPort:   80, // lint:allow_raw_number
		},
		"STANDARD: HTTP://u:p@example.com:8080/foo?q=1#bar": {
			Canonicalizer: "standard",
			InputURL:      "HTTP://u:p@example.com:8080/foo?q=1#bar",
			Href:          "http://u:p@example.com:8080/foo?q=1#bar",
			Protocol:      "http:",
			Scheme:        "http",
			Username:      "u",
			Password:      "p",
			Hostname:      "example.com",
			Host:          "example.com:8080",
			Port:          "8080",
			Path:          "/foo",
			Search:        "?q=1",
			Query:         "q=1",
			Hash:          "#bar",
			Fragment:      "bar",
			DecodedPort:   8080, // lint:allow_raw_number
		},
		"STANDARD: HTTPs://example.com": {
			Canonicalizer: "standard",
			InputURL:      "HTTPs://example.com",
			Href:          "https://example.com/",
			Protocol:      "https:",
			Scheme:        "https",
			Username:      "",
			Password:      "",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/",
			Search:        "",
			Query:         "",
			Hash:          "",
			Fragment:      "",
			DecodedPort:   443, // lint:allow_raw_number
		},

		// ---------------------------------------------------------------------
		// GoogleSafe canonicalizer specified.

		"GOOGLE SAFE: HTTP://u:p@example.com:80/foo?q=1#bar": {
			Canonicalizer: "google_safe_browsing",
			InputURL:      "HTTP://u:p@example.com:80/foo?q=1#bar",
			Href:          "http://u:p@example.com/foo?q=1",
			Protocol:      "http:",
			Scheme:        "http",
			Username:      "u",
			Password:      "p",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/foo",
			Search:        "?q=1",
			Query:         "q=1",
			Hash:          "",
			Fragment:      "",
			DecodedPort:   80, // lint:allow_raw_number
		},
		"GOOGLE SAFE: HTTP://u:p@example.com/foo?q=1#bar": {
			Canonicalizer: "google_safe_browsing",
			InputURL:      "HTTP://u:p@example.com/foo?q=1#bar",
			Href:          "http://u:p@example.com/foo?q=1",
			Protocol:      "http:",
			Scheme:        "http",
			Username:      "u",
			Password:      "p",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/foo",
			Search:        "?q=1",
			Query:         "q=1",
			Hash:          "",
			Fragment:      "",
			DecodedPort:   80, // lint:allow_raw_number
		},
		"GOOGLE SAFE: HTTP://u:p@example.com:8080/foo?q=1#bar": {
			Canonicalizer: "google_safe_browsing",
			InputURL:      "HTTP://u:p@example.com:8080/foo?q=1#bar",
			Href:          "http://u:p@example.com/foo?q=1",
			Protocol:      "http:",
			Scheme:        "http",
			Username:      "u",
			Password:      "p",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/foo",
			Search:        "?q=1",
			Query:         "q=1",
			Hash:          "",
			Fragment:      "",
			DecodedPort:   8080, // lint:allow_raw_number
		},
		"GOOGLE SAFE: HTTPs://example.com": {
			Canonicalizer: "google_safe_browsing",
			InputURL:      "HTTPs://example.com",
			Href:          "https://example.com/",
			Protocol:      "https:",
			Scheme:        "https",
			Username:      "",
			Password:      "",
			Hostname:      "example.com",
			Host:          "example.com",
			Port:          "",
			Path:          "/",
			Search:        "",
			Query:         "",
			Hash:          "",
			Fragment:      "",
			DecodedPort:   443, // lint:allow_raw_number
		},
	}

	// URLDecodeTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	URLDecodeTestTable = map[string]struct { // lint:no_dupe
		Input       string
		Expected    string
		ExpectedErr bool
	}{
		"abc123-_": {
			Input:       "abc123-_",
			Expected:    "abc123-_",
			ExpectedErr: false,
		},
		"foo%3Abar%40localhost%3Ffoo%3Dbar%26bar%3Dbaz": {
			Input:       "foo%3Abar%40localhost%3Ffoo%3Dbar%26bar%3Dbaz",
			Expected:    "foo:bar@localhost?foo=bar&bar=baz",
			ExpectedErr: false,
		},
		"mailto%3Aemail%3Fsubject%3Dthis%2Bis%2Bmy%2Bsubject": {
			Input:       "mailto%3Aemail%3Fsubject%3Dthis%2Bis%2Bmy%2Bsubject",
			Expected:    "mailto:email?subject=this+is+my+subject",
			ExpectedErr: false,
		},
		"foo%2Fbar": {
			Input:       "foo%2Fbar",
			Expected:    "foo/bar",
			ExpectedErr: false,
		},
		"foo% bar": {
			Input:       "foo% bar",
			Expected:    "",
			ExpectedErr: true,
		},
		"foo%2 bar": {
			Input:       "foo%2 bar",
			Expected:    "",
			ExpectedErr: true,
		},
		"%GGfoo%2bar": {
			Input:       "%GGfoo%2bar",
			Expected:    "",
			ExpectedErr: true,
		},
		"foo%00, bar!": {
			Input:       "foo%00, bar!",
			Expected:    "foo\x00, bar!",
			ExpectedErr: false,
		},
		"hello%20%E4%B8%96%E7%95%8C": {
			Input:       "hello%20%E4%B8%96%E7%95%8C", // Unicode character support
			Expected:    "hello 世界",
			ExpectedErr: false,
		},
		"hello%20%D8%AF%D9%86%DB%8C%D8%A7": {
			Input:       "hello%20%D8%AF%D9%86%DB%8C%D8%A7", // Unicode character support
			Expected:    "hello دنیا",
			ExpectedErr: false,
		},
	}
)
