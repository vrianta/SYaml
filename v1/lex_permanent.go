package syml

import "bytes"

func (l *lexer) matchPermanentValue() bool {
	for _, value := range l.Settings.PermanentValues {
		if !l.match(value) {
			continue
		}

		// Ensure it isn't a prefix of another identifier.
		if !l.eof() {
			c := l.peek()

			if isLetter(c) || isDigit(c) || c == '_' {
				return false
			}
		}

		l.emit(TokenPermanentValue, value)
		return true
	}

	return false
}

func (l *lexer) matchPermanentType(value []byte) bool {
	for _, t := range l.Settings.PermanentTypes {
		if bytes.Equal(value, t) {
			return true
		}
	}

	return false
}
