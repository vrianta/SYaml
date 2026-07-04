package syml

func (l *lexer) lexNumber() {
	start := l.pos
	hasDot := false

	for !l.eof() {
		c := l.peek()

		if isDigit(c) {
			l.advance()
			continue
		}

		if c == '.' && !hasDot {
			hasDot = true
			l.advance()
			continue
		}

		break
	}

	l.emit(TokenNumber, l.data[start:l.pos])
}