package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	testNum := 0
	if testNum++; MaxInt32 != myAtoi("2147483648") {
		println(fmt.Sprintf("Fail %d", testNum))
		os.Exit(1)
	}

	if testNum++; MinInt32 != myAtoi("-6147483648") {
		println(fmt.Sprintf("Fail %d", testNum))
		os.Exit(1)
	}

	if testNum++; MaxInt32 != myAtoi("20000000000000000000") {
		println(fmt.Sprintf("Fail %d", testNum))
		os.Exit(1)
	}

	if testNum++; 12345678 != myAtoi("  0000000000012345678") {
		println(fmt.Sprintf("Fail %d", testNum))
		os.Exit(1)
	}

	if testNum++; -1 != myAtoi("-000000000000001") {
		println(fmt.Sprintf("Fail %d", testNum))
		os.Exit(1)
	}

	if testNum++; -1 != myAtoi("-000000000000001") {
		println(fmt.Sprintf("Fail %d", testNum))
		os.Exit(1)
	}

	if testNum++; 2147483646 != myAtoi("2147483646") {
		println(fmt.Sprintf("Fail %d ", testNum))
		os.Exit(1)
	}

	// 8
	if testNum++; 0 != myAtoi("- 234") {
		println(fmt.Sprintf("Fail %d ", testNum))
		os.Exit(1)
	}

	// 9
	if testNum++; 0 != myAtoi("0-1") {
		println(fmt.Sprintf("Fail %d ", testNum))
		os.Exit(1)
	}

	// 10
	if testNum++; -5 != myAtoi("-5-") {
		println(fmt.Sprintf("Fail %d ", testNum))
		os.Exit(1)
	}

	println("pass")
}

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
	minus    = "-"
	plus     = "+"
)

func myAtoi(str string) int {
	if len(str) < 1 {
		return 0
	}
	// count to keep track of the number of digits
	count := 0
	// container for digits
	values := make([]rune, 0)
	// constant we are going to need
	minus := "-"
	plus := "+"
	lastWasPrefix := false
	// is the number negative or lead with dash
	negative := false

	clean := TrimNum(str)
	for i, char := range clean {
		isNumber := unicode.IsNumber(char)
		isMinus := string(char) == minus
		isPlus := string(char) == plus
		isPrefix := isMinus || isPlus
		// checking for two prefixes following eachother
		if isPrefix && lastWasPrefix {
			return 0
		}
		if isPrefix {
			if i != 0 && i != len(clean) - 1 {
				return 0
			}
			lastWasPrefix = true
		} else {
			lastWasPrefix = false
		}
		curLen := len(values)
		if !isNumber && !isPrefix && curLen == 0 {
			return 0
		}

		if !isNumber && !isPrefix && curLen != 0 {
			break
		}

		if isMinus {
			negative = true
			continue
		}

		if isNumber {
			// add value in order and inc. count
			values = append(values, char)
			count++
		}
	}

	postLen := len(values)

	if postLen == 0 {
		return 0
	}

	multiplier := int32(1)
	ten := int32(10)
	total := int32(0)
	multOver := false
	for i := postLen - 1; i >= 0; i-- {
		diff := MaxInt32 - total

		// if we are overflowing the added values
		// we know the overall value is going to overflow
		if int64(CharToNum(values[i]))*int64(multiplier) > MaxInt32 {
			return overflow(negative)
		}
		added := CharToNum(values[i]) * multiplier
		// added will be zero if we overflow the int
		if added > diff || added < 0 {
			return overflow(negative)
		}
		if added > 0 && multOver {
			return overflow(negative)
		}
		total += added

		// if our multiplier overflows our number will overflow
		if int64(multiplier)*int64(ten) > MaxInt32 && i != 0 {
			multOver = true
		}
		multiplier *= ten
	}

	if negative {
		return int(total * int32(-1))
	} else {
		return int(total)
	}
}

func overflow(neg bool) int {
	if neg {
		return MinInt32
	} else {
		return MaxInt32
	}
}

/**
a rune is a uni code char so we need to conver from unicode int to digit
 */
func CharToNum(r rune) (int32) {
	for i := 48; i <= 57; i++ {
		if int(r) == i {
			return int32(r) - 48
		}
	}
	return -1
}

func TrimNum(s string) string {
	runes := []rune(s)
	neg := false
	pos := false
	if string(runes[0]) == minus {
		s = s[1:]
		neg = true
	}

	if string(runes[0]) == plus {
		s = s[1:]
		pos = true
	}
	//trim all space from left and right of string
	s = strings.TrimRightFunc(s, unicode.IsSpace)
	// we care about trimming left because `- 234` is not a valid number I guess previous alg
	// turned that into -234 which I thought was right
	preTrimSpaceLen := len(s)
	s = strings.TrimLeftFunc(s, unicode.IsSpace)
	if neg || pos {
		if preTrimSpaceLen > len(s) { // if the length before trimming and after are different in case of pos and neg sign
			return "0"
		}
	}

	//s = strings.TrimLeft(s, "0")

	if neg {
		s = "-" + s
	} else if pos {
		s = "+" + s
	}
	return s
}
