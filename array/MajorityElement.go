package array

// https://leetcode.com/problems/majority-element/

// Time Complexity O(n)
// Space Complexity O(1)
func majorityElement(nums []int) int {

	majority := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if num == majority {
			count += 1
		} else {
			count -= 1
		}
	}
	return majority
}
