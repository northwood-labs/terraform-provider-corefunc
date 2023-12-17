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

package testfixtures // lint:no_dupe

import (
	"log"
	"os/user"
	"path/filepath"
)

var (
	u = func() *user.User {
		us, err := user.Current()
		if err != nil {
			log.Fatalf("err: %s", err)
		}

		return us
	}()

	// HomedirExpandTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	HomedirExpandTestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
		Err      bool
	}{
		"root-foo": {
			Input:    "/foo",
			Expected: "/foo",
			Err:      false,
		},
		"home-foo": {
			"~/foo",
			filepath.Join(u.HomeDir, "foo"),
			false,
		},
		"empty": {
			"",
			"",
			false,
		},
		"home": {
			"~",
			u.HomeDir,
			false,
		},
		"invalid": {
			"~foo/foo",
			"",
			true,
		},
	}
)
