package wildcard_matching

import "strings"

func main() {
	println(isMatch("abcdefg", "*a*g"))
}

func isMatch(s string, p string) bool {
	if !containsSpecialCharacters(s) { // exact match
		return s == p
	}
	values := []rune(s)
	for i, v := range p {
		if v != '*' && v != '?'
	}
}

func containsSpecialCharacters(val string) bool {
	return strings.Contains(val, "*") ||
		strings.Contains(val, "?")
}
