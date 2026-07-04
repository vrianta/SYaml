package syml

import "fmt"

type TokenType uint8

const (
	TokenEOF TokenType = iota

	TokenIdentifier
	TokenNumber
	TokenString

	TokenColon
	TokenDash
	TokenIndent
	TokenNewLine
)

var TokenTypeString = [...]string{
	TokenEOF:        "EOF",
	TokenIdentifier: "Identifier",
	TokenNumber:     "Number",
	TokenString:     "String",
	TokenColon:      "Colon",
	TokenDash:       "Dash",
	TokenIndent:     "Indent",
	TokenNewLine:    "NewLine",
}

type Token struct {
	Type   TokenType
	Value  []byte
	Line   int
	Column int
}

func (t *Token) String() string {
	return fmt.Sprintf(
		"Token{Type:%v, Value:%q, Line:%d, Column:%d}",
		TokenTypeString[t.Type],
		t.Value,
		t.Line,
		t.Column,
	)
}
