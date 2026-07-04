package syml

func DefaultSettings() Settings {
	return Settings{
		// Comments
		SingleLineComment: '#',

		MultiLineComments: true,
		MultiLineCommentDelimiter: []Delimiter{
			{
				Start: []byte("/*"),
				End:   []byte("*/"),
			},
		},

		// Line endings
		LineEndings: [][]byte{
			[]byte("\r\n"),
			[]byte("\n"),
			[]byte(";"),
		},

		// Whitespace
		Whitespace: [][]byte{
			[]byte(" "),
		},

		// Strings
		StringDelimiters: [][]byte{
			[]byte(`"`),
			[]byte(`'`),
			[]byte("`"),
		},

		MultiLineStringDelimiter: []Delimiter{
			{
				Start: []byte(`"""`),
				End:   []byte(`"""`),
			},
			{
				Start: []byte(`'''`),
				End:   []byte(`'''`),
			},
		},

		// End of input
		EOF: nil,

		// Indentation
		TabsAsIndent: true,
		IndentWidth:  4,
	}
}
