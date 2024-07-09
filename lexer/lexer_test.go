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
// Assertion Checks: The test asserts that the tokens returned by NextToken match the expected types and literals.
func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	!-/*5;
	### this code below add code to the tests add create tokens left off at 21:40 09/07/2021 to go padel - will update on return
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
	 	return false;
	}
	`

	tests := []struct { // The expected token type and literal values are stored in a slice of structs
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, // The first test case expects the lexer to return a token with the type LET and the literal "let".
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""}, // The last test case expects the lexer to return a token with the type EOF and an empty literal.
	}

	l := New(input) // A new lexer instance is created with the input string =+(){},;

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
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
