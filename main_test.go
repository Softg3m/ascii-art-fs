package main

import (
	"strings"
	"testing"
)

func TestGenerateAscii_InvalidBanner(t *testing.T) {
	result := GenerateAscii("A", "unknownbanner")

	expected := "Error reading banner file"

	if strings.TrimSpace(result) != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
