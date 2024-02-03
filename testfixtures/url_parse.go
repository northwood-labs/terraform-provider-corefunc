// Copyright 2023-2024, Ryan Parman
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

// URLParseTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var URLParseTestTable = map[string]struct { // lint:no_dupe
	InputURL    string
	Href        string
	Protocol    string
	Scheme      string
	Username    string
	Password    string
	Hostname    string
	Host        string
	Port        string
	Path        string
	Search      string
	Query       string
	Hash        string
	Fragment    string
	DecodedPort int
}{
	"HTTP://u:p@example.com:80/foo?q=1#bar": {
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
	"HTTP://u:p@example.com/foo?q=1#bar": {
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
	"HTTP://u:p@example.com:8080/foo?q=1#bar": {
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
	"HTTPs://example.com": {
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
}

//   {
//     "input": "fass.de",
//     "output": "fass.de"
//   },
//   {
//     "input": "fa\u00df.de",
//     "output": "xn--fa-hia.de"
//   },
//   {
//     "input": "Fa\u00df.de",
//     "output": "xn--fa-hia.de"
//   },
//   {
//     "input": "xn--fa-hia.de",
//     "output": "xn--fa-hia.de"
//   },
//   {
//     "input": "\u00e0.\u05d0\u0308",
//     "output": "xn--0ca.xn--ssa73l"
//   },
//   {
//     "input": "a\u0300.\u05d0\u0308",
//     "output": "xn--0ca.xn--ssa73l"
//   },
//   {
//     "input": "A\u0300.\u05d0\u0308",
//     "output": "xn--0ca.xn--ssa73l"
//   },
//   {
//     "input": "\u00c0.\u05d0\u0308",
//     "output": "xn--0ca.xn--ssa73l"
//   },
//   {
//     "input": "xn--0ca.xn--ssa73l",
//     "output": "xn--0ca.xn--ssa73l"
//   },
//  {
//     "input": "a\u094d\u200db",
//     "output": "xn--ab-fsf014u"
//   },
//   {
//     "input": "A\u094d\u200dB",
//     "output": "xn--ab-fsf014u"
//   },
//   {
//     "input": "A\u094d\u200db",
//     "output": "xn--ab-fsf014u"
//   },
//   {
//     "input": "xn--ab-fsf014u",
//     "output": "xn--ab-fsf014u"
//   },
