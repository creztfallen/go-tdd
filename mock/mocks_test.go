package mock

import (
	"bytes"
	"testing"
)

func TestCounting(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	Counting(buffer, sleeperSpy)

	result := buffer.String()
	expected := `3
2
1
Go!`

	if result != expected {
		t.Errorf("resultado %s, expected %s", result, expected)
	}

	if sleeperSpy.Calls != 4 {
		t.Errorf("não houve chamadas suficientes do sleeper, esperado 4, resultado %d", sleeperSpy.Calls)
	}
}
