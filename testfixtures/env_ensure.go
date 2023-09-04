package testfixtures

import "errors"

// EnvEnsureTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var EnvEnsureTestTable = map[string]struct {
	EnvVarName  string
	SetValue    string
	ExpectedErr error
}{
	"AWS_DEFAULT_REGION": {
		EnvVarName:  "AWS_DEFAULT_REGION",
		SetValue:    "us-east-1",
		ExpectedErr: nil,
	},
	"AWS_REGION": {
		EnvVarName:  "AWS_REGION",
		SetValue:    "us-east-1",
		ExpectedErr: nil,
	},
	"AWS_PAGER": {
		EnvVarName:  "AWS_PAGER",
		SetValue:    "",
		ExpectedErr: errors.New("environment variable AWS_PAGER is not defined"), // lint:allow_errorf
	},
	"BREW_PREFIX": {
		EnvVarName:  "BREW_PREFIX",
		SetValue:    "/opt/homebrew",
		ExpectedErr: nil,
	},
	"GO111MODULE": {
		EnvVarName:  "GO111MODULE",
		SetValue:    "on",
		ExpectedErr: nil,
	},
	"LANG": {
		EnvVarName:  "LANG",
		SetValue:    "en_US.UTF-8",
		ExpectedErr: nil,
	},
	"LC_ALL": {
		EnvVarName:  "LC_ALL",
		SetValue:    "en_US.UTF-8",
		ExpectedErr: nil,
	},
	"LSCOLORS": {
		EnvVarName:  "LSCOLORS",
		SetValue:    "gxBxhxDxfxhxhxhxhxcxcx",
		ExpectedErr: nil,
	},
	"SHELL": {
		EnvVarName:  "SHELL",
		SetValue:    "/opt/homebrew/bin/bash",
		ExpectedErr: nil,
	},
	"TERM": {
		EnvVarName:  "TERM",
		SetValue:    "xterm-256color",
		ExpectedErr: nil,
	},
}
