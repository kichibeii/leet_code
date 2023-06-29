package main

import "fmt"

func main() {
	data := []int{1, 3, 5, 6}
	data2 := 7
	fmt.Println(searchInsert(data, data2))
}

func searchInsert(nums []int, target int) int {
	for key, val := range nums {
		if val >= target {
			return key
		}
	}

	return len(nums)
}
