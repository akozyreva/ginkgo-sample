# ginkgo-sample

## Installation

```
go install github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
```

check the installation:

```
ginkgo version
```

## Project structure
```
mkdir books
cd books
ginkgo bootstrap  # create file books_suite_test.go
```

generate test file:

```
ginkgo generate book
```

to run tests:

```
ginkgo
```

to get detailed output:

```
ginko -v
```

## Troubleshooting:

to fix error like this:

```
../../../go-workspace/pkg/mod/github.com/onsi/ginkgo/v2@v2.27.5/core_dsl.go:25:2: missing go.sum entry for module providing package github.com/go-logr/logr (imported by github.com/onsi/ginkgo/v2); to add:
```
run:
```
go mod tidy
```
## Logging

Out-of-box tool: `GinkgoWriter` Sample:

```
GinkgoWriter.Println(data)
```

https://onsi.github.io/ginkgo/#mental-model-spec-timelines

