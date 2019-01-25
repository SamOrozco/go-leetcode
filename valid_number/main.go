package main

import (
	"os"
	"strings"
)

func main() {
	if !isNumber("0.1") {
		println("Fail")
		os.Exit(1)
	}

	if !isNumber(" 0.1") {
		println("Fail")
		os.Exit(1)
	}

	if isNumber(" ") {
		println("Fail")
		os.Exit(1)
	}

	if !isNumber(" -90e3") {
		println("Fail")
		os.Exit(1)
	}

	if isNumber(" 99e2.5") {
		println("Fail")
		os.Exit(1)
	}

	if isNumber(".") {
		println("Fail")
		os.Exit(1)
	}

	if isNumber("0e") {
		println("Fail")
		os.Exit(1)
	}

	if isNumber("-.") {
		println("Fail")
		os.Exit(1)
	}

	if !isNumber(".1") {
		println("Fail")
		os.Exit(1)
	}

	println("pass")
}

var dec = '.'
var plus = '+'
var minus = '-'
var e = 'e'

func isNumber(s string) bool {
	validCharacters := getValidChars()
	seenInt := false
	seenE := false
	seenIntAfterE := false
	eInd := -21311 // arbitrary
	seenDec := false
	seenIntAfterDec := false
	s = trim(s)
	for characterIndex, currentCharacter := range s {
		// if we have encountered a bad char
		if ok := validCharacters[currentCharacter]; !ok {
			return false
		}

		if currentCharacter == e {
			if seenE { // no duplicate E's
				return false
			}

			if !seenInt { // no int prefix
				return false
			}
			seenE = true
			eInd = characterIndex
			continue
		}

		if currentCharacter == dec {
			if seenE { // no dec after the e
				return false
			}
			if seenDec { // no duplicate dec
				return false
			}

			seenDec = true
			continue
		}

		if currentCharacter == minus {
			if characterIndex > 0 && (characterIndex != (eInd + 1)) {
				// assuming with my forward thinking that
				// at the zero index and after the e is the only time we should see a minus
				return false
			}
			continue
		}

		if currentCharacter == plus {
			if characterIndex > 0 && (characterIndex != (eInd + 1)) {
				// assuming with my forward thinking that
				// at the zero index and after the e is the only time we should see a minus
				return false
			}
			continue
		}

		if !seenInt {
			seenInt = true
		}

		if seenE && !seenIntAfterE {
			seenIntAfterE = true
		}

		if seenDec && !seenIntAfterDec {
			seenIntAfterDec = true
		}

	}

	if seenE {
		if !seenIntAfterE { // if no ints after e
			return false
		}
	}

	if len(s) == 1 && !seenInt {
		return false
	}

	if !seenInt { // saw no valid ints
		return false
	}

	return true
}

func getValidChars() map[rune]bool {
	return map[rune]bool{
		'0':   true,
		'1':   true,
		'2':   true,
		'3':   true,
		'4':   true,
		'5':   true,
		'6':   true,
		'7':   true,
		'8':   true,
		'9':   true,
		dec:   true,
		e:     true,
		plus:  true,
		minus: true,
	}
}

func trim(s string) string {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return "k"
	}
	return s
}
