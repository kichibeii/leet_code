package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := " "
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	s = strings.ToLower(s)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "")

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
