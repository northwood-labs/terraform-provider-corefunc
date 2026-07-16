// Copyright 2024-2026, Northwood Labs, LLC <license@northwood-labs.com>
// Copyright 2023-2025, Ryan Parman <rparman@northwood-labs.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package corefunc

import (
	"strings"
	"testing"

	"pgregory.net/rapid"
)

// drawByteString generates a random string of arbitrary bytes with length between min and max.
func drawByteString(t *rapid.T, imin, imax int, label string) string {
	b := rapid.SliceOfN(rapid.Byte(), imin, imax).Draw(t, label)

	return string(b)
}

// Property 1: Output Length Invariant
// **Validates: Requirements 1.1, 1.2, 1.3, 1.4**.
func TestStrLeftPad_OutputLength(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		str := drawByteString(t, 0, 1024, "str")
		padWidth := rapid.IntRange(0, 2048).Draw(t, "padWidth")
		// Restrict to ASCII range (0x01–0x7F) because string(byte) for values >= 0x80
		// produces a multi-byte UTF-8 encoding, making byte-length assertions invalid.
		padChar := rapid.ByteRange(0x01, 0x7F).Draw(t, "padChar")

		result := StrLeftPad(str, padWidth, padChar)

		// When len(str) < padWidth, the output length must equal padWidth.
		if len(str) < padWidth {
			if len(result) != padWidth {
				t.Fatalf("len(str) < padWidth: expected len(result) == %d, got %d", padWidth, len(result))
			}
		}

		// When len(str) >= padWidth, the output length must equal len(str).
		if len(str) >= padWidth {
			if len(result) != len(str) {
				t.Fatalf("len(str) >= padWidth: expected len(result) == %d, got %d", len(str), len(result))
			}
		}

		// When padWidth == 0, the output length must equal len(str).
		if padWidth == 0 {
			if len(result) != len(str) {
				t.Fatalf("padWidth == 0: expected len(result) == %d, got %d", len(str), len(result))
			}
		}
	})
}

// Property 2: Content Preservation (Suffix)
// **Validates: Requirements 2.1, 2.2, 2.3**.
func TestStrLeftPad_ContentPreservation(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		str := drawByteString(t, 0, 1000, "str")
		padWidth := rapid.IntRange(0, 1000).Draw(t, "padWidth")
		padChar := rapid.ByteRange(0x01, 0xFF).Draw(t, "padChar")

		result := StrLeftPad(str, padWidth, padChar)

		// The original string must always appear as a suffix of the result.
		if !strings.HasSuffix(result, str) {
			t.Fatalf("result %q does not have suffix %q", result, str)
		}

		// When padWidth exceeds the input string length, the padding prefix
		// must consist entirely of repeated padChar characters.
		if padWidth > len(str) {
			padCount := padWidth - len(str)
			expectedPrefix := strings.Repeat(string(rune(padChar)), padCount)

			actualPrefix := result[:len(result)-len(str)]
			if actualPrefix != expectedPrefix {
				t.Fatalf(
					"padding prefix mismatch for padChar=0x%02X, padCount=%d: got %q, want %q",
					padChar, padCount, actualPrefix, expectedPrefix,
				)
			}
		}
	})
}

// Property 3: Pad Character Correctness
// **Validates: Requirements 5.1, 5.2, 5.3, 5.4**.
func TestStrLeftPad_PadCharCorrectness(t *testing.T) {
	t.Run("explicit_padchar", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			str := drawByteString(t, 0, 1024, "str")
			padChar := rapid.Byte().Draw(t, "padChar")

			// Ensure padWidth exceeds string length so padding is applied.
			extra := rapid.IntRange(1, 1024).Draw(t, "extra")
			padWidth := len(str) + extra

			result := StrLeftPad(str, padWidth, padChar)

			// The padding prefix must consist of exactly (padWidth - len(str))
			// repetitions of string(rune(padChar)). The function converts the byte
			// to a rune and then to its UTF-8 string representation for padding.
			padCount := padWidth - len(str)
			expectedPrefix := strings.Repeat(string(rune(padChar)), padCount)
			actualPrefix := result[:len(result)-len(str)]

			if actualPrefix != expectedPrefix {
				t.Fatalf(
					"padding prefix mismatch for padChar=%#x: got %q, want %q",
					padChar, actualPrefix, expectedPrefix,
				)
			}
		})
	})

	t.Run("default_padchar_is_space", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			str := drawByteString(t, 0, 1024, "str")

			// Ensure padWidth exceeds string length so padding is applied.
			extra := rapid.IntRange(1, 1024).Draw(t, "extra")
			padWidth := len(str) + extra

			// Call without specifying padChar to test the default.
			result := StrLeftPad(str, padWidth)

			// Every byte in the padding prefix must equal space (0x20).
			// Space is in the ASCII range so string(byte(0x20)) is a single byte.
			padLen := padWidth - len(str)

			prefix := result[:padLen]
			for i := range len(prefix) {
				if prefix[i] != 0x20 {
					t.Fatalf("byte at index %d is %#x, want 0x20 (space)", i, prefix[i])
				}
			}
		})
	})

	t.Run("no_padding_when_width_not_exceeded", func(t *testing.T) {
		rapid.Check(t, func(t *rapid.T) {
			str := drawByteString(t, 1, 1024, "str")
			padChar := rapid.Byte().Draw(t, "padChar")

			// padWidth <= len(str): output must equal input unchanged.
			padWidth := rapid.IntRange(0, len(str)).Draw(t, "padWidth")

			result := StrLeftPad(str, padWidth, padChar)

			if result != str {
				t.Fatalf(
					"expected output to equal input when padWidth(%d) <= len(str)(%d), got %q",
					padWidth,
					len(str),
					result,
				)
			}
		})
	})
}

// Property 4: Idempotence
// **Validates: Requirements 3.1, 3.2**.
func TestStrLeftPad_Idempotence(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		str := drawByteString(t, 0, 1000, "str")
		padWidth := rapid.IntRange(0, 1000).Draw(t, "padWidth")
		padChar := rapid.ByteRange(0x20, 0x7E).Draw(t, "padChar")

		// Apply StrLeftPad once.
		once := StrLeftPad(str, padWidth, padChar)

		// Apply StrLeftPad twice with the same parameters.
		twice := StrLeftPad(once, padWidth, padChar)

		// Idempotence: applying twice equals applying once.
		if twice != once {
			t.Fatalf(
				"idempotence violated: StrLeftPad(StrLeftPad(%q, %d, %q), %d, %q) = %q, want %q",
				str, padWidth, string(padChar), padWidth, string(padChar), twice, once,
			)
		}

		// When len(str) >= padWidth, result equals input unchanged.
		if len(str) >= padWidth {
			if once != str {
				t.Fatalf(
					"identity violated: len(%q)=%d >= padWidth=%d, but got %q, want %q",
					str, len(str), padWidth, once, str,
				)
			}
		}
	})
}

// Property 5: Identity (No-Op)
// **Validates: Requirements 3.3, 4.1, 4.2, 4.3, 4.4**.
func TestStrLeftPad_Identity(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		str := drawByteString(t, 0, 200, "str")
		negPadWidth := rapid.IntRange(-1000, 0).Draw(t, "negPadWidth")
		padChar := rapid.ByteRange(0x01, 0x7F).Draw(t, "padChar")

		// When padWidth <= 0, result must equal input unchanged.
		result := StrLeftPad(str, negPadWidth, padChar)
		if result != str {
			t.Fatalf(
				"identity violated: padWidth=%d, expected %q, got %q",
				negPadWidth, str, result,
			)
		}

		// When len(str) >= padWidth (positive), result must equal input unchanged.
		if str != "" {
			smallWidth := rapid.IntRange(0, len(str)).Draw(t, "smallWidth")

			result2 := StrLeftPad(str, smallWidth, padChar)
			if result2 != str {
				t.Fatalf(
					"identity violated: len(str)=%d >= padWidth=%d, expected %q, got %q",
					len(str), smallWidth, str, result2,
				)
			}
		}

		// Varying padChar does not influence output when padWidth is non-positive.
		otherChar := rapid.ByteRange(0x01, 0x7F).Draw(t, "otherChar")

		result3 := StrLeftPad(str, negPadWidth, otherChar)
		if result3 != str {
			t.Fatalf(
				"padChar should not influence output for non-positive padWidth=%d: got %q, want %q",
				negPadWidth, result3, str,
			)
		}
	})
}

// Property 6: Monotonic Growth (Metamorphic)
// **Validates: Requirements 6.1, 6.2, 6.3, 6.5**.
func TestStrLeftPad_MonotonicGrowth(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		str := drawByteString(t, 0, 200, "str")
		// Restrict padChar to ASCII range so string(byte) produces a single byte,
		// making the metamorphic assertion straightforward.
		padChar := rapid.ByteRange(0x01, 0x7F).Draw(t, "padChar")

		// Draw a padWidth strictly greater than len(str) for growth assertions.
		minWidth := min(len(str)+1, 300)

		padWidth := rapid.IntRange(minWidth, 300).Draw(t, "padWidth")

		result := StrLeftPad(str, padWidth, padChar)

		// 6.1: Output is never shorter than the input.
		if len(result) < len(str) {
			t.Fatalf(
				"output shorter than input: len(result)=%d < len(str)=%d",
				len(result), len(str),
			)
		}

		// 6.2: Padding with padWidth == len(str) returns input unchanged.
		identityResult := StrLeftPad(str, len(str), padChar)
		if identityResult != str {
			t.Fatalf(
				"padding with padWidth==len(str) should return input unchanged: got %q, want %q",
				identityResult, str,
			)
		}

		// 6.3: Metamorphic — increasing padWidth by one prepends exactly one pad character.
		// StrLeftPad(str, w+1, c) == string(c) + StrLeftPad(str, w, c) when w >= len(str).
		w := padWidth - 1 // w >= len(str) since padWidth > len(str).
		left := StrLeftPad(str, w+1, padChar)

		right := string(rune(padChar)) + StrLeftPad(str, w, padChar)
		if left != right {
			t.Fatalf(
				"metamorphic property violated: StrLeftPad(str, %d, %#x) = %q, but string(c)+StrLeftPad(str, %d, %#x) = %q",
				w+1,
				padChar,
				left,
				w,
				padChar,
				right,
			)
		}

		// 6.5: Output does NOT equal input when padWidth > len(str) and len(str) > 0.
		if str != "" && result == str {
			t.Fatalf(
				"expected output != input when padWidth(%d) > len(str)(%d), but got equal: %q",
				padWidth, len(str), result,
			)
		}
	})
}
