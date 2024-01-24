package playwright

import (
	"regexp"
	"strings"
)

var escapedChars = map[rune]bool{
	'$':  true,
	'^':  true,
	'+':  true,
	'.':  true,
	'*':  true,
	'(':  true,
	')':  true,
	'|':  true,
	'\\': true,
	'?':  true,
	'{':  true,
	'}':  true,
	'[':  true,
	']':  true,
}

func globMustToRegex(glob string) *regexp.Regexp {
	tokens := []string{"^"}
	inGroup := false

	for i := 0; i < len(glob); i++ {
		c := rune(glob[i])
		if c == '\\' && i+1 < len(glob) {
			char := rune(glob[i+1])
			if _, ok := escapedChars[char]; ok {
				tokens = append(tokens, "\\"+string(char))
			} else {
				tokens = append(tokens, string(char))
			}
			i++
		} else if c == '*' {
			beforeDeep := rune(0)
			if i > 0 {
				beforeDeep = rune(glob[i-1])
			}
			starCount := 1
			for i+1 < len(glob) && glob[i+1] == '*' {
				starCount++
				i++
			}
			afterDeep := rune(0)
			if i+1 < len(glob) {
				afterDeep = rune(glob[i+1])
			}
			isDeep := starCount > 1 && (beforeDeep == '/' || beforeDeep == 0) && (afterDeep == '/' || afterDeep == 0)
			if isDeep {
				tokens = append(tokens, "((?:[^/]*(?:/|$))*)")
				i++
			} else {
				tokens = append(tokens, "([^/]*)")
			}
		} else {
			switch c {
			case '?':
				tokens = append(tokens, ".")
			case '[':
				tokens = append(tokens, "[")
			case ']':
				tokens = append(tokens, "]")
			case '{':
				inGroup = true
				tokens = append(tokens, "(")
			case '}':
				inGroup = false
				tokens = append(tokens, ")")
			case ',':
				if inGroup {
					tokens = append(tokens, "|")
				} else {
					tokens = append(tokens, string(c))
				}
			default:
				if _, ok := escapedChars[c]; ok {
					tokens = append(tokens, "\\"+string(c))
				} else {
					tokens = append(tokens, string(c))
				}
			}
		}
	}

	tokens = append(tokens, "$")
	return regexp.MustCompile(strings.Join(tokens, ""))
}
