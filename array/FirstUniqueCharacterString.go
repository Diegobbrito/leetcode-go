package array

// https://leetcode.com/problems/first-unique-character-in-a-string/
// Time Complexity: O(n)
// Space Complexity: O(1)
func firstUniqChar(s string) int {
	chars := [26]int{}

	for _, c := range s {
		chars[c-'a']++
	}

	for i, c := range s {
		if chars[c-'a'] == 1 {
			return i
		}
	}

	return -1
}
