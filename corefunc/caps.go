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

package corefunc

import "github.com/chanced/caps"

/*
StrCamel converts a string to `camelCase`, removing any non-alphanumeric
characters.

----

  - s (string): The string to reformat.

  - opts (caps.Opts): Options to pass to the underlying `caps` library.
*/
func StrCamel(s string, opts ...caps.Opts) string {
	return caps.ToLowerCamel(s, opts...)
}

/*
StrConstant converts a string to `CONSTANT_CASE` (aka, `SCREAMING_SNAKE_CASE`),
removing any non-alphanumeric characters.

----

  - s (string): The string to reformat.

  - opts (caps.Opts): Options to pass to the underlying `caps` library.
*/
func StrConstant(s string, opts ...caps.Opts) string {
	return caps.ToScreamingSnake(s, opts...)
}

/*
StrKebab converts a string to `kebab-case`, removing any non-alphanumeric
characters.

----

  - s (string): The string to reformat.

  - opts (caps.Opts): Options to pass to the underlying `caps` library.
*/
func StrKebab(s string, opts ...caps.Opts) string {
	return caps.ToKebab(s, opts...)
}

/*
StrPascal converts a string to `PascalCase` (aka, `StudlyCase`,
`UpperCamelCase`), removing any non-alphanumeric characters.

----

  - s (string): The string to reformat.

  - acronymCaps (bool): If `true`, acronyms will be converted to initialisms
    (e.g., `ID` instead of `Id`).

  - opts (caps.Opts): Options to pass to the underlying `caps` library.
*/
func StrPascal(s string, acronymCaps bool, opts ...caps.Opts) string {
	if acronymCaps {
		opts = append(opts, caps.WithReplaceStyleScreaming())
	} else {
		opts = append(opts, caps.WithReplaceStyleCamel())
	}

	return caps.ToCamel(s, opts...)
}

/*
StrSnake converts a string to `snake_case`, removing any non-alphanumeric
characters.

----

  - s (string): The string to reformat.

  - opts (caps.Opts): Options to pass to the underlying `caps` library.
*/
func StrSnake(s string, opts ...caps.Opts) string {
	return caps.ToSnake(s, opts...)
}
