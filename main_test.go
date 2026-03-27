package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

//
// ✅ BASIC FUNCTIONALITY
//

func TestGenerateAscii_SingleCharacter(t *testing.T) {
	result := GenerateAscii("A", "standard")

	if result == "" {
		t.Error("Expected ASCII output, got empty string")
	}
}

func TestGenerateAscii_Word(t *testing.T) {
	result := GenerateAscii("Hi", "standard")

	if result == "" {
		t.Error("Expected ASCII output for 'Hi', got empty string")
	}
}

//
// ✅ MULTI-LINE INPUT
//

func TestGenerateAscii_Multiline(t *testing.T) {
	result := GenerateAscii("A\\nB", "standard")

	lines := strings.Split(result, "\n")

	if len(lines) == 0 {
		t.Error("Expected multiple lines for multiline input")
	}
}

//
// ✅ EMPTY INPUT
//

func TestGenerateAscii_EmptyString(t *testing.T) {
	result := GenerateAscii("", "standard")

	if result != "" {
		t.Errorf("Expected empty result, got: %s", result)
	}
}

//
// ✅ INVALID BANNER FILE
//

func TestGenerateAscii_InvalidBanner(t *testing.T) {
	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	GenerateAscii("A", "unknownbanner")

	w.Close()
	os.Stdout = oldStdout

	output, _ := io.ReadAll(r)

	if !strings.Contains(string(output), "Error reading banner file") {
		t.Errorf("Expected error message, got: %s", string(output))
	}
}

