package syml

func (l *lexer) lexString() {
	start := l.pos

	for !l.eof() {
		c := l.peek()

		if c == '\n' {
			break
		}

		l.advance()
	}

	end := l.pos

	for end > start && l.data[end-1] == _Space {
		end--
	}

	l.emit(TokenString, l.data[start:end])
}

func (l *lexer) match(b []byte) bool {
	if len(b) == 0 || l.pos+len(b) > len(l.data) {
		return false
	}

	for i := range b {
		if l.data[l.pos+i] != b[i] {
			return false
		}
	}

	return true
}

func (l *lexer) matchWhitespace() bool {
	for _, ws := range l.Settings.Whitespace {
		if l.match(ws) {
			l.pos += len(ws)
			l.col += len(ws)
			return true
		}
	}

	return false
}

func (l *lexer) matchStringDelimiter() bool {
	for _, d := range l.Settings.StringDelimiters {
		if l.match(d) {
			return true
		}
	}

	return false
}

func (l *lexer) lexQuotedString() {
	var delimiter []byte

	// Find the opening delimiter.
	for _, d := range l.Settings.StringDelimiters {
		if l.match(d) {
			delimiter = d
			break
		}
	}

	if delimiter == nil {
		return
	}

	// Skip opening delimiter.
	l.pos += len(delimiter)
	l.col += len(delimiter)

	valueStart := l.pos

	for !l.eof() {
		// Handle escaped characters.
		if l.peek() == '\\' {
			l.advance()

			if !l.eof() {
				l.advance()
			}

			continue
		}

		// Closing delimiter.
		if l.match(delimiter) {
			l.emit(TokenString, l.data[valueStart:l.pos])

			l.pos += len(delimiter)
			l.col += len(delimiter)

			return
		}

		// Unterminated single-line string.
		if l.matchLineEnding() {
			break
		}

		l.advance()
	}

	// Emit what we have. You may later replace this with an error.
	l.emit(TokenString, l.data[valueStart:l.pos])
	l.updateLineEnding()
}

func (l *lexer) lexPlainString() {
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

		switch l.peek() {
		case _Space, _Tab:
			goto done
		}

		l.advance()
	}

done:
	if start != l.pos {
		l.emit(TokenString, l.data[start:l.pos])
	}
	l.updateLineEnding()
}

func (l *lexer) lexMultiLineString() bool {

	// Find the opening delimiter.
	for _, d := range l.Settings.MultiLineStringDelimiter {
		if !l.match(d.Start) {
			break
		}

		start := d.Start

		// Skip opening delimiter.
		l.pos += len(start)
		l.col += len(start)

		valueStart := l.pos

		for !l.eof() {
			// Handle escaped characters.
			if l.peek() == '\\' {
				l.advance()

				if !l.eof() {
					l.advance()
				}

				continue
			}

			if l.matchWhitespace() {
				continue
			}

			// Closing delimiter.
			if l.match(d.End) {
				l.emit(TokenString, l.data[valueStart:l.pos])

				l.pos += len(d.End)
				l.col += len(d.End)

				return true
			}

			l.advance()
		}

		// Emit what we have. You may later replace this with an error.
		l.emit(TokenString, l.data[valueStart:l.pos])
	}

	return false
}
