package lexer

type Lexer struct { // defines a new struct names Lexer. A struct is a collection of fields/properties.
	input        string // input is the string being parsed
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

// Constructor function. it takes an input string and returns a pointer to a new Lexer instance.
func New(input string) *Lexer {
	l := &Lexer{input: input} // : A new Lexer instance is created with the provided input, and a pointer to it is assigned to l.
	l.readChar()              // : The readChar method is called to initialize l.ch and l.readPosition.
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
