package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	result := buffer.String()
	expected := "Olá, Chris"

	if result != expected {
		t.Errorf("resultado '%s', esperado '%s'", result, expected)
	}
}
