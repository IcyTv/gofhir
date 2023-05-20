package expression

import (
	"strconv"
	"strings"

	"github.com/gofhir/fhirpath/interpreter"
)

const stringDelimiterChar = '\''

type StringLiteral struct {
	node interpreter.StringAccessor
}

func NewRawStringLiteral(value string) *StringLiteral {
	return &StringLiteral{interpreter.NewString(value)}
}

func ParseStringLiteral(value string) *StringLiteral {
	return &StringLiteral{interpreter.NewString(parseStringLiteral(value, stringDelimiterChar))}
}

// TODO realistically this should be the responsibility of the interpreter
// TODO there's actually a restriction on what characters can and should be escaped, which we're not enforcing
func parseStringLiteral(value string, delimiter byte) string {
	l := len(value)
	if l < 2 || value[0] != delimiter || value[l-1] != delimiter {
		return value
	}

	var b strings.Builder
	b.Grow(1)

	var esc bool
	var unicode bool
	var unicodeValue strings.Builder
	for _, char := range value[1 : l-1] {
		if unicode {
			unicodeValue.WriteRune(char)
			if unicodeValue.Len() == 4 {
				if r, err := strconv.ParseInt(unicodeValue.String(), 16, 32); err != nil {
					b.WriteRune('u')
					b.WriteString(unicodeValue.String())
				} else {
					b.WriteRune(int32(r))
				}
				unicode = false
			}
		} else if esc {
			switch char {
			case 'r':
				b.WriteRune('\r')
			case 'n':
				b.WriteRune('\n')
			case 't':
				b.WriteRune('\t')
			case 'f':
				b.WriteRune('\f')
			case 'u':
				unicode = true
				unicodeValue.Reset()
				unicodeValue.Grow(4)
			default:
				b.WriteRune(char)
			}
			esc = false
		} else if char == '\\' {
			esc = true
		} else {
			b.WriteRune(char)
		}
	}

	if unicode {
		b.WriteRune('u')
		b.WriteString(unicodeValue.String())
	}

	return b.String()
}

func (e *StringLiteral) Evaluate(interpreter.ContextAccessor, interface{}, interpreter.Looper) (interface{}, error) {
	return e.node, nil
}
