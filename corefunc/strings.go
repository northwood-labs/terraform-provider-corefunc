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

package corefunc

import "strings"

/*
StrStartsWith checks if a string starts with a given prefix. It is functionally
identical to strings.HasPrefix() from the standard library.

----

  - str (string): The string to check.

  - prefix (string): The prefix to compare against the string.

  - Returns: (bool) True if `str` starts with `prefix`, otherwise false.
*/
func StrStartsWith(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

/*
StrEndsWith checks if a string ends with a given suffix. It is functionally
identical to strings.HasSuffix() from the standard library.

----

  - str (string): The string to check.

  - suffix (string): The suffix to compare against the string.

  - Returns: (bool) True if `str` ends with `suffix`, otherwise false.
*/
func StrEndsWith(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

/*
StrContains checks if a string contains a given substring. It is functionally
identical to strings.Contains() from the standard library.

----

  - str (string): The string to check.

  - substr (string): The substring to search for within the string.

  - Returns: (bool) True if `str` contains `substr`, otherwise false.
*/
func StrContains(str, substr string) bool {
	return strings.Contains(str, substr)
}
