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

import "strings"

/*
StrLeftPad pads a string on the left side with a given string until it reaches
the desired width.

----

  - str (string): The string to pad.

  - padWidth (int): The total width of the padded string. If `str` is equivalent
    length or longer, then no padding will be applied.

  - padChar (string): The string to use as padding. If unspecified, will use a
    single space character.
*/
func StrLeftPad(str string, padWidth int, padChar ...byte) string {
	if len(str) >= padWidth {
		return str
	}

	var pStr byte
	pStr = 0x20

	if padWidth < 0 {
		padWidth = 0
	}

	if len(padChar) > 0 {
		pStr = padChar[0]
	}

	return strings.Repeat(string(pStr), padWidth-len(str)) + str
}
