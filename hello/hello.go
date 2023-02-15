package main

import "fmt"

func main() {
	fmt.Println(Hello("Macumba"))
}

const prefixHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return prefixHello + name + "."
}
