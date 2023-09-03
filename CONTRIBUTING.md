# Contribution Guide

## Build provider from source

> **IMPORTANT:** You probably want **"Install provider (as a user)"** instead.

This will build the provider for the current OS and CPU architecture, and install it into your Go binary path.

1. Clone this Git repository.

1. Build this Terraform provider from source.

    ```bash
    make build
    ```

### Configure Terraform Provider development mode

By default, Terraform expects to download providers over the internet. Instead, we're going to download our own from GitHub Enterprise, and we need to tell Terraform how to find it.

We're going to assume you already have [Homebrew](https://mac.install.guide/homebrew/index.html) and the [Xcode CLI tools](https://mac.install.guide/commandlinetools/1.html) installed.

1. Install [Go](https://go.dev). For Mac users, it's as simple as…

    ```bash
    brew install go
    ```

1. Find where your installed Go binaries live.

    ```bash
    ./find-go-bin.sh
    ```

1. Update your `~/.terraformrc` file with these contents:

    ```hcl
    provider_installation {
        dev_overrides {
            "northwood-labs/corefunc" = "<GO_BIN_PATH>"
        }

        # Other configuration options
    }
    ```

    1. Replace `<GO_BIN_PATH>` with the path to your Go binaries, above.

        ```bash
        # After running `make build`...
        dirname $(which terraform-provider-corefunc)
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

The `nproc` binary is commonly available on most Linux distributions. To install in Alpine Linux or macOS, you should install the `coreutils` package from `apk` (Alpine Linux) or `brew` (macOS).

### Unit tests (and code coverage)

This will run Unit and Fuzz tests. This tests the low-level Go code, but not the Terraform integration wrapped around it.

```bash
go test -count=1 -parallel=$(nproc) -timeout 30s -coverpkg=./... -coverprofile=_coverage.txt -v ./...
```

You can view the code coverage report with:

```bash
go tool cover -html=_coverage.txt
```

### Terraform provider acceptance tests (and code coverage)

This will run all tests: Unit, Acceptance, and Fuzz. Acceptance tests run the code through Terraform and test the Terraform communication pathway.

```bash
TF_ACC=1 go test -count=1 -parallel=$(nproc) -timeout 30s -coverpkg=./... -coverprofile=_coverage.txt -v ./...
```

You can view the code coverage report with:

```bash
go tool cover -html=_coverage.txt
```

### Fuzzer

This will run the fuzzer for 10 minutes. [Learn more about fuzzing](https://go.dev/doc/tutorial/fuzz).

```bash
go test -fuzz=Fuzz -fuzztime 10m -v ./corefunc/.
```

### Benchmarks

Benchmarks test the performance/scalability of a package. Different versions of the benchmarks execute different code paths inside the functions.

```bash
go test \
    -bench=. \
    -benchtime=10s \
    -benchmem \
    -cpuprofile=_cpu.out \
    -memprofile=_mem.out \
    -trace=_trace.out \
    ./corefunc \
    | tee _bench.txt
```

Then:

```bash
go tool pprof -http :8080 _cpu.out
```

```bash
go tool pprof -http :8081 _mem.out
```

```bash
go tool trace _trace.out
```

```bash
go install golang.org/x/perf/cmd/benchstat@latest
benchstat _bench.txt
```

When you're done, clean-up.

```bash
rm -f _*.out _*.txt *.test
```

## Scanning for vulnerabilities

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck -test ./...
```

```bash
go install github.com/google/osv-scanner/cmd/osv-scanner@v1
osv-scanner -r .
```

## Previewing documentation

### Terraform Documentation

Generate the Terraform Registry-facing documentation.

```bash
go generate -v ./...
```

### CLI Go Documentation

You can get the package documentation on the CLI using the `go doc` command.

```bash
go doc -C corefunc/ -all
```

### Local HTTP Go Documentation Server

```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc -index -links
```

Then open <http://localhost:6060/pkg/github.com/northwood-labs/terraform-provider-corefunc/corefunc/> in your web browser.

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
