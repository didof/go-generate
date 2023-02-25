package main

//go:generate go run gen.go

import "fmt"

func init() {
	fmt.Println("main.go#main.init")
}

func main() {
	fmt.Println("World!")
}
