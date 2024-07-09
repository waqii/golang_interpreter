# golang Interpreter - personal brain dump



This is a repository following my journey going through the book Writing An Interpreter In Go by Thorsten Ball.




Previously: 

Took a break from book. got to page 20. 

The book will be going through 

Lexical Analysis - This is the act of taking a sequence of characters and turning them into *lexical tokens*.

Tokens are characters/strings with a *meaning* and are *identified* for example.


	for(int i = 0; i < 10; ++i)_

for = keywords
( = open parent
int = keywords
) = close parent
i = symbol
= = equals
0 = numbers
; = semi colon

this act of taking source code and characterising it is called lexical analysis. 

A lexer can many types of tokens *whitespaces* *line numbers* *file name* 


![[Pasted image 20240429164411.png]]

 

Created the files

`lexer.go` -	
	lexer_test.go - Go test function for a lexer package During the test, each token produced by the lexer is compared against expected values. If the actual type or literal does not match the expected ones, the test will fail, indicating that there’s an issue with the lexer implementation. This approach helps ensure that the lexer behaves as intended and can accurately parse the input into tokens.
	

`token.go - used to define our token outputs. 


Lexer.go file -

`readChar()` -
	This method gives us the *next character* and advances our position in the *input* string. 
	First checks to see if *EOF* if yes - set l.ch to 0 - which is the ASCII for *NUL* character and signifies Either *end of file* or *we have not learnt anything yet*
		If not - sets l.ch to next character by accessing *l.input[l.readPosition]*.		*l.position* is updated to The *l.readPosition* which is incremented +1 
			This way
				l.readPosition always points to the next position
				l.position always points to the position where we last read.

`nextToken()` -
	The nextToken() function is a key part of the lexer in our Monkey interpreter. Its purpose is to read the input source code character by character and produce the next token in the sequence.
		It starts by skipping any whitespace characters using the `skipWhitespace()` method.
		It then uses a switch statement to handle different characters:
		    - For single-character tokens like `+`, `-`, `(`, `)`, etc., it simply creates a new token with the corresponding token type.
		    - For potential two-character tokens like `==` and `!=`, it peeks at the next character using `peekChar()` and creates the appropriate token.
		    - For string literals, it reads the entire string using `readString()` and creates a string token.
		    - For identifiers (e.g., variable names, keywords), it reads the identifier using `readIdentifier()`, looks up its type using `LookupIdent()`, and returns the token.
		    - For integer literals, it reads the number using `readNumber()` and creates an integer token.
		    - If the character is not recognized, it creates an illegal token.
		After the switch statement, it advances to the next character using `readChar()` and returns the token.

`isLetter ()`
This function takes a byte (a single character) as input and returns a boolean indicating whether the character is a letter or an underscore.
The function uses three conditions combined with logical OR operators:

1. `'a' <= ch && ch <= 'z'`: This checks if the character is a lowercase letter between 'a' and 'z' (inclusive).
2. `'A' <= ch && ch <= 'Z'`: This checks if the character is an uppercase letter between 'A' and 'Z' (inclusive).
3. `ch == '_'`: This checks if the character is an underscore.

If any of these conditions are true, the function returns `true`, indicating that the character is considered a letter. Otherwise, it returns `false`.

In the `NextToken()` method, if `isLetter(l.ch)` returns `true`, it means the current character is a letter or an underscore, so we proceed to read the entire identifier using the `readIdentifier()` method:

`readIdentifier()` - 

The `readIdentifier()` method reads characters until it encounters a non-letter character, building the identifier string. After reading the identifier, we use the `LookupIdent()` function to determine if the identifier is a keyword (like `let`, `fn`, `if`, etc.) or a regular identifier.

If `isLetter(l.ch)` returns `false`, we move on to check if the character is a digit using the `isDigit(l.ch)` condition. If it's a digit, we read the entire number using the `readNumber()` method.

If the character is neither a letter nor a digit, we create an illegal token using `newToken(token.ILLEGAL, l.ch)`.

