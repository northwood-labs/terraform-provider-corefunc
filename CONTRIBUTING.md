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

    * Go 1.21+ (primary language)
    * Node.js 18+ (linting tools)
    * Python 3.11+ (linting tools)
    * `jq` (shell scripting tools)
    * `pre-commit` (linting tools)

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

## Build provider from source

> [!IMPORTANT]
> You probably want **"Install provider (as a user)"** instead.

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

This will run Unit and Fuzz tests. This tests the low-level Go code, but not the Terraform integration wrapped around it.

```bash
# Run all unit tests
make unit

# Run one unit test
make unit NAME=TruncateLabel
```

You can view the code coverage report with:

```bash
make view-cov
```

### Terraform provider acceptance tests (and code coverage)

This will run all tests: Unit, Acceptance, and Fuzz. Acceptance tests run the code through Terraform and test the Terraform communication pathway.

```bash
# Run all acceptance tests
make acc

# Run one acceptance test
make acc NAME=TruncateLabel
```

You can view the code coverage report with:

```bash
make view-cov
```

### Fuzzer

This will run the fuzzer for 10 minutes. [Learn more about fuzzing](https://go.dev/doc/tutorial/fuzz).

```bash
# May only run one fuzz test at a time
make fuzz NAME=TruncateLabel
```

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
> If you put old before new, values in the right-most column represent the size of the decrease (negative values are better). If you reverse them and push new before old, values in the right-most column represent the size of the increase (positive values are better).

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

But the middle tests from ~10–80 are most likely to execute _all_ of the code in the function, which makes it the most intersting.

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

### Go CLI Documentation

You can get the package documentation on the CLI using the `go doc` command.

```bash
make docs-cli
```

### Go Documentation Server

```bash
make docs-serve
```

## Running the debugger

> **IMPORTANT:** At present, we only provide support for debugging in VSCode. Debugging through other tools may work, but they are untested.

If you are unfamiliar with using VSCode's debugging functionality for Go, see:

* <https://code.visualstudio.com/docs/editor/debugging>
* <https://github.com/golang/vscode-go/blob/master/docs/debugging.md>

To enable debugging for this Terraform provider:

1. In VSCode, open the extension pane for _Run and Debug_.

1. At the top of the pane, select _Debug Terraform Provider_ from the pulldown menu, then click the _Play_ icon (sideways triangle). The bottom statusbar will change color.

1. If the console does not open automatically, you can open it with `⌘ ⇧ P`, search for "debug", and select _Debug Console: Focus on Debug Console View_.

1. In the VSCode console, there will be a line of code that begins with `TF_REATTACH_PROVIDERS=` and then a JSON string. Copy the entire line, paste it into your Terminal, and **prefix** the command with the `export` keyword.

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
