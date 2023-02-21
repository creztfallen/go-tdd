package main

import (
	"fmt"
	"net/http"
	// "os"

	di "github.com/creztfallen/go-tdd/dependency-injection"
)

func HandlerGreeting(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerGreeting))

	if err != nil {
		fmt.Println(err)
	}

	// di.Greet(os.Stdout, "Geralt")
}
