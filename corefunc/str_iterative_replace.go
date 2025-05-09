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
	"strings"

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

/*
StrIterativeReplace iterates over a list of replacements. This allows you to
accept a list of replacements of unknown length from users, and apply them all
in sequence. It is a wrapper around `strings.ReplaceAll()`.

----

  - str (string): The string to which we apply the replacements.

  - replacements ([]types.Replacement): A list of replacements to apply to
    the string, in sequence.
*/
func StrIterativeReplace(str string, replacements []types.Replacement) string {
	s := str

	for i := range replacements {
		replacement := replacements[i]
		s = strings.ReplaceAll(s, replacement.Old, replacement.New)
	}

	return s
}
