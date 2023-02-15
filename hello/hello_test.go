package main

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Geralt")
		expected := "Hello, Geralt."

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("say 'Hello, world' when an empty string is passed", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world."

		verifyCorrectMessage(t, result, expected)
	})
}
