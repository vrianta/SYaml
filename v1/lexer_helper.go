package syml

/*
 * eof reports whether the lexer has reached the end of the input.
 *
 * It returns true when the current position is greater than or equal
 * to the length of the input buffer.
 */
func (l *lexer) eof() bool {
	return l.pos >= len(l.data)
}

/*
 * peek returns the current byte without advancing the lexer.
 *
 * The caller must ensure that eof() is false before calling this method.
 */
func (l *lexer) peek() byte {
	return l.data[l.pos]
}

/*
 * advance moves the lexer forward by one byte.
 *
 * Both the current input position and the column counter are incremented.
 * Line tracking should be handled separately when a newline is encountered.
 */
func (l *lexer) advance() {
	l.pos++
	l.col++
}

/*
 * emit creates a new token and appends it to the lexer's token stream.
 *
 * Parameters:
 *  - t: The type of token being emitted.
 *  - value: The raw byte slice representing the token's value.
 *
 * The emitted token records the current line and column at the time it
 * is created.
 */
func (l *lexer) emit(t TokenType, value []byte) {
	l.tokens = append(l.tokens, Token{
		Type:   t,
		Value:  value,
		Line:   l.line,
		Column: l.col,
	})
}
