package syml

func Lexer(data []byte, settings Settings) []Token {
	l := lexer{
		data:     data,
		line:     1,
		col:      1,
		Settings: settings,
	}

	return l.lex()
}

type lexer struct {
	data     []byte
	pos      int
	line     int
	col      int
	tokens   []Token
	Settings Settings
}

type Delimiter struct {
	Start []byte
	End   []byte
}

type Settings struct {
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

	// ----------------------------
	// End of input
	// ----------------------------

	EOF []byte

	// ----------------------------
	// Indentation
	// ----------------------------

	TabsAsIndent bool
	IndentWidth  int
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
			continue

		case l.matchWhitespace():
			continue

		case l.Settings.TabsAsIndent && l.peek() == _Tab:
			l.emit(TokenIndent, []byte{_Tab})
			l.advance()

		case l.peek() == _Colon:
			l.emit(TokenColon, []byte{_Colon})
			l.advance()

		case l.peek() == _Dash:
			l.emit(TokenDash, []byte{_Dash})
			l.advance()

		case isDigit(l.peek()):
			l.lexNumber()

		case l.matchStringDelimiter():
			l.lexQuotedString()

		case isLetter(l.peek()):
			l.lexIdentifier()

		default:
			l.lexPlainString()
		}
	}

	l.emit(TokenEOF, nil)
	return l.tokens
}
