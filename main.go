package main

import (
	"os"

	"github.com/creztfallen/go-tdd/mock"
)

func main() {
	sleeper := &mock.StandardSleeper{}
	mock.Counting(os.Stdout, sleeper)
}
