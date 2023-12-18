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

// IntLeftPadTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var IntLeftPadTestTable = map[string]struct { // lint:no_dupe
	Expected string
	Input    int64
	PadWidth int
}{
	"Xn1": {
		Input:    123, // lint:allow_raw_number
		PadWidth: -1,  // lint:allow_raw_number
		Expected: "123",
	},
	"X0": {
		Input:    123, // lint:allow_raw_number
		PadWidth: 0,   // lint:allow_raw_number
		Expected: "123",
	},
	"X1": {
		Input:    123, // lint:allow_raw_number
		PadWidth: 1,   // lint:allow_raw_number
		Expected: "123",
	},
	"X2": {
		Input:    123, // lint:allow_raw_number
		PadWidth: 2,   // lint:allow_raw_number
		Expected: "123",
	},
	"X3": {
		Input:    123, // lint:allow_raw_number
		PadWidth: 3,   // lint:allow_raw_number
		Expected: "123",
	},
	"X4": {
		Input:    123, // lint:allow_raw_number
		PadWidth: 4,   // lint:allow_raw_number
		Expected: "0123",
	},
	"X5": {
		Input:    123, // lint:allow_raw_number
		PadWidth: 5,   // lint:allow_raw_number
		Expected: "00123",
	},
	"Xb17": {
		Input:    0b00010001, // lint:allow_raw_number
		PadWidth: 5,          // lint:allow_raw_number
		Expected: "00017",
	},
	"Xb255": {
		Input:    0b11111111, // lint:allow_raw_number
		PadWidth: 5,          // lint:allow_raw_number
		Expected: "00255",
	},
	"Xh32": {
		Input:    0x0020, // lint:allow_raw_number
		PadWidth: 5,      // lint:allow_raw_number
		Expected: "00032",
	},
	"Xh255": {
		Input:    0x00FF, // lint:allow_raw_number
		PadWidth: 5,      // lint:allow_raw_number
		Expected: "00255",
	},
}
