//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Reference <https://github.com/eliben/code-for-blog/blob/master/2021/go-generate-guide/insourcegenerator/mypack/gen.go>
// Reference <https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go>

func main() {
	args := os.Args

	invoker := os.Getenv("GOFILE")

	numerator, _ := strconv.Atoi(args[1])
	denominator, _ := strconv.Atoi(args[2])

	fmt.Println("args", args)
	fmt.Println("invoker", invoker)
	fmt.Printf("%d / %d = %d\n", numerator, denominator, numerator/denominator)

	for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
		fmt.Println(ev, os.Getenv(ev))
	}
}
