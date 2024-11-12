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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

/*
TOMLtoJSON converts a TOML string to a JSON string. The result can be passed to
`json.Unmarshal()` to convert it to a `map[string]any`.

Ported from `pelletier/go-toml`.

<https://pkg.go.dev/github.com/pelletier/go-toml/v2>

----

  - tomlStr (string): A string of TOML content.
*/
func TOMLtoJSON(tomlStr string) (string, error) {
	var (
		v any
		w bytes.Buffer
	)

	tomlDecoder := toml.NewDecoder(strings.NewReader(tomlStr))

	err := tomlDecoder.Decode(&v)
	if err != nil {
		var derr *toml.DecodeError
		if errors.As(err, &derr) {
			row, col := derr.Position()

			return "", fmt.Errorf("error occurred at row %d column %d: %s", row, col, derr.String())
		}

		return "", fmt.Errorf("failed to decode TOML: %w", err)
	}

	jsonEncoder := json.NewEncoder(&w)

	err = jsonEncoder.Encode(v)
	if err != nil {
		return "", fmt.Errorf("failed to encode to JSON: %w", err)
	}

	return w.String(), nil
}

/*
JSONtoTOML converts a JSON string to a TOML string. The result can be passed to
`toml.Unmarshal()` to convert it to a `map[string]any`.

Ported from `pelletier/go-toml`.

<https://pkg.go.dev/github.com/pelletier/go-toml/v2>

----

  - jsonStr (string): A string of JSON content.
*/
func JSONtoTOML(jsonStr string) (string, error) {
	var (
		v any
		w bytes.Buffer
	)

	jsonDecoder := json.NewDecoder(strings.NewReader(jsonStr))
	tomlEncoder := toml.NewEncoder(&w)

	jsonDecoder.UseNumber()
	tomlEncoder.SetMarshalJsonNumbers(true)

	err := jsonDecoder.Decode(&v)
	if err != nil {
		return "", fmt.Errorf("failed to decode to JSON: %w", err)
	}

	if err = tomlEncoder.Encode(v); err != nil {
		return "", fmt.Errorf("failed to encode to TOML: %w", err)
	}

	return w.String() + "\n", nil
}
