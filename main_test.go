package main

import (
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
