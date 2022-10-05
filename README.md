# goleak-example [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![GoReport][report-img]][report] [![Coverage Status][cov-img]][cov] [![GitHub go.mod Go version of a Go module][version-img]][version]

Example project which represents how to use `goleak` module to avoid Goroutine leaks.

## Prolog

Concurrency in Go materializes itself in the form of goroutines (independent activities) and channels (used for communication). While dealing with goroutines programmer needs to be careful to avoid their leakage

### Primary reasons of leakage:

1. The goroutine is waiting to read from a channel and the data never arrives.
2. The goroutine tries to write into a channel but blocked as the existing data is never read (buffered channel).

## Step 1. The Problem

Please checkout to the `problem` branch and explore `main.go` and `main_test.go`.
After running tests `make test` and `make cover` you can recognize that tests are passing and you have 100% of test coverage.
We are running in **False Negative** test situation.

```sh
$ git checkout problem

$ make test

go test -v -race ./...
=== RUN   Test_main
--- PASS: Test_main (2.00s)
PASS
ok      github.com/vbetsun/goleak-example       2.446s

$ make cover

go test -race -coverprofile=cover.out -coverpkg=./... ./...
ok      github.com/vbetsun/goleak-example       2.404s  coverage: 100.0% of statements in ./...
go tool cover -html=cover.out -o cover.html
```

## Step 2. Detecting

Next, checkout to the `detecting` branch 

```sh
$ git checkout detecting
```

At this step we have added `go.uber.org/goleak` package

```diff
import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
	"testing"
+
+	"go.uber.org/goleak"
)

+func TestMain(m *testing.M) {
+	goleak.VerifyTestMain(m)
+}
+
...
```
And after running test one more time - we have fixed our **False Negative** tests

```sh
$ make test

go test -v -race ./...
=== RUN   Test_main
--- PASS: Test_main (2.00s)
PASS
goleak: Errors on successful test run: found unexpected goroutines:
[Goroutine 7 in state chan send, with github.com/vbetsun/goleak-example.main.func1 on top of the stack:
goroutine 7 [chan send]:
github.com/vbetsun/goleak-example.main.func1(0x1)
        /goleak-example/main.go:15 +0x53
created by github.com/vbetsun/goleak-example.main
        /goleak-example/main.go:14 +0x85

 Goroutine 8 in state chan send, with github.com/vbetsun/goleak-example.main.func1 on top of the stack:
goroutine 8 [chan send]:
github.com/vbetsun/goleak-example.main.func1(0x2)
        /goleak-example/main.go:15 +0x53
created by github.com/vbetsun/goleak-example.main
        /goleak-example/main.go:14 +0x85
]
FAIL    github.com/vbetsun/goleak-example       2.837s
FAIL
make: *** [test] Error 1
```
So, now we have clear understanding that we have a goroutine leak and we are already able to fix it!

## Step 3. Solution

Next, checkout to the `solution` branch 
We are going to add `sync.WaitGroup` for synchronize all goroutines

Now< after running tests and coverage we will receive correct results

```sh
$ make test

go test -v -race ./...
=== RUN   Test_main
--- PASS: Test_main (2.00s)
PASS
ok      github.com/vbetsun/goleak-example 

$ make cover

go test -race -coverprofile=cover.out -coverpkg=./... ./...
ok      github.com/vbetsun/goleak-example       2.407s  coverage: 100.0% of statements in ./...
go tool cover -html=cover.out -o cover.html
```

[doc-img]: https://pkg.go.dev/badge/github.com/vbetsun/goleak-example?status.svg
[doc]: https://pkg.go.dev/github.com/vbetsun/goleak-example
[ci-img]: https://github.com/vbetsun/goleak-example/actions/workflows/ci.yml/badge.svg
[ci]: https://github.com/vbetsun/goleak-example/actions/workflows/ci.yml
[report-img]: https://goreportcard.com/badge/github.com/vbetsun/goleak-example
[report]: https://goreportcard.com/report/github.com/vbetsun/goleak-example
[cov-img]: https://codecov.io/gh/vbetsun/goleak-example/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/vbetsun/goleak-example
[version-img]: https://img.shields.io/github/go-mod/go-version/vbetsun/goleak-example.svg
[version]: https://github.com/vbetsun/goleak-example


