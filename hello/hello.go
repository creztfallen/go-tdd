package main

import "fmt"

func main() {
	fmt.Println(hello("Macumba"))
}

func hello(name string) string {
	return "Hello, " + name + "."
}
