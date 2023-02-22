package mock

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

type StandardSleeper struct{}

func (s *StandardSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const lastWord = "Go!"
const countingStart = 3

func Counting(buffer io.Writer, sleeper Sleeper) {
	for i := countingStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, i)
		if i == 1 {
			sleeper.Sleep()
			fmt.Fprint(buffer, lastWord)
		}
	}

}
