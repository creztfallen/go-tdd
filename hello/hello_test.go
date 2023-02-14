package main

import "testing"

func TestHello(t *testing.T) {
	result := hello("Geralt")
	expected := "Hello, Geralt."

	if result != expected {
		t.Errorf("Resultado '%s', esperado '%s'", result, expected)
	}
}
