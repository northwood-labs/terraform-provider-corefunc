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

package corefunc

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/pelletier/go-toml/v2"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

func ExampleTOMLtoJSON() {
	tomlStr := `title = 'TOML Example'

[database]
data = [['delta', 'phi'], [3.14]]
enabled = true
ports = [8000, 8001, 8002]

[database.temp_targets]
case = 72
cpu = 79.5

[owner]
dob = '1979-05-27T07:32:00-08:00'
name = 'Tom Preston-Werner'

[servers]
    [servers.alpha]
    ip = '10.0.0.1'
    role = 'frontend'

    [servers.beta]
    ip = '10.0.0.2'
    role = 'backend'
`

	jsonStr, err := TOMLtoJSON(tomlStr)
	if err != nil {
		panic(err)
	}

	var v any

	err = json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		panic(err)
	}

	if m, ok := v.(map[string]any); ok {
		if title, exists := m["title"]; exists {
			fmt.Println(title)
		}

		if servers, ok := m["servers"].(map[string]any); ok {
			if alpha, ok := servers["alpha"].(map[string]any); ok {
				if role, exists := alpha["role"]; exists {
					fmt.Println(role)
				}
			}
		}
	}

	// Output:
	// TOML Example
	// frontend
}

func ExampleJSONtoTOML() {
	jsonStr := `{
  "title": "TOML Example",
  "owner": {
    "dob": "1979-05-27T07:32:00-08:00",
    "name": "Tom Preston-Werner"
  },
  "database": {
    "data": [
      ["delta", "phi"],
      [3.14]
    ],
    "enabled": true,
    "ports": [8000, 8001, 8002],
    "temp_targets": {
      "case": 72,
      "cpu": 79.5
    }
  },
  "servers": {
    "alpha": {
      "ip": "10.0.0.1",
      "role": "frontend"
    },
    "beta": {
      "ip": "10.0.0.2",
      "role": "backend"
    }
  }
}`

	tomlStr, err := JSONtoTOML(jsonStr)
	if err != nil {
		panic(err)
	}

	var v any

	err = toml.Unmarshal([]byte(tomlStr), &v)
	if err != nil {
		panic(err)
	}

	if m, ok := v.(map[string]any); ok {
		if title, exists := m["title"]; exists {
			fmt.Println(title)
		}

		if servers, ok := m["servers"].(map[string]any); ok {
			if alpha, ok := servers["alpha"].(map[string]any); ok {
				if role, exists := alpha["role"]; exists {
					fmt.Println(role)
				}
			}
		}
	}

	// Output:
	// TOML Example
	// frontend
}

func TestTOMLtoJSON(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.TOMLtoJSONTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := TOMLtoJSON(tc.Input)

			// We expect an error.
			if err != nil && !tc.ExpectedErr {
				t.Errorf("Unexpected error: %v", err)
			}

			if strings.TrimSpace(output) != strings.TrimSpace(tc.Expected) {
				t.Errorf(
					"Expected ----\n%s\n---- got ----\n%s\n----",
					strings.TrimSpace(tc.Expected),
					strings.TrimSpace(output),
				)
			}
		})
	}
}

func TestJSONtoTOML(t *testing.T) { // lint:allow_complexity
	for name, tc := range testfixtures.JSONtoTOMLTestTable {
		t.Run(name, func(t *testing.T) {
			output, err := JSONtoTOML(tc.Input)

			// We expect an error.
			if err != nil && !tc.ExpectedErr {
				t.Errorf("Unexpected error: %v", err)
			}

			if strings.TrimSpace(output) != strings.TrimSpace(tc.Expected) {
				t.Errorf(
					"Expected ----\n%s\n---- got ----\n%s\n----",
					strings.TrimSpace(tc.Expected),
					strings.TrimSpace(output),
				)
			}
		})
	}
}
