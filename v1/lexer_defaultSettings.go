package syml

func DefaultSettings() Settings {
	return Settings{
		AssignmentIndicators: [][]byte{
			{':'},
			{'='},
		},
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
			{'\r', '\n'},
			{'\n'},
			{';'},
		},

		// Whitespace
		Whitespace: [][]byte{
			[]byte(" "),
			[]byte("\t"),
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
			{
				Start: []byte(`|\n`),
				End:   []byte(`\n\n`),
			},
			{
				Start: []byte(`>\n`),
				End:   []byte(`\n\n`),
			},
		},

		// End of input
		EOF: nil,

		// Indentation
		Indentations: []Delimiter{
			{Start: []byte(" "), End: []byte{'\n', '\n'}},  // 4-space indentation
			{Start: []byte("-"), End: []byte("\n\n")},      // 2-space indentation
			{Start: []byte("\t"), End: []byte{'\n', '\n'}}, // tab indentation
			{Start: []byte("{"), End: []byte("}")},
			{Start: []byte("["), End: []byte("]")},
			{Start: []byte("("), End: []byte(")")},
		},

		PermanentValues: [][]byte{
			[]byte("true"),
			[]byte("false"),
			[]byte("null"),
			[]byte("~"),
		},

		PermanentTypes: [][]byte{
			[]byte("string"),
			[]byte("int"),
			[]byte("float"),
			[]byte("bool"),
			[]byte("byte"),
			[]byte("time"),
		},
	}
}
