package syml

func (l *lexer) lexIdentifier() {
	start := l.pos

	for !l.eof() {
		c := l.peek()

		if !isLetter(c) && !isDigit(c) && c != '_' {
			break
		}

		l.advance()
	}

	l.emit(TokenIdentifier, l.data[start:l.pos])
}
