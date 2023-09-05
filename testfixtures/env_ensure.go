package testfixtures

import (
	"errors"
	"regexp"
)

// EnvEnsureTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var EnvEnsureTestTable = map[string]struct {
	EnvVarName  string
	SetValue    string
	ExpectedErr error
	Pattern     *regexp.Regexp
}{
	"AWS_DEFAULT_REGION": {
		EnvVarName: "AWS_DEFAULT_REGION",
		SetValue:   "us-east-1",
	},
	"AWS_REGION": {
		EnvVarName: "AWS_REGION",
		SetValue:   "us-east-1",
		Pattern:    regexp.MustCompile(`^([a-z]{2})-([^-]+)-(\\d)$`),
	},
	"AWS_VAULT": {
		EnvVarName:  "AWS_VAULT",
		SetValue:    "dev",
		Pattern:     regexp.MustCompile(`(non)?prod$`),
		ExpectedErr: errors.New("environment variable AWS_VAULT does not match pattern (non)?prod$"), // lint:allow_errorf
	},
	"AWS_PAGER": {
		EnvVarName:  "AWS_PAGER",
		ExpectedErr: errors.New("environment variable AWS_PAGER is not defined"), // lint:allow_errorf
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
