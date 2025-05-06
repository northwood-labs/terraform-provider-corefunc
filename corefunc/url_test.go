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

package corefunc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleURLParse() {
	parsedURL, err := URLParse("HTTP://u:p@example.com:80/foo?q=1#bar")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedURL.Href(false))
	fmt.Println(parsedURL.Href(true))
	fmt.Println(parsedURL.Protocol())
	fmt.Println(parsedURL.Scheme())
	fmt.Println(parsedURL.Username())
	fmt.Println(parsedURL.Password())
	fmt.Println(parsedURL.Host())
	fmt.Println(parsedURL.Port())
	fmt.Println(parsedURL.Pathname())
	fmt.Println(parsedURL.Search())
	fmt.Println(parsedURL.Query())
	fmt.Println(parsedURL.Hash())
	fmt.Println(parsedURL.Fragment())
	fmt.Println(parsedURL.DecodedPort())

	// Output:
	// http://u:p@example.com/foo?q=1#bar
	// http://u:p@example.com/foo?q=1
	// http:
	// http
	// u
	// p
	// example.com
	//
	// /foo
	// ?q=1
	// q=1
	// #bar
	// bar
	// 80
}

func ExampleURLParse_googleSafeBrowsing() {
	// https://developers.google.com/safe-browsing/v4/urls-hashing#canonicalization
	parsedURL, err := URLParse("HTTP://u:p@example.com:80/foo?q=1#bar", types.GoogleSafeBrowsing)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("." + parsedURL.Href(false))
	fmt.Println("." + parsedURL.Href(true))
	fmt.Println("." + parsedURL.Protocol())
	fmt.Println("." + parsedURL.Scheme())
	fmt.Println("." + parsedURL.Username())
	fmt.Println("." + parsedURL.Password())
	fmt.Println("." + parsedURL.Host())
	fmt.Println("." + parsedURL.Port())
	fmt.Println("." + parsedURL.Pathname())
	fmt.Println("." + parsedURL.Search())
	fmt.Println("." + parsedURL.Query())
	fmt.Println("." + parsedURL.Hash())
	fmt.Println("." + parsedURL.Fragment())
	fmt.Println("." + fmt.Sprint(parsedURL.DecodedPort()))

	// Output:
	// .http://u:p@example.com/foo?q=1
	// .http://u:p@example.com/foo?q=1
	// .http:
	// .http
	// .u
	// .p
	// .example.com
	// .
	// ./foo
	// .?q=1
	// .q=1
	// .
	// .
	// .80
}

func ExampleURLDecode() {
	output, err := URLDecode("hello%20%E4%B8%96%E7%95%8C")
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// Output:
	// hello 世界
}

func TestURLParse(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.URLParseTestTable {
		t.Run(name, func(t *testing.T) {
			// Handle canonicalizer-switching
			opts := types.Standard

			if strings.EqualFold(tc.Canonicalizer, "google_safe_browsing") {
				opts = types.GoogleSafeBrowsing
			}

			u, err := URLParse(tc.InputURL, opts)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got := u.Href(false); got != tc.Href {
				t.Errorf("Href() = %v, want %v", got, tc.Href)
			}

			if got := u.Protocol(); got != tc.Protocol {
				t.Errorf("Protocol() = %v, want %v", got, tc.Protocol)
			}

			if got := u.Scheme(); got != tc.Scheme {
				t.Errorf("Scheme() = %v, want %v", got, tc.Scheme)
			}

			if got := u.Username(); got != tc.Username {
				t.Errorf("Username() = %v, want %v", got, tc.Username)
			}

			if got := u.Password(); got != tc.Password {
				t.Errorf("Password() = %v, want %v", got, tc.Password)
			}

			if got := u.Hostname(); got != tc.Hostname {
				t.Errorf("Hostname() = %v, want %v", got, tc.Host)
			}

			if got := u.Host(); got != tc.Host {
				t.Errorf("Host() = %v, want %v", got, tc.Host)
			}

			if got := u.Port(); got != tc.Port {
				t.Errorf("Port() = %v, want %v", got, tc.Port)
			}

			if got := u.DecodedPort(); got != tc.DecodedPort {
				t.Errorf("DecodedPort() = %v, want %v", got, tc.DecodedPort)
			}

			if got := u.Pathname(); got != tc.Path {
				t.Errorf("Pathname() = %v, want %v", got, tc.Path)
			}

			if got := u.Search(); got != tc.Search {
				t.Errorf("Search() = %v, want %v", got, tc.Search)
			}

			if got := u.Query(); got != tc.Query {
				t.Errorf("Query() = %v, want %v", got, tc.Query)
			}

			if got := u.Hash(); got != tc.Hash {
				t.Errorf("Hash() = %v, want %v", got, tc.Hash)
			}

			if got := u.Fragment(); got != tc.Fragment {
				t.Errorf("Fragment() = %v, want %v", got, tc.Fragment)
			}
		})
	}
}

func TestURLDecode(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.URLDecodeTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := URLDecode(tc.Input)

			// We expect an error.
			if err != nil && !tc.ExpectedErr {
				t.Errorf("Unexpected error: %v", err)
			}

			if output != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, output)
			}
		})
	}
}
