package corefunc

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	// Terratest displays a Markdown link to Terratest.
	Terratest = "[Terratest](https://terratest.gruntwork.io)"

	// TPF displays a Markdown link to the Terraform Plugin Framework.
	TPF = "[Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework)"
)

func traceFuncName() string { // lint:allow_unused
	pc := make([]uintptr, 15)   // lint:allow_raw_number
	n := runtime.Callers(2, pc) // lint:allow_raw_number
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return fmt.Sprintf(
		"%s\n",
		strings.TrimPrefix(
			frame.Function,
			"github.com/northwood-labs/terraform-provider-corefunc/corefunc.",
		),
	)
}

func linkPackage(functionName string) string {
	return "[`corefunc." +
		functionName +
		"()`](https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/corefunc#" +
		functionName +
		")"
}
