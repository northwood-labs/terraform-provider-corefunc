# Contribution Guide

## Initial setup

The majority of development is done on macOS, so we have some helpers set-up to streamline setup.

### Prerequisites

* [Xcode CLI tools](https://mac.install.guide/commandlinetools/1.html)
* [Homebrew](https://mac.install.guide/homebrew/index.html)

### Setting-up

1. For Mac users, you can install all of the Homebrew and Go packages together.

    ```bash
    make install-tools-mac
    ```

    Obviously, this won't work on Linux, so ensure that the following packages are installed from your system's package manager.

    * [Go] 1.21+ (primary language)
    * [Node.js] 18+ (linting tools)
    * [Python] 3.11+ (linting tools)
    * [Git LFS] (storage of binary data)
    * [`jq`][jq] (shell scripting tools)
    * [`pre-commit`][pre-commit] (linting tools)

1. Running `make` in the root of the repo, by itself, will display a list of tasks and what they do. The ones highlighted in yellow are the ones that are most frequently used, or combine running multiple sub-tasks with one convenient command.

    ```bash
    make
    ```

1. After the core dependencies are installed, install the Git hooks. These will validate code on commit and reject anything that does not meet this project's best practices.

    ```bash
    make install-hooks
    ```

1. This project uses [Conventional Commits](https://www.conventionalcommits.org). If you are unfamiliar with this pattern of writing your commit messages, please read the spec. This project supports the following keywords:

    | Keyword    | Description                                                                        |
    |------------|------------------------------------------------------------------------------------|
    | `build`    | Changes that affect the build system or external dependencies.                     |
    | `ci`       | Changes to our CI configuration files and scripts.                                 |
    | `deps`     | Updating dependencies.                                                             |
    | `docs`     | Documentation only changes.                                                        |
    | `feat`     | A new feature.                                                                     |
    | `fix`      | A bug fix.                                                                         |
    | `lint`     | Linter or static analysis-related changes.                                         |
    | `perf`     | A code change that improves performance.                                           |
    | `refactor` | A code change that neither fixes a bug nor adds a feature.                         |
    | `style`    | Changes that do not affect the meaning of the code (whitespace, formatting, etc.). |
    | `test`     | Adding missing tests or correcting existing tests, benchmarks, fuzzing, etc.       |

<!-- Code of Conduct? -->

## Contributing a code change

In order to contribute a code change, you should fork the repository, make your changes, and then submit a pull request. Crucially, all code changes should be preceded by an issue that you've been assigned to. If an issue for the change you'd like to introduce already exists, please communicate in the issue that you'd like to take ownership of it. If an issue doesn't yet exist, please create one expressing your interest in working on it and discuss it first, prior to working on the code. Code changes without a related issue will generally be rejected.

In order for a code change to be accepted, you'll also have to accept the (very lightweight) [Developer Certificate of Origin](https://developercertificate.org) (DCO). Acceptance is accomplished by signing-off your commits; you can do this by adding a `Signed-off-by` line to your commit message, like here:

```plain
This is my commit message

Signed-off-by: Random Developer <random@developer.example.org>
```

Git has a built-in flag to append this line automatically:

```bash
git commit -s -m 'This is my commit message'
```

You can find more details about the _Developer Certificate of Origin_ checker in the [DCO app repo](https://github.com/dcoapp/app).

<!-- https://cla-assistant.io? -->

## Build provider from source

This will build the provider for the current OS and CPU architecture, and install it into your Go binary path.

1. Clone this Git repository.

1. Build this Terraform provider from source.

    ```bash
    make build
    ```

### Configure Terraform Provider development mode

By default, Terraform expects to download providers over the internet. Instead, we're going to download our own from GitHub Enterprise, and we need to tell Terraform how to find it.

1. Find where your installed Go binaries live.

    ```bash
    ./find-go-bin.sh
    ```

1. Update your `~/.terraformrc` file with these contents, replacing `<GO_BIN_PATH>` with the result from `./find-go-bin.sh`.

    ```hcl
    provider_installation {
        dev_overrides {
            "northwood-labs/corefunc" = "<GO_BIN_PATH>"
        }

        # Other configuration options
    }
    ```

## Ensure the provider is installed correctly

```bash
cd ./examples/data-sources/corefunc_str_truncate_label

# As a developer using `make build`, SKIP `terraform init`.
terraform plan
```

> **IMPORTANT:** **DO NOT** run `terraform init`. When running internal/custom providers, this can break things.

### Success looks like…

```plain
data.corefunc_str_truncate_label.label: Reading...
data.corefunc_str_truncate_label.label: Read complete after 0s

Changes to Outputs:
  + name = "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…"

You can apply this plan to save these new output values to the Terraform state, without changing any real infrastructure.
```

## Testing and fuzzing

The `nproc` binary is commonly available on most Linux distributions. If it's not installed, go back to the top of this document and follow instructions.

### Unit tests (and code coverage)

This will run Unit tests. This tests the low-level Go code, but not the Terraform integration wrapped around it.

```bash
# Run all unit tests
make unit

# Run one unit test
make unit NAME=TruncateLabel
```

You can view the code coverage report with either:

```bash
make view-cov-cli  # on the CLI
make view-cov-html # in the browser
```

#### How to write a unit test

In the `testfixtures/` directory, there is a file for each function. This leverages a pattern known as [Table-Driven Testing](https://github.com/golang/go/wiki/TableDrivenTests), and composes all of the test cases in one place.

The `corefunc/` directory contains the code and tests for the Go library. The `corefuncprovider/` directory contains the code and tests for the Terraform provider. The relevant `_test.go` files in each directory leverage the same test fixture. This ensures that the Go library code and the Terraform provider which implements the Go library function both pass the test cases.

If we leverage a third-party package for functionality, there will not be a test in the `corefunc/` directory, but the Terraform provider implementation will have a test in the `corefuncprovider/` directory.

* A unit test function lives in the `corefunc/` directory, and begins with the word `Test`.
* [Tutorial: Add a test](https://go.dev/doc/tutorial/add-a-test)
* [Documentation: Coverage profiling](https://go.dev/testing/coverage/)
* [Documentation: Testing flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags)
* [Tutorial: Code coverage for Go integration tests](https://go.dev/blog/integration-test-coverage)

### Terraform provider acceptance tests (and code coverage)

This will run Acceptance tests. Acceptance tests run the code through Terraform and test the Terraform communication pathway.

```bash
# Run all acceptance tests
make acc

# Run one acceptance test
make acc NAME=TruncateLabel

# Run all acceptance tests with debug-level output
make acc DEBUG=true

# Run one acceptance test with debug-level output
make acc NAME=TruncateLabel DEBUG=true
```

You can view the code coverage report with either:

```bash
make view-cov-cli  # on the CLI
make view-cov-html # in the browser
```

#### How to write an acceptance test

In the `testfixtures/` directory, there is a file for each function. This leverages a pattern known as [Table-Driven Testing](https://github.com/golang/go/wiki/TableDrivenTests), and composes all of the test cases in one place.

The `corefunc/` directory contains the code and tests for the Go library. The `corefuncprovider/` directory contains the code and tests for the Terraform provider. The relevant `_test.go` files in each directory leverage the same test fixture. This ensures that the Go library code and the Terraform provider which implements the Go library function both pass the test cases.

If we leverage a third-party package for functionality, there will not be a test in the `corefunc/` directory, but the Terraform provider implementation will have a test in the `corefuncprovider/` directory.

To help keep things easy to understand, the acceptance tests use Go's [`text/template`](https://pkg.go.dev/text/template) package to generate the Terraform code that is used for the acceptance test. The documentation for the [`templatefile()`](https://developer.hashicorp.com/terraform/language/functions/templatefile) function says that templates for use with Terraform should use the `*.tftpl` extension:

> `*.tftpl` is the recommended naming pattern to use for your template files. Terraform will not prevent you from using other names, but following this convention will help your editor understand the content and likely provide better editing experience as a result.

* An acceptance test function lives in the `corefuncprovider/` directory, and begins with the word `TestAcc`.
* [Tutorial: Add a test](https://go.dev/doc/tutorial/add-a-test)
* [Documentation: Coverage profiling](https://go.dev/testing/coverage/)
* [Documentation: Testing flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags)
* [Tutorial: Code coverage for Go integration tests](https://go.dev/blog/integration-test-coverage)

### Documentation Examples as tests (and code coverage)

This will run the Documentation Examples as tests. This ensures that the examples we put in front of users actually work.

```bash
# Run all example tests
make examples
```

You can view the code coverage report with either:

```bash
make view-cov-cli   # on the CLI
make view-cov-html # in the browser
```

#### How to write a documentation example

The `corefunc/` directory contains the code and tests for the Go library. If we leverage a third-party package for functionality, there will not be a test in the `corefunc/` directory.

* An example test function lives in the `corefunc/` directory, and begins with the word `Example`.
* [Tutorial: Documentation examples](https://go.dev/blog/examples)

### Test the provider schema as tests

This will use `tfschema` (reads the provider schema) and `bats` (CLI testing framework) to verify that the provider exposes the correct schema. This test requires compiling and installing the provoider (which `go test` doesn't require.)

```bash
# Run all BATS tests
make bats
```

#### How to write a BATS test

### Fuzzer

Fuzzing is a type of automated testing which continuously manipulates inputs to a program to find bugs. Go fuzzing uses coverage guidance to intelligently walk through the code being fuzzed to find and report failures to the user. Since it can reach edge cases which humans often miss, fuzz testing can be particularly valuable for finding security exploits and vulnerabilities.

This will run the fuzzer for 10 minutes. [Learn more about fuzzing](https://go.dev/doc/tutorial/fuzz).

```bash
# May only run one fuzz test at a time
make fuzz NAME=TruncateLabel
```

#### How to write a fuzz test

The `corefunc/` directory contains the code and tests for the Go library. If we leverage a third-party package for functionality, there will not be a test in the `corefunc/` directory.

* A fuzzer test function lives in the `corefunc/` directory, and begins with the word `Fuzz`.
* [Documentation: Go fuzzing](https://go.dev/security/fuzz/)
* [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)

## Benchmarks

Benchmarks test the performance of a package.

### Quick benchmarks

You can run a _quick_ benchmark (relatively speaking) that gets some data we can measure. It benchmarks the results of:

* Each test
* …with each test case
* …serially as well as in parallel

```bash
make quickbench
```

### Benchmark analysis

If you want to compare benchmarks between code changes, you'll need to run a _more complete_ benchmark which will provide enough data for analysis.  It benchmarks the results of:

* Each test
* …with each test case
* …serially as well as in parallel
* …multiple times over

```bash
make bench
```

Benchmark files are timestamped, so you can compare benchmark results across runs. See [benchstat documentation](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) to understand how to most effectively slice the results.

> [!NOTE]
> If you put old (previous) before new (current), values in the right-most column represent the size of the decrease (negative values are better). If you reverse them and push new (current) before old (previous), values in the right-most column represent the size of the increase (positive values are better).

```bash
benchstat Current=__bench-{new}.out Previous=__bench-{old}.out
```

It will show you something like this:

```plain
goos: darwin
goarch: arm64
pkg: github.com/northwood-labs/terraform-provider-corefunc/corefunc
                                       │    Current    │                 Previous                 │
                                       │    sec/op     │     sec/op      vs base                  │
TruncateLabel/balanced0-10                2.100n ±  1%   103.600n ±  3%   +4833.33% (p=0.002 n=6)
TruncateLabel/balanced3-10                2.104n ±  0%   103.750n ±  2%   +4831.08% (p=0.002 n=6)
TruncateLabel/balanced5-10                2.103n ±  1%   104.000n ±  1%   +4845.32% (p=0.002 n=6)
TruncateLabel/balanced7-10                69.41n ±  3%    213.20n ±  1%    +207.16% (p=0.002 n=6)
TruncateLabel/balanced8-10                68.75n ±  2%    214.55n ±  4%    +212.07% (p=0.002 n=6)
TruncateLabel/balanced10-10               67.67n ±  4%    215.85n ±  2%    +219.00% (p=0.002 n=6)
TruncateLabel/balanced30-10               72.97n ±  5%    220.50n ±  5%    +202.18% (p=0.002 n=6)
TruncateLabel/balanced64-10               74.80n ±  1%    220.05n ±  1%    +194.20% (p=0.002 n=6)
TruncateLabel/balanced80-10               33.37n ±  3%    103.45n ±  2%    +210.06% (p=0.002 n=6)
TruncateLabel/balanced100-10              33.83n ±  2%    103.45n ±  1%    +205.75% (p=0.002 n=6)
TruncateLabel/balanced120-10              33.48n ±  1%    103.75n ±  1%    +209.84% (p=0.002 n=6)
TruncateLabel/balanced128-10              34.49n ±  5%    103.90n ±  1%    +201.20% (p=0.002 n=6)
TruncateLabel/balanced256-10              34.15n ±  5%    104.20n ±  3%    +205.17% (p=0.002 n=6)

[…snip…]
```

I had to go into each file with VS Code, select the benchmarking lines, and sort _ascending_, then sort _numerically_ to get the lines in an order that makes sense.

The way that these are written, `TruncateLabel` is the test. `balanced{number}` is the test-case. And the `-10` is the number of CPU cores I have on my machine.

For the middle part (`balanced{number}`), the number represents the number of characters I truncated the label to. I know from the implementation that different truncation lengths can trigger different code paths, and that anything under `6` results in a near-immediate return with no calculation necessary. We also know that the longer lengths similarly have less truncation logic to perform.

But the middle tests from ~10–80 are most likely to execute _all_ of the code in the function, which makes it the most interesting.

#### How to write a benchmark suite

The `corefunc/` directory contains the code and tests for the Go library. If we leverage a third-party package for functionality, there will not be a test in the `corefunc/` directory.

* A fuzzer test function lives in the `corefunc/` directory, and begins with the word `Benchmark`.
* There is one `Benchmark` test which runs the tests serially. There is a second `Benchmark` test which runs the tests in parallel. This latter function has `Parallel` as the suffix of its name.
* [API Reference: Benchmarks](https://pkg.go.dev/testing@go1.21.1#hdr-Benchmarks)
* [API Reference: Benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)

### Exploring profiler data

Then, you can view the CPU profiler results…

```bash
make view-cpupprof
```

…view the memory profiler results…

```bash
make view-mempprof
```

…or view the trace results.

```bash
make view-trace
```

When you're done, clean-up.

```bash
make clean-bench
```

#### How to understand profiling and tracing data

* [Tutorial: Profiling Go Programs](https://go.dev/blog/pprof)
* [Documentation: Diagnostics](https://go.dev/doc/diagnostics)
* [Go: The Complete Guide to Profiling Your Code](https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz)

## Profile-guided optimization

Profile-guided optimization (PGO), also known as feedback-directed optimization (FDO), is a compiler optimization technique that feeds information (a profile) from representative runs of the application back into to the compiler for the next build of the application, which uses that information to make more informed optimization decisions. For example, the compiler may decide to more aggressively inline functions which the profile indicates are called frequently.

* [Documentation: Profile-guided optimization](https://go.dev/doc/pgo)
* [Tutorial: Profile-guided optimization in Go 1.21](https://go.dev/blog/pgo)

## Scanning for vulnerabilities

```bash
make vuln
```

## Previewing documentation

### Terraform Documentation

Generate the Terraform Registry-facing documentation.

```bash
make docs-provider
```

#### Terraform Provider Documentation

These are the patterns we follow for generating Terraform documentation. _Every_ resource/data source in the provider has at least one example. With a resource-specific template, we can implement multiple examples.

See `examples/data-sources/corefunc_env_ensure/` as an example of a custom template with multiple examples.

* [tfplugindocs: Conventional Paths](https://github.com/hashicorp/terraform-plugin-docs?tab=readme-ov-file#conventional-paths)
* [Terraform: Provider Documentation](https://developer.hashicorp.com/terraform/registry/providers/docs)
* [Terraform: Implement documentation generation](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-documentation-generation)

### Go CLI Documentation

You can get the package documentation on the CLI using the `go doc` command.

```bash
make docs-cli
```

#### How to write Go documentation comments and examples

* [Documentation: Go Doc Comments](https://tip.golang.org/doc/comment)
* [Tutorial: Testable Examples in Go](https://go.dev/blog/examples)

### Go Documentation Server

```bash
make docs-serve
```

#### How to write Go documentation comments and examples

* [Documentation: Go Doc Comments](https://tip.golang.org/doc/comment)
* [Tutorial: Testable Examples in Go](https://go.dev/blog/examples)

## Running the debugger

> [!IMPORTANT]
> At present, we only provide support for debugging in VS Code. Debugging through other tools may work, but they are untested.

If you are unfamiliar with using VS Code's debugging functionality for Go, see:

* <https://code.visualstudio.com/docs/editor/debugging>
* <https://github.com/golang/vscode-go/blob/master/docs/debugging.md>

To enable debugging for this Terraform provider:

1. In VS Code, open the extension pane for _Run and Debug_.

1. At the top of the pane, select _Debug Terraform Provider_ from the pulldown menu, then click the _Play_ icon (sideways triangle). The bottom statusbar will change color.

1. If the console does not open automatically, you can open it with `⌘ ⇧ P`, search for "debug", and select _Debug Console: Focus on Debug Console View_.

1. In the VS Code console, there will be a line of code that begins with `TF_REATTACH_PROVIDERS=` and then a JSON string. Copy the entire line, paste it into your Terminal, and **prefix** the command with the `export` keyword.

    ```bash
    export TF_REATTACH_PROVIDERS='{"@TODO":…}'
    ```

1. In your Terminal, go to the data source example directory.

    ```bash
    cd @TODO
    ```

1. Once in the example directory, run a Terraform plan.

    ```bash
    terraform plan
    ```

1. The Terraform provider will run and block (with "time elapsed" updating every 10 seconds or so). Any breakpoints you have set in your code will show up in the VS Code debugger, and you can review variable values and step through the breakpoints to understand what's happening in the code.

1. When you're done debugging, choose "Stop" (square) from the floating debugger controller. Alternatively, type `⇧ F5`. Once the debugger stops, and the bottom statusbar switches back to its standard color, cleanup your Terminal environment by unsetting the `TF_REATTACH_PROVIDERS` environment variable.

    ```bash
    unset TF_REATTACH_PROVIDERS
    ```

## Code

### Dotfiles

| File                 | Description                                                                                                                        |
|----------------------|------------------------------------------------------------------------------------------------------------------------------------|
| `.github/`           | GitHub repository project files.                                                                                                   |
| `.vscode/`           | Standard configurations for VS Code users. (This is not the same as `.idea/` folders for JetBrains IDEs. This is actually useful.) |
| `.dcignore`          | Configuration file for the [Snyk IDE Extension].                                                                                   |
| `.editorconfig`      | Configuration file for [EditorConfig].                                                                                             |
| `.gitattributes`     | Configuration file for Git. Helps ensure that file types are processed appropriately.                                              |
| `.gitignore`         | Configuration file for Git. Helps determine which files should be ignored by Git.                                                  |
| `.golangci.yml`      | Configuration file for [golangci-lint]. Helps ensure a high level of Go code quality.                                              |
| `.goreleaser.yml`    | Configuration file for [GoReleaser].                                                                                               |
| `.markdownlint.json` | Configuration file for [Markdownlint].                                                                                             |
| `.tflint.hcl`        | Configuration file for [tflint].                                                                                                   |
| `.yamllint`          | Configuration file for [yamllint].                                                                                                 |

### Project files

| File                               | Description                                                               |
|------------------------------------|---------------------------------------------------------------------------|
| `CHANGELOG.md`                     | The file containing the list of changes in each versionb of the software. |
| `CONTRIBUTING.md`                  | This document.                                                            |
| `find-go-bin.sh*`                  | Helps discover where installed Go executables live.                       |
| `LICENSE`                          | The license for this project.                                             |
| `Makefile`                         | A list of repeatable tasks used for managing this project.                |
| `README.md`                        | User-facing documentation for the project.                                |
| `terraform-registry-manifest.json` | A file that is required by the Terraform Registry.                        |
| `VERSION`                          | The version of the release of the software.                               |

### Source code

| File                | Description                                                                                                            |
|---------------------|------------------------------------------------------------------------------------------------------------------------|
| `cmd/`              | This provider also has a couple of subcommands. These are implemented using [Cobra].                                   |
| `corefunc/`         | The core Go library we provide to users. All core functionality is exposed through this library first.                 |
| `corefuncprovider/` | The files which wrap the methods from the Go library as Terraform data sources.                                        |
| `docs/`             | These are auto-generated from the `corefuncprovider/` files.                                                           |
| `examples/`         | These are HCL examples which implement the data sources, are used for real tests, and as examples used in the `docs/`. |
| `templates/`        | These are the templates used by the `tfplugindocs` app.                                                                |
| `testfixtures/`     | Fixtures used by the test suites for both the Go tests, as well as the Terraform tests.                                |
| `tools/`            | Includes the libraries that are part of the build process, but not part of the application.                            |
| `go.mod`            | Go modules definition.                                                                                                 |
| `go.sum`            | Go modules checksums.                                                                                                  |
| `main.go`           | Bootstraps the application binary.                                                                                     |

[Cobra]: https://cobra.dev
[EditorConfig]: https://editorconfig.org
[Git LFS]: https://git-lfs.com
[Go]: https://go.dev
[golangci-lint]: https://golangci-lint.run
[GoReleaser]: https://goreleaser.com
[jq]: https://jqlang.github.io/jq/
[Markdownlint]: https://www.npmjs.com/package/markdownlint-cli
[Node.js]: https://nodejs.org
[pre-commit]: https://pre-commit.com
[Python]: https://www.python.org
[Snyk IDE Extension]: https://docs.snyk.io/integrations/ide-tools/visual-studio-code-extension
[tflint]: https://github.com/terraform-linters/tflint
[yamllint]: https://github.com/adrienverge/yamllint
