// Copyright 2024-2025, Northwood Labs, LLC <license@northwood-labs.com>
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

/*
See "corefunc" package for the Go library code.
*/
package main

import (
	"github.com/northwood-labs/terraform-provider-corefunc/v2/cmd"
)

// Provider documentation generation.
//go:generate go tool -modfile=go.tools.mod tfplugindocs generate --provider-name corefunc --rendered-provider-name "Core Functions" # lint:ignore_length

func main() {
	cmd.Execute()
}
