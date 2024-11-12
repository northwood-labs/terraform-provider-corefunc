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

package corefunc

import (
	"fmt"
	"io"
	"os"
)

/*
File reads the contents of a file and returns it as a string. It is meant to
simulate the behavior of the `file()` function in Terraform/OpenTofu.

----

  - path (string): A file name path to read.
*/
func File(path string) (string, error) {
	r, err := os.OpenFile(path, os.O_RDONLY, 0o644) // lint:allow_possible_insecure lint:allow_raw_number
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("failed to read the file: %w", err)
	}

	if err := r.Close(); err != nil {
		return "", fmt.Errorf("failed to close file: %w", err)
	}

	return string(b), nil
}
