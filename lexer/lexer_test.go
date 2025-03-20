package lexer

import (
	"os"
	"testing"
)

func TestTokenize(t *testing.T) {
	// Read the content of lib.rs
	content, err := os.ReadFile("./lib.rs")
	if err != nil {
		t.Fatalf("Failed to read lib.rs: %v", err)
	}

	// Tokenize the content
	tokens := Tokenize(string(content))

	// Verify the output
	// Note: This is a basic verification. You can add more detailed checks based on expected tokens.
	if len(tokens) == 0 {
		t.Errorf("Expected tokens, but got none")
	}

	// Print the tokens for manual verification
	for _, token := range tokens {
		t.Logf("Token: %+v", token)
	}
}
