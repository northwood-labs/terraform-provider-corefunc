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

package corefuncprovider

import (
	"context"
	"runtime"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

	out := strings.TrimPrefix(
		frame.Function,
		"github.com/northwood-labs/terraform-provider-corefunc/corefunc.",
	)

	out = strings.TrimPrefix(
		out,
		"github.com/northwood-labs/terraform-provider-corefunc/corefuncprovider.",
	)

	return out + "\n"
}

func linkPackage(functionName string) string {
	return "[`corefunc." + functionName + "()`]" +
		"(https://pkg.go.dev/github.com/northwood-labs/terraform-provider-corefunc/" +
		"corefunc#" + functionName + ")"
}

func code(s string) string {
	return "`" + s + "`"
}

func convertTFMapToGoMap(ctx context.Context, m types.Map) (map[string]attr.Value, diag.Diagnostics) {
    goObject := make(map[string]attr.Value)

	// Read from the Terraform data into the map
    if diags := m. ElementsAs(ctx, &goObject, false); diags != nil {
        return nil, diags
    }

	return goObject, nil
}

func convertGoMapToTFMap(ctx context.Context, gm map[string]any) (types.Map, diag.Diagnostics) {
	mappedValue, diags := types.MapValueFrom(ctx, types.StringType, gm)

	if diags != nil {
        return types.MapNull(nil), diags
    }

	return mappedValue, nil
}
