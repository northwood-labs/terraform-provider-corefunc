package corefunc

import (
	"fmt"
	"runtime"
	"strings"
)

func traceFuncName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
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
