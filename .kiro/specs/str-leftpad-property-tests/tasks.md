# Implementation Plan: StrLeftPad Property-Based Tests

## Overview

Add property-based tests for the `StrLeftPad` function using `pgregory.net/rapid`. The implementation creates a new test file `corefunc/str_pad_rapid_test.go` containing six property test functions that verify algebraic invariants of `StrLeftPad` across randomly generated inputs. The `rapid` dependency is added to `go.mod` as a test-only dependency.

## Tasks

* [x] 1. Add rapid dependency and create test file scaffold
  * [x] 1.1 Add `pgregory.net/rapid` dependency to go.mod
    * Run `go get pgregory.net/rapid` to add the dependency
    * Verify it appears in `go.mod` and `go.sum`
    * This is a test-only dependency (imported only in `_test.go` files)
    * _Requirements: 7.4_

  * [x] 1.2 Create `corefunc/str_pad_rapid_test.go` with package declaration and imports
    * File must use `package corefunc` declaration
    * Import `strings`, `testing`, and `pgregory.net/rapid`
    * Add copyright header matching project convention
    * _Requirements: 7.1, 7.3, 7.4_

* [x] 2. Implement output length and content preservation properties
  * [x] 2.1 Implement `TestStrLeftPad_OutputLength`
    * Use `rapid.Check(t, ...)` integration pattern
    * Generate random strings (0–1024 bytes) and pad widths (0–2048)
    * Assert `len(result) == padWidth` when `len(str) < padWidth`
    * Assert `len(result) == len(str)` when `len(str) >= padWidth`
    * Assert `len(result) == len(str)` when `padWidth == 0`
    * Include property comment: `// Property 1: Output Length Invariant`
    * _Requirements: 1.1, 1.2, 1.3, 1.4_

  * [x] 2.2 Implement `TestStrLeftPad_ContentPreservation`
    * Generate random strings (0–1000 bytes), pad widths (0–1000), pad chars (0x01–0xFF)
    * Assert `strings.HasSuffix(result, str)` for all inputs including empty strings
    * Assert padding prefix bytes each equal padChar when `padWidth > len(str)`
    * Include property comment: `// Property 2: Content Preservation (Suffix)`
    * _Requirements: 2.1, 2.2, 2.3_

* [x] 3. Implement pad character correctness and idempotence properties
  * [x] 3.1 Implement `TestStrLeftPad_PadCharCorrectness`
    * Generate random strings, pad widths exceeding string length, and pad chars (0x00–0xFF full byte range)
    * Assert every byte in `output[0 : padWidth-len(str)]` equals the specified padChar
    * Verify default pad char is space (0x20) when no padChar is specified
    * Assert output equals input unchanged when `padWidth <= len(str)`
    * Include property comment: `// Property 3: Pad Character Correctness`
    * _Requirements: 5.1, 5.2, 5.3, 5.4_

  * [x] 3.2 Implement `TestStrLeftPad_Idempotence`
    * Generate strings (0–1000 bytes), pad widths (0–1000), pad chars (0x20–0x7E)
    * Assert `StrLeftPad(StrLeftPad(str, w, c), w, c) == StrLeftPad(str, w, c)`
    * Also verify that when `len(str) >= padWidth`, result equals input unchanged
    * Include property comment: `// Property 4: Idempotence`
    * _Requirements: 3.1, 3.2_

* [x] 4. Implement identity and monotonic growth properties
  * [x] 4.1 Implement `TestStrLeftPad_Identity`
    * Generate strings (0–200 bytes), negative pad widths (-1000 to 0), and pad chars
    * Assert `result == str` when `len(str) >= padWidth`
    * Assert `result == str` when `padWidth <= 0`
    * Verify varying padChar does not influence output when padWidth is non-positive
    * Include property comment: `// Property 5: Identity (No-Op)`
    * _Requirements: 3.3, 4.1, 4.2, 4.3, 4.4_

  * [x] 4.2 Implement `TestStrLeftPad_MonotonicGrowth`
    * Generate strings (0–200 bytes), pad widths from `len(str)+1` to 300, and pad chars
    * Assert `len(result) >= len(str)` for all inputs
    * Assert padding with `padWidth == len(str)` returns input unchanged
    * Assert `StrLeftPad(str, w+1, c) == string(c) + StrLeftPad(str, w, c)` when `w >= len(str)`
    * Assert output does NOT equal input when `padWidth > len(str)` and `len(str) > 0`
    * Include property comment: `// Property 6: Monotonic Growth (Metamorphic)`
    * _Requirements: 6.1, 6.2, 6.3, 6.5_

* [x] 5. Checkpoint - Verify all tests pass
  * Ensure all tests pass by running `go test ./corefunc/...`
  * Verify property tests are discoverable alongside existing tests
  * Confirm no existing tests are broken by the new file
  * Ensure all tests pass, ask the user if questions arise.
  * _Requirements: 7.2, 7.5, 7.6_

## Notes

* All tasks use Go as the implementation language with `pgregory.net/rapid` for property-based testing
* Each property test function uses `rapid.Check(t, ...)` for standard `testing.T` integration
* Generators produce random inputs within ranges specified by requirements
* The `rapid` library defaults to 100 iterations per property; use `-rapid.checks=1000` flag for extended coverage (Requirement 6.4)
* Rapid automatically shrinks failing inputs to minimal counterexamples for easy debugging
* Tasks marked with `*` are optional and can be skipped for faster MVP
* Each task references specific requirements for traceability

## Task dependency graph

```json
{
  "waves": [
    { "id": 0, "tasks": ["1.1"] },
    { "id": 1, "tasks": ["1.2"] },
    { "id": 2, "tasks": ["2.1", "2.2", "3.1", "3.2", "4.1", "4.2"] }
  ]
}
```
