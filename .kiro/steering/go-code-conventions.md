---
inclusion: fileMatch
fileMatchPattern: "**/*.go"
---

<!-- @config-manager:start zdp -->
# Zero Diagnostics Policy

All Go source files must be free of diagnostic errors and warnings. Code is not considered complete until `getDiagnostics` reports zero issues for every file touched during a change.

## Philosophy

* **Small, surgical edits** are strongly preferred over wide-scale refactoring. Without a larger spec to guide the work, keep changes minimal and focused on the specific diagnostic being resolved.
* **Fix the problem, not the symptom.** Strongly favor resolving the root cause over suppressing the error. Suppression is a last resort reserved for cases where a fix is genuinely impossible or would introduce worse problems.
* The goal is **zero remaining diagnostic issues** when running `golangci-lint run --fix ./...` from the project root.

## Verification workflow

1. After editing any `.go` file, run `getDiagnostics` on that file.
2. If diagnostics are reported, fix every one before moving on.
3. After all fixes, run `golangci-lint run --fix ./...` from the project root to auto-fix what the linter can and surface anything remaining.
4. Run `go vet ./...` to confirm the build is clean.
5. Run `go test` for the affected package to confirm no regressions.

Do not present work as finished while diagnostics remain.

## Golangci-lint output handling

* When running `golangci-lint run`, focus on the final lines of the output beginning with `{number} issues:` to understand the scope of work.
* Save the full results to a temporary file for easy reference during the resolution process.
* None of the changes should break the code. The project must still compile after every change.

## Resolution workflow

When resolving lint issues across the project (not just a single file), follow this process:

1. Run `golangci-lint run` (no path filter) to produce the full report.
2. Parse the output to count how many errors each individual linter produced. Group errors by linter name (the identifier in parentheses at the end of each line, e.g., `(sloglint)`, `(err113)`, `(gocritic)`).
3. Sort the linters by error count ascending — fewest errors first.
4. Resolve linters in that order, starting with the linter that has the fewest errors and working up to the linter with the most. This maximizes early progress and keeps changesets small.
5. After clearing all errors for a given linter, re-run `golangci-lint run` to confirm the count dropped and no new issues were introduced.
6. Continue until the full run reports zero issues.

## Linter categories and fix patterns

This project enables ~60 linters via `.golangci.yml` (v2 format, `default: none`). The most common diagnostic categories and their fixes are listed below.

### Declaration order (decorder)

The enforced order within a file is: `const`, `var`, `type`, `func`. Moving a declaration block to the wrong position triggers a warning.

**Multiple var declarations:** The `disable-dec-num-check: false` setting means only ONE `var` block is allowed per file. This applies to the package-level `var()` block AND any `var` declarations inside function literals that are assigned within that block (e.g., `RunE` closures inside Cobra command definitions).

Fix pattern: convert `var x type` inside these closures to short declarations:

* `var x string` → `x := ""`
* `var x []T` → `x := []T{}`
* `var x int` → `x := 0`
* `var x bool` → `x := false`
* Multi-var blocks (`var ( ... )`) → individual short declarations

Standalone functions (declared with `func` outside the `var()` block) are NOT affected — `var` inside those is fine because the linter counts declarations per file-level scope, not per function.

### Structured logging (sloglint)

* **context**: all slog calls must use `*Context` variants (`InfoContext`, `DebugContext`, `WarnContext`, `ErrorContext`). Pass the `context.Context` from the function signature.
* **no-raw-keys**: define string constants for log keys and reference them instead of inline strings.
* **forbidden-keys**: the keys `time`, `level`, `msg`, `source`, and `foo` must not be used.
* **key-naming-case**: all log keys must be `snake_case`.
* **kv-only**: use key-value pairs, not `slog.Attr` typed attributes.
* **msg-style**: log messages must be capitalized.

### Error handling (err113, wrapcheck, errorlint)

* **err113**: do not create dynamic errors with `fmt.Errorf("...")` alone. Define package-level sentinel errors with `errors.New` and wrap them: `fmt.Errorf("%w: %s", ErrSentinel, detail)`.
* **wrapcheck**: errors returned from external packages (including other internal packages) must be wrapped with `%w`.
* **errorlint**: use `errors.Is` / `errors.As` instead of `==` or type assertions on errors.

**Error message clarity:** All error messages — both sentinel definitions and dynamic details — must be understandable to users who have no familiarity with the source code. Avoid referencing internal variable names, struct fields, or implementation details. Instead, describe the problem in terms of what the user was trying to do and what went wrong. For example, prefer `"failed to read AWS config file"` over `"awsConfigFilePath read error"`.

**Sentinel errors location:** All sentinel errors live in `cmd/errors.go`. Group them by domain (auth, cache, config, input validation, preconditions, validation). Each sentinel uses `errors.New` with a short, stable message. Dynamic details are added at the call site via `fmt.Errorf("%w: %s", ErrSentinel, detail)`.

**Naming convention:** Sentinel errors use the `Err` prefix followed by a PascalCase domain noun and condition (e.g., `ErrConfigKeyNotSet`, `ErrAccountIdentAmbiguous`, `ErrCachePathLookup`).

**Test files:** Test-specific dynamic errors (e.g., `errors.New("spy: prompt called")`, `errors.New("boom")`) are acceptable with `// lint:allow_errorf` suppression on the same line. These are intentionally dynamic for test clarity.

**Comment format for exported vars in errors.go:** The `revive` linter requires each exported variable's doc comment to start with the variable name (e.g., `// ErrFoo indicates ...`). When sentinel errors are grouped in a single `var()` block, each error must have its own individual doc comment starting with its name — group comments like `// Authentication errors.` are not sufficient.

**Single var block rule:** The `decorder` linter requires exactly one `var()` block per file. Even `cmd/errors.go` must use a single `var ( ... )` block with individual doc comments inside it, not separate `var` declarations.

### Context propagation (contextcheck)

Functions that have access to a `context.Context` must pass it to callees that accept one. If a helper is called from a context-bearing function, add `ctx context.Context` to the helper's signature.

### Argument count (revive argument-limit)

`revive` enforces a maximum of 4 parameters per function. When a function exceeds this limit, extract all parameters (except the testing handle) into a named input struct and pass it by pointer.

**Naming convention:** name the struct after the function it serves, using the `Input` suffix. For example, `assertFileRuleEqual` → `FileRuleInput`.

**Testing functions:** `t *testing.T` and `t *rapid.T` must remain as direct parameters. All other arguments are extracted into the struct.

**Struct placement:** define the struct in the file's existing `type (...)` block. If no block exists yet, add one following the `const → var → type → func` order required by `decorder`.

**Pass by pointer:** always pass input structs by pointer so that large structs do not also trigger `gocritic hugeParam`. Apply this consistently even for structs below the 64-byte threshold, so the pattern is uniform across the codebase.

```go
// Wrong — too many parameters
func assertFileRuleEqual(t *rapid.T, expected, actual FileRule, label string, idx int)

// Right
type (
    FileRuleInput struct {
        Expected FileRule
        Actual   FileRule
        Label    string
        Idx      int
    }
)

func assertFileRuleEqual(t *rapid.T, input *FileRuleInput)
```

Call sites construct and take the address of the struct inline:

```go
assertFileRuleEqual(t, &FileRuleInput{
    Expected: expected.Files[i],
    Actual:   actual.Files[i],
    Label:    label,
    Idx:      i,
})
```

### Control flags (revive flag-parameter)

`revive` flags boolean parameters that switch between two code paths inside a function. Renaming the parameter does not fix the diagnostic — the type and usage pattern are what the linter checks.

The fix is to wrap the bool in a named options struct so the call site is self-documenting and the parameter is no longer a bare boolean. See [[Code conventions — Function parameters]] for the full pattern.

```go
// Wrong — bare bool control flag
func Discover(configFlag string, flagChanged bool) (*DiscoverResult, error) {
    if flagChanged { ... }
}

// Right
type DiscoverOptions struct {
    FlagChanged bool
}

func Discover(configFlag string, opts DiscoverOptions) (*DiscoverResult, error) {
    if opts.FlagChanged { ... }
}
```

When the calling function already has an options struct that contains the relevant bool (e.g., `SyncOptions.DryRun`), pass that struct directly rather than creating a redundant single-field wrapper.

### Parameter size (gocritic hugeParam)

Structs larger than 64 bytes should be passed by pointer. Common in this codebase: `config.FileRule` (88 bytes). Use `*config.FileRule` in function signatures and dereference at call sites where the original API expects a value.

Instead of allowing or suppressing `hugeParam`, work to adapt the code to use pointers instead. Even if there are a large number of call sites, fixing them is the goal.

### Variable shadowing (govet shadow)

The `govet` linter with `shadow` enabled flags any `:=` declaration that shadows a variable from an outer scope. Common patterns and fixes:

* **`if err := someFunc(); err != nil`** inside a function that already has `err` declared: rename the inner variable (e.g., `authErr`, `writeErr`, `fetchErr`, `delErr`).
* **Closures inside spinner/action callbacks** that use `accts, err := ...` where `err` exists in the enclosing function: rename to `fetchErr` or similar.
* **Package-level variables** (e.g., `userHomeDir`) redeclared inside a function: rename the local to `homeDir` or similar.
* **`if err := x.Close(); err != nil`** in cleanup blocks: rename to `closeErr`.
* **`if _, err := file.WriteString(...); err != nil`** when `err` is already in scope: rename to `writeErr`.

The fix should always rename the INNER variable, not the outer one. Choose a name that describes the operation (e.g., `mkdirErr`, `marshalErr`, `releaseErr`).

### Line length (lll)

Line length is 120 characters. Do not break too early. Do not break too late. Common sources and fixes:

* **Struct tags with `jsonschema` descriptions:** These are inherently long and cannot be broken across lines in Go. Use `// lint:ignore_length` on the same line. Remove alignment padding between tag groups to minimize length.
* **Format strings in `fmt.Sprintf`/`fmt.Errorf`:** Break the string using `+` concatenation across lines.
* **Long comments:** Wrap to the next line.
* **Test string literals:** Use `+` concatenation to break across lines.

**Struct tag alignment:** When struct tags trigger lll, first remove all alignment padding (extra spaces between tag groups like `json:`, `jsonschema:`, `toml:`). If still over 120 chars after removing padding, add `// lint:ignore_length`.

**Interaction with tagliatelle:** When adding `// lint:ignore_length` to struct fields with non-camelCase JSON tags (e.g., `json:"snake_case_name"`), also add `// lint:allow_format` on the same line to suppress the `tagliatelle` linter. These tag names are intentional (matching TOML config format).

### Comment line length

Comments, including any whitespace (where tabs count as 4 spaces), must not have individual lines longer than 80 characters. Wrap to the next line instead of continuing on the same line.

Wrong:

```go
// validateManagedMarkers checks all profiles at once, used by the validate
// command to give a comprehensive report rather than stopping at the first error.
```

Right:

```go
// validateManagedMarkers checks all profiles at once, used by the validate
// command to give a comprehensive report rather than stopping at the first
// error.
```

### Cognitive complexity (gocognit)

Functions with complexity > 20 trigger a warning. Reduce complexity by extracting helper functions for distinct logical branches (e.g., "handle existing destination", "initialize from stub"). Each helper should have a single responsibility and a clear doc comment.

Look for opportunities to extract sections of code into separate functions in order to reduce the overall complexity of the solution.

### Security (gosec)

* **G104** (unhandled errors): address by properly handling the errors. Do not suppress them.
* **G117** (dangerous policy): report to the user at the end of the job for follow-up. Do not silently suppress.
* **G302** (file permissions): obey all diagnostics. File permissions for directories must always be `0o0755`. File permissions for files must always be `0o0666` (to support Microsoft Windows).
* **G304** (file inclusion via variable): while dynamic file names are allowed, security handling must be extraordinarily robust. Do not suppress without addressing real-world underlying security issues. Read <https://go.dev/blog/osroot> for path-safety patterns.
* **G703** (unhandled error on type assertion): read <https://go.dev/blog/osroot> to learn how best to resolve the issue. Report to the user at the end of the job for follow-up.
* **G204** (command injection): already suppressed in `hook.RunHook` — do not add new instances without review.
* **G301** (directory permissions): use `0o0755` for `os.MkdirAll`.

### Duplicate code (dupl)

The `dupl` linter is sometimes overly sensitive. If there is an opportunity to extract duplicated blocks into shared code, evaluate that approach first. If there is not a straightforward way to do that, use a suppression comment.

### Magic numbers (mnd)

Convert raw numbers into appropriately-named constants, then use those constants in place of the raw numbers. If this cannot be done cleanly, report the issue in the final summary instead of suppressing it.

### Commented-out code (gocritic commentedOutCode)

Remove commented-out code blocks. If the code is needed for reference, move it behind a build tag or delete it and rely on version control.

### Nil safety

`SyncOptions.Logger` can be nil. Any function that uses the logger outside of `RunSync` (which sets the default) must guard against nil — use a `safeLogger` helper or check before calling.

### Test performance (huh TUI blocking)

`huh.NewInput().Run()` and other `huh` form calls start a bubbletea TUI program that **blocks indefinitely** without a TTY (common in CI and `go test`). Tests that call `cmd.RunE` directly will hang at the test timeout if the code path reaches any `huh` input.

**Fix pattern:** Provide all config values via viper so the command skips every `huh` prompt:

```go
asmConfig = viper.New()
asmConfig.Set("profile-name", "test-profile")
```

**Cobra context nil panic:** Calling `cmd.RunE(cmd, args)` directly (without Cobra's `Execute` dispatch) leaves `cmd.Context()` as nil. Any code that calls `context.WithTimeout(cmd.Context(), ...)` will panic with `cannot create context from nil parent`. Fix by calling `cmd.SetContext(context.Background())` before `RunE`.

### Sorting (slices package)

* **Avoid `sort.Slice`**: Use the appropriate function from the `slices` package instead.
  * Simple ascending sort: `slices.Sort(s)`
  * Custom comparator: `slices.SortFunc(s, cmp)`
  * Stable sort with comparator: `slices.SortStableFunc(s, cmp)`
* **String comparators**: When the sort key is a string, use `strings.Compare` as the comparator rather than inline `<`/`>` expressions.

```go
// Wrong
sort.Slice(items, func(i, j int) bool { return items[i].Name < items[j].Name })

// Right
slices.SortFunc(items, func(a, b Item) int { return strings.Compare(a.Name, b.Name) })
```

### Performance

Using `fmt.Println`, `fmt.Printf`, `fmt.Print`, `fmt.Sprintln`, `fmt.Sprintf`, `fmt.Sprint`, `fmt.Fprintln`, `fmt.Fprintf`, or `fmt.Fprint` often degrades performance. There is performance overhead every time we write to stdout or stderr. We can improve performance by using `strings.Builder` to collect what needs to be written up-front with `fmt.Fprint*`, then printed to stdout or stderr with standard methods from the `fmt` package. This should be the default pattern when printing text.

### Suppression comment scope

`lint:allow_unhandled` suppresses `gosec` G104 but does NOT suppress `errcheck`. `lint:allow_defer_close` suppresses `errcheck` but does NOT suppress `gosec`. When both linters flag the same unhandled error, prefer handling the error properly (e.g., `if closeErr := f.Close(); closeErr != nil { t.Fatal(closeErr) }`) rather than stacking suppression comments.

## Code conventions

These rules apply to all Go code in the project, independent of linter enforcement.

### Comment blocks

When making any code change, review and update all comment blocks relevant to the changed code. Doc comments on functions, types, constants, and variables must remain accurate. Do not leave stale comments that describe old behavior.

### Interfaces

Anonymous interfaces are not allowed. Every interface must be named and defined at the package level, separately from the functions or types that use it.

```go
// Wrong
func Process(r interface{ Read() ([]byte, error) }) {}

// Right
type Reader interface {
    Read() ([]byte, error)
}

func Process(r Reader) {}
```

### Function parameters

If a function requires more than 3 input parameters (excluding `context.Context`), group them into a named struct and pass it as a single argument. `context.Context` must never be included in a struct — always pass it as a direct argument.

```go
// Wrong — too many parameters
func Sync(ctx context.Context, source, dest, profile string, dryRun bool) error

// Right
type SyncOptions struct {
    Source  string
    Dest    string
    Profile string
    DryRun  bool
}

func Sync(ctx context.Context, opts SyncOptions) error
```

### Function return values

If a function returns more than 3 values (excluding `error`), group them into a named struct. `context.Context` must never be included in a struct.

```go
// Wrong — too many return values
func Parse() (string, int, bool, error)

// Right
type ParseResult struct {
    Name  string
    Count int
    Valid bool
}

func Parse() (ParseResult, error)
```

### os.Exit

`os.Exit` must only be called from the outermost functions in the call chain: `main` or `init`. All other functions must signal failure by returning a wrapped error via `fmt.Errorf` or a sentinel via `errors.New`. This ensures deferred cleanup runs and callers can handle errors.

### Windows path compatibility

Always use `filepath.Join` (from `path/filepath`) when constructing file paths. Never concatenate path segments with `/` directly. When locating executables, use `exec.LookPath` rather than appending `.exe` manually.

```go
// Wrong
path := dir + "/" + filename

// Right
path := filepath.Join(dir, filename)
```

### Error check separation

Do not combine a function call and a nil-check in the same `if` statement using `;`. Separate the call from the check.

```go
// Wrong
if err := someFunc(); err != nil {

// Right
err := someFunc()
if err != nil {
```

### File locking

Implement file locking with platform-specific files using file name suffixes rather than inline build guards:

* `_unix` suffix — non-Windows builds. Use `golang.org/x/sys/unix.Flock`.
* `_windows` suffix — Windows builds. Use `golang.org/x/sys/windows.LockFileEx`.

### Git operations

When Git functionality is required, prefer the `go-git` library
(`github.com/go-git/go-git/v5`) over shell `git` invocations. Only fall back
to `os/exec` when the required functionality is not available from the library.

### Performance

Using `fmt.Println`, `fmt.Printf`, `fmt.Print`, `fmt.Sprintln`, `fmt.Sprintf`, `fmt.Sprint`, `fmt.Fprintln`, `fmt.Fprintf`, or `fmt.Fprint` often degrades performance. There is performance overhead every time we write to stdout or stderr. We can improve performance by using `strings.Builder` to collect what needs to be written up-front with `fmt.Fprint*`, then printed to stdout or stderr with standard methods from the `fmt` package. This should be the default pattern when printing text.

## Suppression

When a diagnostic cannot be resolved cleanly, use the project's lint suppression comments. Always include a justification. Suppression is a last resort — STRONGLY prefer fixing the root cause.

**Critical rules:**

* NEVER use `// nolint:xxx` directives. All valid and supported suppression comments begin with `// lint:`.
* Do NOT fall back to suppression comments except as a last resort. The goal is to resolve the issues, not hide them. Even if there are a large number of call sites, fixing them is the goal.
* When function signatures are split across multiple lines, and there is just cause to suppress an error, the suppression comment must be on the same line that triggers the diagnostic error.
* Context should never be passed as part of a struct. It must be passed as a direct argument.
* For anything deferred (not fixed), present and explain it to the user at the end of the job so the user can follow-up.
* Report on anything using `lint:allow_possible_insecure` as a suppression comment so the user can follow-up.

### Available suppression comments

These are defined in `.golangci.yml` under `exclusions.rules` and matched via `source:` patterns:

| Comment                        | Suppresses                          |
|--------------------------------|-------------------------------------|
| `lint:ignore_length`           | `lll` (line length)                 |
| `lint:allow_666`               | `gosec` (file permissions)          |
| `lint:allow_755`               | `gosec`, `mnd` (dir permissions)    |
| `lint:allow_possible_insecure` | `gosec`                             |
| `lint:allow_param`             | `unparam`                           |
| `lint:allow_raw_number`        | `mnd` (magic numbers)               |
| `lint:allow_nesting`           | `nestif`                            |
| `lint:no_dupe`                 | `dupl`                              |
| `lint:allow_errorf`            | `err113`, `goerr113`, `staticcheck` |
| `lint:allow_tls_min_version`   | `staticcheck` SA1019, `gosec` G402  |
| `lint:not_a_secret`            | `gosec` G101                        |
| `lint:allow_unhandled`         | `gosec` G104                        |
| `lint:allow_dynamic_filename`  | `gosec` G304                        |
| `lint:not_crypto`              | `gosec` G404                        |
| `lint:allow_large_memory`      | `gocritic` hugeParam                |
| `lint:allow_format`            | `gofumpt`, `tagliatelle`            |
| `lint:allow_init`              | `gochecknoinits`                    |
| `lint:allow_cuddle`            | `wsl`                               |
| `lint:allow_complexity`        | `gocognit`                          |
| `lint:no_const`                | `goconst`                           |
| `lint:allow_defer_close`       | `errcheck`                          |

Additionally, `errcheck` is automatically suppressed on lines containing `fmt.Fprint` or `fmt.Println` calls (matched via source regex, no comment needed).

The `.golangci.yml` file may NEVER be edited as a way to resolve diagnostics.
<!-- @config-manager:end zdp -->
