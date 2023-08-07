package parser

import (
	"fmt"
	"strings"
	"unicode"
)

// parseArgs will parse a string that contains quoted strings the same as bash does
// (same as most other *nix shells do). This is secure in the sense that it doesn't do any
// executing or interpeting.
func parseArgs(str string) ([]string, error) {
	if str == "" {
		return []string{}, nil
	}
	var m []string
	var s string

	str = strings.TrimSpace(str) + " "

	nullStr := rune(0)
	lastQuote := nullStr
	isSpace := false
	for i, c := range str {
		switch {
		// If we're ending a quote, break out and skip this character
		case c == lastQuote:
			lastQuote = nullStr

		// If we're in a quote, count this character
		case lastQuote != nullStr:
			s += string(c)

		// If we encounter a quote, enter it and skip this character
		case unicode.In(c, unicode.Quotation_Mark):
			isSpace = false
			lastQuote = c

		// If it's a space, store the string
		case unicode.IsSpace(c):
			if 0 == i || isSpace {
				continue
			}
			isSpace = true
			m = append(m, s)
			s = ""

		default:
			isSpace = false
			s += string(c)
		}

	}

	if lastQuote != nullStr {
		return nil, fmt.Errorf("quotes did not terminate")
	}

	return m, nil
}
