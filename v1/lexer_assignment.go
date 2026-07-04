package syml

func (l *lexer) isAssignmentIndicator() bool {
	for _, indicator := range l.Settings.AssignmentIndicators {
		if len(indicator) == 0 {
			continue
		}

		if l.pos+len(indicator) > len(l.data) {
			continue
		}

		matched := true
		for i := range indicator {
			if l.data[l.pos+i] != indicator[i] {
				matched = false
				break
			}
		}

		if matched {
			l.emit(TokenAssignment, indicator)
			l.pos += len(indicator)
			l.col += len(indicator)

			l.expectValue = true
			l.lineStart = false
			return l.expectValue
		}
	}

	return false
}
