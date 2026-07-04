package syml

func (l *lexer) lexIndentation() bool {
	// l.expectValue = false
	if !l.lineStart {
		return false
	}

	for _, d := range l.Settings.Indentations {

		// do not match the indent it should return false if the current position does not match the start of the indentation delimiter
		if !l.match(d.Start) {
			// l.emit(TokenIndent, d.Start)
			// l.pos += len(d.Start)
			// l.col += len(d.Start)
			continue
		}

		// Consume the opening delimiter.
		l.pos += len(d.Start)
		l.col += len(d.Start)
		l.emit(TokenIndent, d.Start)

		// Indentation without an explicit end.
		if len(d.End) == 0 {
			l.emit(TokenDedent, d.End)
			l.pos += len(d.Start)
			l.col += len(d.Start)
			return true
		}

		// Lex until the matching end delimiter.
		for !l.eof() {

			// Closing delimiter.
			if l.match(d.End) {
				l.emit(TokenDedent, d.End)
				l.pos += len(d.End)
				l.col += len(d.End)
				return true
			}

			switch {

			case l.matchSingleLineComment():
				continue

			case l.matchMultiLineComment():
				continue

			case l.matchEOF():
				l.emit(TokenEOF, l.Settings.EOF)
				l.pos += len(l.Settings.EOF)
				return true

			case l.matchLineEnding():
				l.updateLineEnding()
				continue
			case l.lexIndentation():
				continue
			case l.matchWhitespace():
				continue

			case l.isAssignmentIndicator():
				l.expectValue = true
				continue

			case l.peek() == _Dash:
				l.emit(TokenDash, []byte{_Dash})
				l.advance()

			case isDigit(l.peek()):
				l.lexNumber()

			case l.matchStringDelimiter():
				l.lexQuotedString()

			case isLetter(l.peek()) && !l.expectValue:
				l.lexIdentifier()

			default:
				l.expectValue = false
				l.lexPlainString()
			}
		}

		// EOF reached before finding the closing delimiter.
		l.pos += l.data_len
		return true
	}

	return false
}
