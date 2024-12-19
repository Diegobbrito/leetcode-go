package array

// https://leetcode.com/problems/valid-palindrome/
// Time Complexity O(n)
// Space Complexity O(1)
func maxProfit(prices []int) int {

	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
