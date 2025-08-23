# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

When implementing new functionality, please follow these patterns exactly to maintain consistency.

You are assisting with a Terraform/OpenTofu provider named "corefunc". Generate code that strictly matches existing repository conventions.

## High-level structure

* Core logic: corefunc/
* Provider wrapper: corefuncprovider/
* For each feature X (snake_case file stem without corefunc_ prefix):
    * Library: corefunc/X.go
    * Data source: corefuncprovider/X_data_source.go
    * Provider function: corefuncprovider/X_function.go
    * Tests: *_data_source_test.go, *_function_test.go
    * Fixture templates: *_data_source_fixture.tftpl, *_function_fixture.tftpl
    * Shared table-driven cases: testfixtures/*.go
* Examples: examples/data-sources/corefunc_X/, examples/functions/corefunc_X/
* Docs generated from provider sources + templates.

## Headers and docs

* Every Go file begins with the Apache 2.0 header (years range updated) exactly as in existing files.
* Each package has a doc.go with a short package comment.
* Each exported function/method has a godoc block with:
    * One-sentence summary.
    * Blank line.
    * Parameter section using a hyphen list, matching existing formatting.
    * Return value description.
* In provider schema or function Definition markdown: wrap multi-line strings with strings.TrimSpace(dedent.Dedent(`...`)).
* Use backticks for identifiers and code.
* Use linkPackage("FuncName") helper when referencing underlying Go functions (see helpers).
* Mention Terratest usage phrase: “which can be used in Terratest” consistent with current wording.

## Naming and style

* Exported helpers: PascalCase (e.g. StrSnake, URLDecode).
* Internal struct types for data sources/functions: lowerCamel with suffix DataSource / Function (e.g. urlParseDataSource, strIterativeReplaceFunction).
* Receiver names: short (f, d).
* Keep attribute names stable: input attribute first, computed outputs after.
* Avoid introducing new dependencies unless absolutely necessary.

## Logging

* At start and end of every provider method: tflog.Debug(ctx, "Starting <Name> <Type> <Method> method.") and matching Ending line.
* Do not add excessive logs inside unless handling errors.

## Provider data sources

* Implement interfaces explicitly; add compile-time interface assertion lines (var _ datasource.DataSource = &xyzDataSource{}).
* Schema layout:
    * MarkdownDescription summarizing capability + mapping to underlying Go func with linkPackage().
    * Required arguments at top; Computed outputs at bottom.
    * Use types.StringType / types.Int64Type etc. and schema.StringAttribute / List / Map as already done.
* Read():
    * Get state into model struct.
    * Perform transformation via corefunc.<Func>.
    * Set state; append diagnostics; return early if diagnostics.HasError().

## Provider functions

* Provide Metadata(), Definition(), Run() with logging wrappers.
* Definition():
  * Summary + MarkdownDescription (dedent).
  * Parameters typed using function.StringParameter, ListParameter, MapParameter etc.
  * Return: function.StringReturn{}, or object/map mirroring existing examples (see url_parse / str_iterative_replace).
* Run():
  * Extract arguments with req.Arguments.Get.
  * On error: set resp.Error using function.ConcatFuncErrors.
  * Call corefunc.<Func>.
  * Set resp.Result; no panics.
  * Validate enum-like inputs (see canonicalizer switch in url_parse); produce function.NewArgumentFuncError on invalid values.

## Error handling

* Never panic in provider runtime paths.
* Add diagnostics.AddError for data source errors; for functions use function.NewFuncError / NewArgumentFuncError.
* Use clear, user-facing messages (“Could not parse the URL.”) plus underlying error detail.

## Testing

* Table-driven tests share central maps in testfixtures/*.go (e.g. URLParseTestTable, URLDecodeTestTable).
* Test naming:
    * Data source acceptance: TestAcc<PascalName>DataSource
    * Function acceptance: TestAcc<PascalName>Function
* Use template fixture (*.tftpl) rendered with text/template inside tests.
* For failing cases include ExpectError with regexp.MustCompile(".*") pattern (mirrors existing style).
* Example outputs in .tftpl start with '#=>' lines listing each expected line in order.
* Use resource.UnitTest or resource.Test depending on pattern already established.
* Keep providerConfig prefix reuse pattern.

## Fixture templates

* For functions: output blocks using provider::corefunc::<snake_name>(...) syntax.
* Keep spacing and trailing commas consistent with existing templates.
* Represent replacement lists/maps exactly as current examples (indentation, alignment).
* Each expected value echoed after config using '#=> ' lines; order must match test parsing expectations.

## Table-driven data

* Keys are descriptive, often prefixed with mode (e.g. "DEFAULT: ", "GOOGLE SAFE: ").
* Use lint override comments where raw numeric literals or duplication occur (e.g. // lint:allow_raw_number, // lint:no_dupe).

## Code generator compatibility

* Preserve patterns that the generator depends on (filenames *_data_source.go, *_function.go, *_fixture.tftpl).
* When adding a new feature ensure generator output shape remains compatible.

## Performance/cleanliness

* No premature optimization.
* Follow existing import grouping and ordering.
* Use strings.TrimSpace + dedent.Dedent for multi-line docs; never raw multi-line without trimming.

## Do not

* Introduce inconsistent logging messages.
* Change attribute names for existing schemas.
* Bypass shared testfixtures for overlapping logic.
* Add global mutable state.

## When unsure

* Mirror the closest existing analogous implementation:
    * Simple string transform: see str_snake function run path.
    * Multi-field object return: see url_parse function + data source.
    * Iterative replacement pattern: see str_iterative_replace.

## Output

* Always produce code already conforming to golangci-lint expectations in this repo (imports, formatting, receivers, error handling, logging, doc comments).
