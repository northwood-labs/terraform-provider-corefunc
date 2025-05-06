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

package testfixtures

import (
	"errors"
	"regexp"
)

// EnvEnsureTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var EnvEnsureTestTable = map[string]struct {
	Pattern       *regexp.Regexp
	ExpectedErrRE *regexp.Regexp
	ExpectedErr   error
	EnvVarName    string
	SetValue      string
	PatternStr    string
}{
	"AWS_DEFAULT_REGION": {
		EnvVarName: "AWS_DEFAULT_REGION",
		SetValue:   "us-east-1",
	},
	"AWS_REGION": {
		EnvVarName: "AWS_REGION",
		SetValue:   "us-east-1",
		// Pattern:    regexp.MustCompile(`^([a-z]{2})-([^-]+)-(\\d)$`),
		PatternStr: `^([a-z]{2})-([^-]+)-(\\d)$`,
	},
	"AWS_VAULT": {
		EnvVarName: "AWS_VAULT",
		SetValue:   "dev",
		Pattern:    regexp.MustCompile(`(non)?prod$`),
		PatternStr: `(non)?prod$`,
		ExpectedErr: errors.New(
			"environment variable AWS_VAULT does not match pattern (non)?prod$",
		), // lint:allow_errorf
		ExpectedErrRE: regexp.MustCompile(
			`environment variable AWS_VAULT does not match pattern \(non\)\?prod\$`,
		), // lint:allow_errorf
	},
	"AWS_PAGER": {
		EnvVarName:    "AWS_PAGER",
		ExpectedErr:   errors.New("environment variable AWS_PAGER is not defined"), // lint:allow_errorf
		ExpectedErrRE: regexp.MustCompile(`environment variable AWS_PAGER is not defined`),
	},
	"BREW_PREFIX": {
		EnvVarName: "BREW_PREFIX",
		SetValue:   "/opt/homebrew",
	},
	"GO111MODULE": {
		EnvVarName: "GO111MODULE",
		SetValue:   "on",
	},
	"LANG": {
		EnvVarName: "LANG",
		SetValue:   "en_US.UTF-8",
	},
	"LC_ALL": {
		EnvVarName: "LC_ALL",
		SetValue:   "en_US.UTF-8",
	},
	"LSCOLORS": {
		EnvVarName: "LSCOLORS",
		SetValue:   "gxBxhxDxfxhxhxhxhxcxcx",
	},
	"SHELL": {
		EnvVarName: "SHELL",
		SetValue:   "/opt/homebrew/bin/bash",
	},
	"TERM": {
		EnvVarName: "TERM",
		SetValue:   "xterm-256color",
	},
}
