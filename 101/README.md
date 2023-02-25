How to use `go:generate`?

## Setup

`terminal`
```bash
go mod init github.com/<username>/<reponame>
```

Create two files. `main.go` will use `gen.go` as generator.

`terminal`
```bash
touch main.go gen.go
```

`main.go` simply logs "World!" and runs the generator via the special comment.

_main.go_
```go
package main

//go:generate go run gen.go

import "fmt"

func init() {
	fmt.Println("main.go#main.init")
}

func main() {
	fmt.Println("World!")
}

```

> NOTE: There is no space between the slashes and `go:generate`.

> NOTE: not only `gen.go`, but `go run gen.go`.

`gen.go` declares `package main`, which would be forbidden were it not for the special comment.

_gen.go_
```go
//go:build ignore

package main

import "fmt"

func init() {
	fmt.Println("gen.go#main.init")
}

func main() {
	fmt.Println("Hello")
}

```

> NOTE: There is not space between the slashes and `go:build`.

## Usage

`terminal`
```bash
go generate ./...
```

`Output`
```output
gen.go#main.init
Hello
sh $
```

`terminal`
```bash
go run .
```

`terminal`
```output
main.go#main.init
World!
```

*Although the subcommand name is `generate`, the **generator** need not generate a file.*

