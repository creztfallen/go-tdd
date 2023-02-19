package main

import (
	"fmt"
	"github.com/creztfallen/go-tdd/hello"
	"github.com/creztfallen/go-tdd/integers"
)

func main() {
	fmt.Println(hello.Hello("Macumba", "espanhol"))
	fmt.Println(integers.Add(2, 2))
}
