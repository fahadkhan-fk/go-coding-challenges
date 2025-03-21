// Problem Statement
// You are given an array prices[], where prices[i] represents the stock price on the ith day.

// Your task is to maximize your profit by:

// Buying one stock on a certain day.
// Selling it later on another day after buying it.
// Rules & Constraints
// You must buy before you sell (i.e., cannot sell before buying).
// You cannot buy and sell on the same day.
// If no profit can be made, return 0

// Example 1

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation:

// Buy on Day 2 (price = 1).
// Sell on Day 5 (price = 6).
// Profit = 6 - 1 = 5 (maximum possible profit).

// Example:2

// Input: prices = [7,6,4,3,1]
// Output: 0

package main

import "fmt"

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	profit := getMaximumProfit(prices)
	fmt.Println("Profit: ", profit)

}

func getMaximumProfit(prices []int) int {
	maxProfit := 0
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]

	for _, price := range prices {

		if price <= minPrice {
			minPrice = price
		}

		sell := price - minPrice

		if sell > maxProfit {
			maxProfit = sell
		}

	}
	return maxProfit
}
