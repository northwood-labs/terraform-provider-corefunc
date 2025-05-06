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

import "github.com/mattn/go-runewidth"

/*
StrByteLength returns the number of bytes/runes in a string. This is different
from the number of characters, which is what Terraform and OpenTofu's `length()`
function returns for strings.

----

  - str (string): The string with which to count bytes/runes.
*/
func StrByteLength(str string) int {
	return runewidth.StringWidth(str)
}
