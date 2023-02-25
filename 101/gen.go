//go:build ignore

package main

import "fmt"

func init() {
	fmt.Println("gen.go#main.init")
}

func main() {
	fmt.Println("Hello")
}
