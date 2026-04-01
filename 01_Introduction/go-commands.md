# The Essential Go Command Guide for Developers

## 1. Starting with Go Modules

Go modules help manage dependencies and versioning. When you start a new project, run:

```bash
go mod init myproject
````

This creates a `go.mod` file that declares your project and will track all your dependencies.

To tidy up dependencies later, use:

```bash
go mod tidy
```

It automatically adds missing ones and removes anything unused.

---

## 2. Building and Running Go Code

Running your application during development is simple:

```bash
go run main.go      # Runs a single file
go run .            # Runs all files in the current directory/package
```

Once you're ready to build a standalone binary:

```bash
go build            # Builds a binary (default name is your folder)
go build -o myapp   # Builds a binary with a custom name
```

To install the binary system-wide (in your Go bin path):

```bash
go install
```

### What does `go install` do?

Go compiles your code and installs the resulting binary into a specific directory on your system called the Go bin path.

This allows you to run your binary (executable) from anywhere in your terminal, without needing to be in your project directory.

Go will:

* Compile the program
* Place the binary in `~/go/bin` (or wherever `GOBIN` is set to)

If your module name is `github.com/you/mycli`, the binary will be named `mycli`.

You can now execute this binary directly from the command line by typing `mycli`, but only if the Go bin path is in your system’s `PATH`.

### What is the Go bin path?

The Go bin path is defined by the environment variable `GOBIN`. If you haven’t explicitly set `GOBIN`, it defaults to:

```bash
$GOPATH/bin
```

---

## 3. Adding and Managing Dependencies with `go get`

If you want to use an external package, `go get` is your go-to:

```bash
go get github.com/gin-gonic/gin
```

This downloads the package, adds it to your `go.mod`, and updates a file called `go.sum` to verify the download’s integrity.

Need a specific version?

```bash
go get github.com/pkg/errors@v0.9.1
```

---

## 4. Running Tests and Benchmarks

Go has a built-in testing framework. You just need to create test files ending with `_test.go` and use the `testing` package.

Here’s how to run tests:

```bash
go test             # Runs tests in the current package
go test -v          # Runs with verbose output
go test ./...       # Runs tests in all subdirectories
```

Check your test coverage:

```bash
go test -cover
```

### Sample output

```text
PASS
coverage: 85.7% of statements
ok      mymodule/mypkg    0.123s
```

Test coverage measures how much of your Go code is executed when running tests. It helps you understand how well your code is being tested.

Run performance benchmarks:

```bash
go test -bench .
```

### Sample output

```text
BenchmarkAdd-8       1000000000             0.24 ns/op
```

Benchmarks measure how fast or how resource-intensive a function is.

Think of them as speed tests for your code.

* Go determines how many iterations (`b.N`) to run to get a stable result.
* It runs your code in a loop and measures:

  * Time per operation (`ns/op`)
  * Allocations if `-benchmem` is used

---

## 5. Formatting and Code Quality

Go is famously opinionated about how code should look — and that's actually a huge win. It removes "style debates" from teams and encourages consistency across all Go codebases.

To format your code:

```bash
go fmt ./...
```

`go fmt` automatically rewrites your `.go` source files using standard Go formatting rules. No need for a separate linter or prettier tool.

To catch suspicious or incorrect code patterns:

```bash
go vet
```

`go vet` examines your code for common mistakes that might not be caught by the compiler but are often bugs or bad practices.

For deeper analysis, you can install additional tools:

While `fmt` and `vet` are good starting points, Go’s ecosystem offers deeper static analysis via powerful tools.

### `golint`

Suggests stylistic improvements and catches idiomatic issues.

```bash
go install golang.org/x/lint/golint@latest
```

### `staticcheck`

A super-linter for Go. Combines multiple checks:

* Unused variables
* Dead code
* Inefficient patterns
* Deprecated APIs
* And more

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
```

---

## 6. Exploring and Debugging

A few commands that are handy as you go deeper:

```bash
go doc fmt          # View docs for the fmt package
go env              # Check your Go environment (GOPATH, GOROOT, etc.)
go version          # Show your installed Go version
```

For debugging, tools like Delve are extremely helpful.

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

---

## 7. Understanding `go.sum`

The `go.sum` file contains cryptographic hashes of your dependencies. It helps Go verify the integrity of downloaded modules. You don’t need to manually edit it — Go will manage it for you.

Whenever you run `go get` or `go build` with new dependencies, Go updates `go.sum` to ensure everything is legit and secure.

---

## Quick Command Reference

```bash
go mod init <name>        # Start a new module
go mod tidy               # Clean up unused deps
go run .                  # Run the current package
go build -o binary        # Build the project
go get <pkg>              # Add a dependency
go test -v ./...          # Run all tests
go fmt ./...              # Format code
go vet                    # Analyze for bugs
```

```

Nếu bạn muốn, tôi có thể làm tiếp một bản **Markdown đẹp hơn để đăng blog** với:
- mục lục
- emoji/icon tiêu đề
- bảng cheat sheet cuối bài
- format tối ưu cho Notion / GitHub / Dev.to
```
