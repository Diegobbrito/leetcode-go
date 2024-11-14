package array

// https://leetcode.com/problems/two-sum/
// Time Complexity O(n)
// Space Complexity O(n)

func twoSum(nums []int, target int) []int {
	// Initialize the map
	m := make(map[int]int)

	for i, num := range nums {
		// Check if the complement (target - num) exists in the map
		if j, found := m[target-num]; found {
			return []int{j, i} // Return the indices as a slice
		}
		// Store the index of the current number
		m[num] = i
	}

	// Return an empty slice if no solution is found
	return []int{}
}
