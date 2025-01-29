// хорошенькая задача, раньше сильно дольше просидел бы

func isValid(s string) bool {
	var (
		st   = &stack{s: make([]rune, 0, len(s))}
		dict = map[rune]rune{
			'(': ')',
			'{': '}',
			'[': ']',
		}
	)

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			st.push(v)
		case ')', '}', ']':
			if dict[st.pop()] != v {
				return false
			}
		}
	}

    if len(st.s) != 0 {
        return false
    }

	return true
}

type stack struct {
	s []rune
}

func (s *stack) push(bracket rune) {
	s.s = append(s.s, bracket)
}

func (s *stack) pop() rune {
    if lastEl := len(s.s) - 1; lastEl >= 0 {
        bracket := s.s[lastEl]
        s.s = s.s[:lastEl]

        return bracket
    }

	return rune('k')
}
