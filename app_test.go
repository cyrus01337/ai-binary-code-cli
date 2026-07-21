package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestApp(t *testing.T) {
	const EXPECTED_TEXT = "Hello, world!"
	const EXPECTED_BINARY = "01001000 01100101 01101100 01101100 01101111 00101100 00100000 01110111 01101111 01110010 01101100 01100100 00100001"

	t.Run("Encode", func(t *testing.T) {
		command := exec.Command("go", "run", ".", "--encode", EXPECTED_TEXT)
		var output bytes.Buffer
		command.Stdout = &output

		if err := command.Run(); err != nil {
			t.Fatalf("failed to run encode: %v", err)
		}

		trimmed := strings.TrimSpace(output.String())

		if trimmed != EXPECTED_BINARY {
			t.Errorf("expected %s, got %s", EXPECTED_BINARY, trimmed)
		}
	})

	t.Run("Decode", func(t *testing.T) {
		arguments := append([]string{"run", ".", "--decode"}, strings.Fields(EXPECTED_BINARY)...)
		command := exec.Command("go", arguments...)
		var output bytes.Buffer
		command.Stdout = &output

		if err := command.Run(); err != nil {
			t.Fatalf("failed to run decode: %v", err)
		}

		trimmed := strings.TrimSpace(output.String())

		if trimmed != EXPECTED_TEXT {
			t.Errorf("expected %s, got %s", EXPECTED_TEXT, trimmed)
		}
	})
}
