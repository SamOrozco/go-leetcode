package main

import "os"

func main() {
	if removeOuterParentheses("(()())(())(()(()))") != "()()()()(())" {
		println("Failed")
		os.Exit(1)
	}
}


func removeOuterParentheses(S string) string {
    return ""
}