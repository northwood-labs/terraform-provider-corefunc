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

// StrByteLengthTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var StrByteLengthTestTable = map[string]struct { // lint:no_dupe
	Input    string
	Expected int
}{
	"blank":    {"", 0},
	"abcde":    {"abcde", 5},
	"世":        {"世", 2},
	"界":        {"界", 2},
	"ｾ":        {"ｾ", 1},
	"ｶ":        {"ｶ", 1},
	"ｲ":        {"ｲ", 1},
	"☆":        {"☆", 1},
	"☺":        {"☺", 1},
	"☻":        {"☻", 1},
	"♥":        {"♥", 1},
	"♦":        {"♦", 1},
	"♣":        {"♣", 1},
	"♠":        {"♠", 1},
	"♂":        {"♂", 1},
	"♀":        {"♀", 1},
	"♪":        {"♪", 1},
	"♫":        {"♫", 1},
	"☼":        {"☼", 1},
	"↕":        {"↕", 1},
	"‼":        {"‼", 1},
	"↔":        {"↔", 1},
	"\x00":     {"\x00", 0},
	"\x01":     {"\x01", 0},
	"\u0300":   {"\u0300", 0},
	"\u2028":   {"\u2028", 0},
	"\u2029":   {"\u2029", 0},
	"a":        {"a", 1},
	"⟦":        {"⟦", 1},
	"👁":        {"👁", 1},
	"■㈱の世界①":   {"■㈱の世界①", 10},
	"スター☆":     {"スター☆", 7},
	"つのだ☆HIRO": {"つのだ☆HIRO", 11},
}
