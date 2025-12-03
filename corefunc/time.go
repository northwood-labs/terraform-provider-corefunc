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
	"fmt"
	"net/http"
	"time"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/corefunc/types"
)

/*
TimeParse parses an timestamp string and returns the Unix timestamp.

  - 2006-01-02 (no TZ)
  - 2006-01-02T15:04:05 (no TZ)
  - 2006-01-02T15:04:05-07:00
  - 2006-01-02T15:04:05-0700
  - 2006-01-02T15:04:05.9999999 (no TZ)
  - 2006-01-02T15:04:05.9999999-07:00
  - 2006-01-02T15:04:05.9999999-0700
  - 2006-01-02T15:04:05.9999999+07:00
  - 2006-01-02T15:04:05.9999999+0700
  - 2006-01-02T15:04:05.9999999Z
  - 2006-01-02T15:04:05+07:00
  - 2006-01-02T15:04:05+0700
  - 2006-01-02T15:04:05Z
  - Mon, 02 Jan 06 15:04:05 MST
  - Mon, 02 Jan 06 15:04:05.9999999 MST
  - Mon, 02 Jan 2006 15:04:05 GMT
  - Mon, 02 Jan 2006 15:04:05 MST
  - Mon, 02 Jan 2006 15:04:05.9999999 MST
  - Monday, 02-Jan-06 15:04:05 MST
  - Monday, 02-Jan-2006 15:04:05 MST

Specifications:

  - https://datatracker.ietf.org/doc/rfc822/
  - https://datatracker.ietf.org/doc/rfc850/
  - https://datatracker.ietf.org/doc/rfc1036/
  - https://datatracker.ietf.org/doc/rfc1123/
  - https://datatracker.ietf.org/doc/rfc2822/
  - https://datatracker.ietf.org/doc/rfc3339/
  - https://datatracker.ietf.org/doc/rfc7231/
  - https://en.wikipedia.org/wiki/ISO_8601

----

  - t (string): The RFC3339/ISO8601 timestamp string to parse.

  - Returns (int64, error): The Unix timestamp and an error if parsing fails.
*/
func TimeParse(t string) (int64, error) {
	layouts := []string{
		time.RFC3339,
		time.RFC3339Nano,
		string(types.ISO8601),
		string(types.ISO8601Nano),
		"2006-01-02T15:04:05",         // Basic date-time without timezone
		"2006-01-02T15:04:05.9999999", // Date-time with fractional seconds
		"2006-01-02",                  // Date only
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		http.TimeFormat, // RFC7231
		string(types.COOKIE),

		// Short day, spaced, 2-digit year
		"Mon, _2 Jan 06 15:04:05 MST",
		"Mon, _2 Jan 06 15:04:05.9999999 MST",

		// Short day, spaced, 4-digit year
		"Mon, _2 Jan 2006 15:04:05 MST",
		"Mon, _2 Jan 2006 15:04:05.9999999 MST",

		// Short day, hyphenated, 2-digit year
		"Mon, _2-Jan-06 15:04:05 MST",
		"Mon, _2-Jan-06 15:04:05.9999999 MST",

		// Short day, hyphenated, 4-digit year
		"Mon, _2-Jan-2006 15:04:05 MST",
		"Mon, _2-Jan-2006 15:04:05.9999999 MST",

		// Long day, spaced, 2-digit year
		"Monday, _2 Jan 06 15:04:05 MST",
		"Monday, _2 Jan 06 15:04:05.9999999 MST",

		// Long day, spaced, 4-digit year
		"Monday, _2 Jan 2006 15:04:05 MST",
		"Monday, _2 Jan 2006 15:04:05.9999999 MST",

		// Long day, hyphenated, 2-digit year
		"Monday, _2-Jan-06 15:04:05 MST",
		"Monday, _2-Jan-06 15:04:05.9999999 MST",

		// Long day, hyphenated, 4-digit year
		"Monday, _2-Jan-2006 15:04:05 MST",
		"Monday, _2-Jan-2006 15:04:05.9999999 MST",
	}

	for _, layout := range layouts {
		t1, err := time.Parse(layout, t)
		if err == nil {
			return t1.Unix(), nil
		}
	}

	return 0, fmt.Errorf("could not parse %q as a timestamp", t) // lint:allow_errorf
}

/*
TimeCompare compares two time values, which can be either Unix timestamps
(int64) or timestamp strings. The strings will be parsed into Unix timestamps
before comparison.

----

  - t1 (int64|string): The first time value to compare.

  - t2 (int64|string): The second time value to compare.

  - Returns (int, error): An integer indicating the comparison result, or an
    error if parsing fails. A value of `-1` indicates that `t1` is less than
    `t2`, `0` indicates that they are equal, and `1` indicates that `t1` is
    greater than `t2`.
*/
func TimeCompare(t1, t2 any) (int, error) {
	var (
		v1, v2 int64
		err    error
	)

	switch v := t1.(type) {
	case int64:
		v1 = v
	case int32:
		v1 = int64(v)
	case int16:
		v1 = int64(v)
	case int8:
		v1 = int64(v)
	case int:
		v1 = int64(v)
	case uint32:
		v1 = int64(v)
	case uint16:
		v1 = int64(v)
	case uint8:
		v1 = int64(v)
	case float64:
		v1 = int64(v)
	case float32:
		v1 = int64(v)
	case string:
		v1, err = TimeParse(v)
		if err != nil {
			return 0, fmt.Errorf("failed to parse first time value: %w", err)
		}
	default:
		return 0, fmt.Errorf("unsupported type for t1: %T", t1) // lint:allow_errorf
	}

	switch v := t2.(type) {
	case int64:
		v2 = v
	case int32:
		v2 = int64(v)
	case int16:
		v2 = int64(v)
	case int8:
		v2 = int64(v)
	case int:
		v2 = int64(v)
	case uint32:
		v2 = int64(v)
	case uint16:
		v2 = int64(v)
	case uint8:
		v2 = int64(v)
	case float64:
		v2 = int64(v)
	case float32:
		v2 = int64(v)
	case string:
		v2, err = TimeParse(v)
		if err != nil {
			return 0, fmt.Errorf("failed to parse second time value: %w", err)
		}
	default:
		return 0, fmt.Errorf("unsupported type for t2: %T", t2) // lint:allow_errorf
	}

	// Compare the two Unix timestamps
	if v1 < v2 {
		return -1, nil
	} else if v1 > v2 {
		return 1, nil
	}

	return 0, nil
}
