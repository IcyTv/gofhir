package interpreter

import (
	"strconv"
	"strings"
	"unicode"
)

type normalizedStringStream struct {
	value []rune
	pos   int
	size  int
}

func (s *normalizedStringStream) next() (rune, bool) {
	begin := s.pos == 0
	ws := false
	for s.pos < s.size {
		c := s.value[s.pos]
		if c == ' ' || c == '\r' || c == '\n' || c == '\t' {
			s.pos = s.pos + 1
			ws = true
		} else if ws && !begin {
			return ' ', true
		} else {
			s.pos = s.pos + 1
			return unicode.ToLower(c), true
		}
	}
	return 0, false
}

func NormalizedStringEqual(value1 string, value2 string) bool {
	if value1 == value2 {
		return true
	}
	if len(value1) == 0 || len(value2) == 0 {
		return false
	}

	r1 := []rune(value1)
	s1 := normalizedStringStream{value: r1, size: len(r1)}
	r2 := []rune(value2)
	s2 := normalizedStringStream{value: r2, size: len(r2)}

	for c1, ok1 := s1.next(); ok1; c1, ok1 = s1.next() {
		c2, ok2 := s2.next()
		if !ok2 || c1 != c2 {
			return false
		}
	}
	_, ok := s2.next()
	return !ok
}

func writeStringBuilderInt(b *strings.Builder, value int, digits int) {
	formatted := strconv.FormatInt(int64(value), 10)
	l := len(formatted)
	for i := l; i < digits; i++ {
		b.WriteByte('0')
	}
	b.WriteString(formatted)
}
