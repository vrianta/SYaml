package syml

func (l *lexer) lexIdentifier() {
	start := l.pos

	for !l.eof() {

		// Stop at line endings.
		if l.matchLineEnding() {
			break
		}

		// Stop before comments.
		if l.matchSingleLineComment() || l.matchMultiLineComment() {
			break
		}

		c := l.peek()

		// Stop at whitespace.
		if c == _Space || c == _Tab {
			break
		}

		// Valid identifier characters.
		if !isLetter(c) && !isDigit(c) && c != '_' {
			break
		}

		l.advance()
	}

	if start != l.pos {
		identifier := l.data[start:l.pos]
		l.lineStart = false
		if l.matchPermanentType(identifier) {
			l.emit(TokenPermanentType, identifier)
		} else {
			l.emit(TokenIdentifier, l.data[start:l.pos])
		}
	}
	l.updateLineEnding()
}
