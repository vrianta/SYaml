package syml

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}

func (l *lexer) matchLineEnding() bool {
	for _, ending := range l.Settings.LineEndings {

		if len(ending) == 0 {
			continue
		}

		if l.pos+len(ending) > len(l.data) {
			continue
		}

		match := true
		for i := range ending {
			if l.data[l.pos+i] != ending[i] {
				match = false
				break
			}
		}

		if match {
			l.emit(TokenNewLine, ending)

			l.pos += len(ending)
			l.line++
			l.col = 1

			return true
		}
	}

	return false
}

func (l *lexer) matchEOF() bool {

	if len(l.Settings.EOF) == 0 {
		return false
	}

	if l.pos+len(l.Settings.EOF) > len(l.data) {
		return false
	}

	for i := range l.Settings.EOF {
		if l.data[l.pos+i] != l.Settings.EOF[i] {
			return false
		}
	}

	return true
}

func (l *lexer) matchSingleLineComment() bool {

	if l.Settings.SingleLineComment == 0 {
		return false
	}

	if l.peek() != l.Settings.SingleLineComment {
		return false
	}

	for !l.eof() {

		if l.matchLineEnding() {
			return true
		}

		l.advance()
	}

	return true
}

func (l *lexer) matchMultiLineComment() bool {
	if !l.Settings.MultiLineComments {
		return false
	}

	for _, d := range l.Settings.MultiLineCommentDelimiter {

		if !l.match(d.Start) {
			continue
		}

		// Skip the opening delimiter.
		l.pos += len(d.Start)
		l.col += len(d.Start)

		for !l.eof() {

			if l.match(d.End) {
				l.pos += len(d.End)
				l.col += len(d.End)
				return true
			}

			if l.matchLineEnding() {
				continue
			}

			l.advance()
		}

		// EOF before finding the closing delimiter.
		return true
	}

	return false
}
