package main

import "fmt"

func main() {
	data := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(data))
}

func maxProfit(prices []int) int {
	buy := prices[0]
	sell := prices[0]
	totalNumber := 0

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			continue
		}

		if prices[i] < buy {
			fmt.Println("kesini 1", prices[i])
			buy = prices[i]
			sell = prices[i]
		} else if prices[i] > sell {
			sell = prices[i]
			if (sell - buy) > totalNumber {
				fmt.Println("kesini 3", prices[i])
				totalNumber = sell - buy
			}
			fmt.Println("kesini 2", prices[i], totalNumber)

		}
	}

	return totalNumber
}
