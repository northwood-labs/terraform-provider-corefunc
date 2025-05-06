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

import "github.com/mitchellh/go-homedir"

/*
Homedir returns the home directory for the executing user without requiring CGO
for macOS.
*/
func Homedir() (string, error) {
	return homedir.Dir() // lint:allow_unwrapped_errors
}

/*
HomedirExpand expands the path to include the home directory if the path is prefixed
with `~`. If it isn't prefixed with `~`, the path is returned as-is.
*/
func HomedirExpand(path string) (string, error) {
	return homedir.Expand(path) // lint:allow_unwrapped_errors
}
