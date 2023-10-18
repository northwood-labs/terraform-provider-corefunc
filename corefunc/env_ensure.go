// Copyright 2023, Ryan Parman
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
	"errors"
	"os"
	"regexp"
)

/*
EnvEnsure ensures that a given environment variable is set to a non-empty value.
If the environment variable is unset or if it is set to an empty string,
EnvEnsure will respond with an error.

Not every Terraform provider checks to ensure that the environment variables it
requires are properly set before performing work, leading to late-stage errors.
This will force an error to occur early in the execution if the environment
variable is not set, or if its value doesn't match the expected patttern.

----

  - name (string): The name of the environment variable to check.

  - pattern (*regexp.Regexp): The result of a call to `regexp.Compile()` or
    `regexp.MustCompile()`.
*/
func EnvEnsure(name string, pattern ...*regexp.Regexp) error {
	if os.Getenv(name) == "" {
		return errors.New("environment variable " + name + " is not defined") // lint:allow_errorf
	}

	if len(pattern) > 0 && pattern[0] != nil {
		if !pattern[0].MatchString(os.Getenv(name)) {
			return errors.New( // lint:allow_errorf
				"environment variable " + name + " does not match pattern " + pattern[0].String(),
			)
		}

		return nil
	}

	return nil
}
