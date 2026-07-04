package syml

func Lexer(data []byte, settings Settings) []Token {
	l := lexer{
		data:        data,
		data_len:    len(data),
		line:        1,
		col:         1,
		lineStart:   true,
		indentStack: []int{0},
		Settings:    settings,
	}

	return l.lex()
}

type lexer struct {
	data     []byte
	data_len int
	pos      int
	line     int
	col      int

	expectValue bool
	ending      []byte

	tokens []Token

	indentStack []int
	lineStart   bool

	Settings Settings
}

type Delimiter struct {
	Start []byte
	End   []byte
}

type Settings struct {
	AssignmentIndicators [][]byte

	// ----------------------------
	// Comments
	// ----------------------------

	SingleLineComment byte

	MultiLineComments         bool
	MultiLineCommentDelimiter []Delimiter

	// ----------------------------
	// Line endings
	// ----------------------------

	LineEndings [][]byte // e.g. {"\n"}, {"\r\n"}

	// ----------------------------
	// Whitespace
	// ----------------------------

	Whitespace [][]byte // default: {" "}

	// ----------------------------
	// Strings
	// ----------------------------

	// Single-line string delimiters.
	// Examples: {"\""}, {"'"}, {"`"}
	StringDelimiters [][]byte

	// Multi-line string delimiters.
	// Examples: {`"""`}, {`'''`}
	MultiLineStringDelimiter []Delimiter

	Indentations []Delimiter

	PermanentValues [][]byte
	PermanentTypes  [][]byte

	// ----------------------------
	// End of input
	// ----------------------------

	EOF []byte
}

func (l *lexer) lex() []Token {

	for !l.eof() {

		switch {

		case l.matchSingleLineComment():
			continue

		case l.matchMultiLineComment():
			continue

		case l.matchEOF():
			l.emit(TokenEOF, l.Settings.EOF)
			l.pos += len(l.Settings.EOF)
			return l.tokens

		case l.matchLineEnding():
			l.updateLineEnding()
			continue

		case l.matchWhitespace():
			continue

		case l.isAssignmentIndicator():
			continue
		}

		if l.expectValue {

			switch {

			case l.matchPermanentValue():
				continue

			case isDigit(l.peek()):
				l.lexNumber()
				continue

			case l.lexMultiLineString():
				continue

			case l.lexIndentation():
				continue

			case l.matchStringDelimiter():
				l.lexQuotedString()
				continue

			default:
				l.expectValue = false
				l.lexPlainString()
				continue
			}

		}

		switch {

		case isLetter(l.peek()):
			l.lexIdentifier()
			continue

		default:
			l.lexPlainString()
			continue
		}
	}

	l.emit(TokenEOF, nil)
	return l.tokens
}
