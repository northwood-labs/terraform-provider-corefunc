// Copyright 2023, Ryan Parman
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
	"math"
	"strings"
)

/*
TruncateLabel supports prepending a prefix to a label, while truncating them
to meet the maximum length constraints. Useful when grouping labels with a
kind of prefix. Both the prefix and the label will be truncated if necessary.

Uses a "balancing" algorithm between the prefix and the label, so that each
section is truncated as a factor of how much space it takes up in the merged
string.

----

  - maxLength (int64): The maximum allowed length of the combined label.

  - prefix (string): The prefix to prepend to the label.

  - label (string): The label itself.

----

The motivation for this function is in working with monitoring systems such
as New Relic and Datadog where there are hundreds of applications in a
monitoring "prod" account, and also hundreds of applications in a monitoring
"nonprod" account. This allows us to group lists of monitors together using a
shared prefix, but also truncate them appropriately to fit length
constraints.
*/
func TruncateLabel(maxLength int64, prefix, label string) string {
	// None.
	if maxLength == 0 {
		return ""
	}

	// Too short.
	if maxLength <= 3 { // lint:allow_raw_number
		return "…"
	}

	// As short as possible.
	if maxLength >= 4 && maxLength <= 6 {
		return "…: …"
	}

	origName := prefix + ": " + label
	strLen := int64(len(origName))

	if strLen > maxLength {
		overage := float64(strLen - maxLength)

		prefixLengthCount := int64(math.Floor(overage/2)) + 1 // lint:allow_raw_number
		labelLengthCount := int64(math.Ceil(overage/2)) + 1   // lint:allow_raw_number

		removeFromPrefixLength := int64(len(prefix)) - prefixLengthCount
		removeFromLabelLength := int64(len(label)) - labelLengthCount

		// Rebalance the character subtractions when one side reaches 0.
		if removeFromPrefixLength < 0 {
			removeFromLabelLength += removeFromPrefixLength
			removeFromPrefixLength = 0
		}

		if removeFromLabelLength < 0 {
			removeFromPrefixLength += removeFromLabelLength
			removeFromLabelLength = 0
		}

		trimmedPrefix := prefix[:removeFromPrefixLength]
		trimmedLabel := label[:removeFromLabelLength]

		return strings.TrimSpace(trimmedPrefix) + "…: " + strings.TrimSpace(trimmedLabel) + "…"
	}

	return origName
}
