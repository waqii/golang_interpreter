/*
Description: Go test function for a lexer package
During the test, each token produced by the lexer is compared against expected values.
If the actual type or literal does not match the expected ones,
the test will fail, indicating that there’s an issue with the lexer implementation.
This approach helps ensure that the lexer behaves as intended and can accurately parse the input into tokens.
*/

package lexer

import (
	"testing"

	"monkey/token"
)

// TestNextToken tests the NextToken method of the lexer
// Test Cases: includes different token types and literals such as ASSIGN, PLUS, {}, etc.
func TestNextToken(t *testing.T) {
	input := `=+(){},;` // A new lexer instance is created with the input string =+(){},;

	// The expected token type and literal values are stored in a slice of structs
	// Assertion Checks: The test asserts that the tokens returned by NextToken match the expected types and literals.
	tests := []struct {
		expectedType      token.TokenType
		epectectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		//expectedType: This field specifies the expected type of token that the lexer should return

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		//expectedLiteral: This field holds the exact string value that the lexer should return for a particular token.
		//For example, if the input is a + character, the expectedLiteral for the PLUS token type would be "+".

		if tok.Literal != tt.expectedLiteral {
			t.Fatal("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
