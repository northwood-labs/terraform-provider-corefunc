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

import "regexp"

// Base64GunzipTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var Base64GunzipTestTable = map[string]struct {
	ExpectError *regexp.Regexp
	Input       string
	Expected    string
}{
	"test": {
		Input:    "H4sIAAAAAAAA/ypJLS4BAAAA//8BAAD//wx+f9gEAAAA",
		Expected: "test",
	},
	"hey hey, we're the monkees": {
		Input:    "H4sIAAAAAAAA/8pIrVTISK3UUShPVS9KVSjJSFXIzc/LTk0tBgQAAP//qz+dmhoAAAA=",
		Expected: "hey hey, we're the monkees",
	},
	"hey hey, we're the monkees (no padding)": {
		Input:    "H4sIAAAAAAAA/8pIrVTISK3UUShPVS9KVSjJSFXIzc/LTk0tBgQAAP//qz+dmhoAAAA",
		Expected: "hey hey, we're the monkees",
	},
	"The open source infrastructure as code tool.": {
		Input:    "H4sIAAAAAAAA/wrJSFXIL0jNUyjOLy1KTlXIzEsrSiwuKSpNLiktSlVILFZIzk9JVSjJz8/RAwQAAP//HPstqCwAAAA=",
		Expected: "The open source infrastructure as code tool.",
	},
	"The open source infrastructure as code tool. (no padding)": {
		Input:    "H4sIAAAAAAAA/wrJSFXIL0jNUyjOLy1KTlXIzEsrSiwuKSpNLiktSlVILFZIzk9JVSjJz8/RAwQAAP//HPstqCwAAAA",
		Expected: "The open source infrastructure as code tool.",
	},
}
