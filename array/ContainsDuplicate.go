package array

// https://leetcode.com/problems/contains-duplicate/
// Time Complexity O(n)
// Space Complexity O(n)
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, num := range nums {
		if set[num] {
			return true
		}
		set[num] = true
	}
	return false
}
