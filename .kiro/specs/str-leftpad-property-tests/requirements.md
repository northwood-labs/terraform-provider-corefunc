# Requirements Document

## Introduction

This feature adds property-based testing using the `pgregory.net/rapid` library to the `StrLeftPad` function in the `corefunc` package. Property-based tests complement the existing table-driven tests, benchmarks, and fuzz tests by verifying algebraic invariants and behavioral properties that hold across all valid inputs, rather than testing specific known input/output pairs.

## Glossary

* **Property_Test_Suite**: The collection of property-based tests for `StrLeftPad` using the `pgregory.net/rapid` library.
* **StrLeftPad**: The Go function that pads a string on the left side with a given character until it reaches the desired width.
* **Rapid**: The `pgregory.net/rapid` property-based testing library for Go.
* **PadWidth**: The total desired width of the resulting string after padding.
* **PadChar**: The single-byte character used for padding (defaults to space `0x20`).
* **Generator**: A Rapid construct that produces random test inputs conforming to specified constraints.

## Requirements

### Requirement 1: output length property

**User Story:** As a developer, I want a property test that verifies the output length invariant of `StrLeftPad`, so that I can confirm the function always produces correctly-sized results regardless of input.

#### Acceptance criteria

1. WHEN the input string byte-length is less than PadWidth, THE Property_Test_Suite SHALL verify that the output byte-length equals PadWidth.
2. WHEN the input string byte-length is greater than or equal to PadWidth, THE Property_Test_Suite SHALL verify that the output byte-length equals the input string byte-length.
3. THE Property_Test_Suite SHALL generate random strings with byte-length between 0 and 1024 and PadWidth values between 0 and 2048 using Rapid generators, ensuring both conditions (input shorter than PadWidth and input at least as long as PadWidth) are tested as separate property checks.
4. WHEN PadWidth is zero, THE Property_Test_Suite SHALL verify that the output byte-length equals the input string byte-length.

### Requirement 2: content preservation property

**User Story:** As a developer, I want a property test that verifies the original string is always preserved in the output, so that I can confirm padding never corrupts the input.

#### Acceptance criteria

1. THE Property_Test_Suite SHALL verify that the output always ends with the original input string (suffix preservation), including when the input is an empty string.
2. WHEN PadWidth exceeds the input string length, THE Property_Test_Suite SHALL verify that the first `PadWidth - len(input)` bytes of the output each equal the PadChar character.
3. THE Property_Test_Suite SHALL generate random strings of length 0 to 1000, pad widths of 0 to 1000, and pad characters in the byte range 0x01 to 0xFF using Rapid generators.

### Requirement 3: idempotence property

**User Story:** As a developer, I want a property test that verifies `StrLeftPad` is idempotent when the string already meets the target width, so that I can confirm repeated application does not alter already-padded strings.

#### Acceptance criteria

1. WHEN the input string byte length is greater than or equal to PadWidth, THE Property_Test_Suite SHALL verify that `StrLeftPad(str, padWidth, padChar)` returns the input string unchanged, using at least 100 generated test cases with PadWidth in the range 0 to 1000 and input strings of byte length 0 to 1000.
2. THE Property_Test_Suite SHALL verify that for any combination of str (byte length 0 to 1000), padWidth (0 to 1000), and padChar (any single byte value 0x20 to 0x7E), applying `StrLeftPad` twice with the same parameters produces the same result as applying it once (i.e., `StrLeftPad(StrLeftPad(str, w, c), w, c) == StrLeftPad(str, w, c)`), using at least 100 generated test cases.
3. WHEN PadWidth is less than or equal to zero, THE Property_Test_Suite SHALL verify that `StrLeftPad(str, padWidth, padChar)` returns the input string unchanged regardless of the padChar value.

### Requirement 4: negative and zero PadWidth property

**User Story:** As a developer, I want a property test that verifies `StrLeftPad` handles non-positive pad widths correctly, so that I can confirm edge cases produce safe, predictable results.

#### Acceptance criteria

1. WHEN PadWidth is zero, THE Property_Test_Suite SHALL verify that the output is byte-for-byte equal to the input string regardless of input string length or PadChar value.
2. WHEN PadWidth is negative, THE Property_Test_Suite SHALL verify that the output is byte-for-byte equal to the input string regardless of input string length or PadChar value.
3. THE Property_Test_Suite SHALL generate PadWidth values in the range of -1000 to 0 (inclusive), random input strings of length 0 to 200 bytes, and random single-byte pad characters using Rapid generators.
4. THE Property_Test_Suite SHALL verify that varying PadChar does not influence the output when PadWidth is non-positive.

### Requirement 5: pad character correctness property

**User Story:** As a developer, I want a property test that verifies only the specified pad character appears in the padding portion, so that I can confirm no unintended characters are introduced.

#### Acceptance criteria

1. WHEN PadWidth exceeds the input string length, THE Property_Test_Suite SHALL verify that every byte in the padding prefix (output bytes from index 0 to PadWidth minus input string length, exclusive) equals the specified PadChar.
2. WHEN PadWidth exceeds the input string length and no PadChar is specified, THE Property_Test_Suite SHALL verify that every byte in the padding prefix equals the space character (`0x20`).
3. THE Property_Test_Suite SHALL generate random PadChar values across the full single-byte range (0x00–0xFF) using Rapid generators.
4. WHEN PadWidth does not exceed the input string length, THE Property_Test_Suite SHALL verify that the output contains no padding prefix (output equals the input string unchanged).

### Requirement 6: metamorphic relationship property

**User Story:** As a developer, I want a property test that verifies known metamorphic relationships of `StrLeftPad`, so that I can confirm consistent behavior across related inputs.

#### Acceptance criteria

1. THE Property_Test_Suite SHALL verify that `len(StrLeftPad(str, w, c)) >= len(str)` for all generated inputs, where `str` is a string of byte-length 0 to 200, `w` is an integer from 0 to 300, and `c` is any single byte value, confirming the output is never shorter than the input in bytes.
2. THE Property_Test_Suite SHALL verify that padding with PadWidth equal to `len(str)` (byte-length) produces the input string unchanged, for strings of byte-length 0 to 200 and any single-byte pad character.
3. IF PadWidth exceeds `len(str)`, THEN THE Property_Test_Suite SHALL verify that increasing PadWidth by one adds exactly one pad character byte to the beginning of the output, for strings of byte-length 0 to 200 and PadWidth from `len(str)+1` to 300.
4. THE Property_Test_Suite SHALL execute a minimum of 1000 generated input combinations per property to provide sufficient coverage of the input domain.
5. WHEN PadWidth exceeds `len(str)`, THE Property_Test_Suite SHALL verify that the output does NOT equal the input string (confirming that padding actually occurs), for strings of byte-length 1 to 200 and PadWidth from `len(str)+1` to 300.

### Requirement 7: test integration

**User Story:** As a developer, I want the property-based tests to integrate with the existing test infrastructure, so that I can run them alongside existing tests using standard Go tooling.

#### Acceptance criteria

1. THE Property_Test_Suite SHALL reside in the `corefunc` package with `package corefunc` declaration in the same directory as the existing `str_pad_test.go` file.
2. THE Property_Test_Suite SHALL use standard `testing.T` integration via `rapid.Check` so tests are discoverable and executable with `go test ./corefunc/...` without additional flags or tooling.
3. THE Property_Test_Suite SHALL follow the project naming convention by using a file named `str_pad_rapid_test.go`.
4. THE Property_Test_Suite SHALL import `pgregory.net/rapid` as the sole property-based testing dependency, while standard library packages and existing project packages (such as `testfixtures`) are permitted.
5. THE Property_Test_Suite SHALL name test functions using the pattern `TestXxx` (matching Go test naming conventions) so that they are included in default `go test` runs without `-run` filtering.
6. WHEN the full package test suite is executed via `go test ./corefunc/...`, THE Property_Test_Suite SHALL not cause failures in existing tests or alter their pass/fail outcomes.
