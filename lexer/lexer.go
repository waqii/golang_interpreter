package lexer

import "monkey/token"

type Lexer struct { // defines a new struct names Lexer. A struct is a collection of fields/properties.
	input        string // input is the string being parsed
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

func New(input string) *Lexer { // Constructor function. it takes an input string and returns a pointer to a new Lexer instance.
	l := &Lexer{input: input} // : A new Lexer instance is created with the provided input, and a pointer to it is assigned to l.
	l.readChar()              // : The readChar method is called to initialize l.ch and l.readPosition.
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() { // readChar method reads the next character in the input and advances the position of the lexer.
	if l.readPosition >= len(l.input) { // If the readPosition is past the end of the input, the lexer assigns 0 to l.ch, indicating the end of the input.
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition] // Otherwise, the lexer reads the character at the current position and assigns it to l.ch.
	}
	l.position = l.readPosition // The lexer advances the position and readPosition by one character.
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string { // it reads in an identifier and advances our lexer’s positions until it encounters a non-letter-character.
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool { // isLetter function checks if a given character is a letter or an underscore.
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
	/*
		As you can see, in our case it contains the check ch == '_', which means that
		we’ll treat _ as a letter and allow it in identifiers and keywords. That means we can use variable
		names like foo_bar. Other programming languages even allow ! and ? in identifiers. If you
		want to allow that too, this is the place to sneak it in.
	*/
}

func isDigit(ch byte) bool { // isDigit function checks if a given character is a digit.
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
