Like everything in existence, a *generator* is a function that (eventually) takes input and (eventually) returns output.

## Setup

> I will assume from this post on that `go mod init` have been executed (if not see 101).

`terminal`
```bash
touch main.go division.generator.go
```

`main.go`
```go
package main

//go:generate go run division.generator.go 12 4

import "fmt"

func main() {}
```

> NOTE: Adding the suffix `*.generator.go` it's not part of the convention it's something I do.

`division.go`
```go
//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	invoker := os.Getenv("GOFILE") // injected since it is executed via go:generate

	numerator, _ := strconv.Atoi(args[1])
	denominator, _ := strconv.Atoi(args[2])

	fmt.Println("args", args)
	fmt.Println("invoker", invoker)
	fmt.Printf("%d / %d = %d\n", numerator, denominator, numerator/denominator)
}

```

## Usage

`terminal`
```bash
go generate ./...
```

`output`
```bash
args [path/to/exe/division.generator 12 4]
invoker main.go
12 / 4 = 3
```

### A lot of env variables

`GOFILE` is not the only env variable provided.

```go
for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
	fmt.Println(ev, os.Getenv(ev))
}
```

`terminal`
```output
GOARCH arm64
GOOS darwin
GOFILE main.go
GOLINE 3 // lineno of //go:generate 
GOPACKAGE main
DOLLAR $
```

### Reuse across different projects

It is necessary to build the *generator* and place it in the `GOPATH`.

`terminal`
```bash
go build -o division division.generator.go
mv ./division $GOPATH/bin
```

`main.go`
```go
package main

//go:generate division 12 4

import "fmt"

func main() {}
```

> NOTE: No `go run`, just the name of the executable.

## References
- [https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate/](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate/)
- [https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go](https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go)