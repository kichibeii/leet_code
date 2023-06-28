package main

import "fmt"

func main() {
	data := "()"
	fmt.Println(isValid(data))
}

func isValid(s string) bool {
	success := true
	lengthArray := 0
	mapData := make(map[rune]rune)
	mapData['('] = ')'
	mapData['['] = ']'
	mapData['{'] = '}'
	mapOrder := make(map[int]rune)
	for _, char := range s {
		fmt.Println(string(char), lengthArray, mapOrder)
		if char == '(' || char == '[' || char == '{' {
			mapOrder[lengthArray] = char
			lengthArray++
		} else {
			if char == mapData[mapOrder[lengthArray-1]] {
				lengthArray--
			} else {
				return false
			}
		}
	}

	return success
}
