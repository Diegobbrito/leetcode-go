package array

// https://leetcode.com/problems/valid-anagram/
// Time Complexity O(n + m)
// Memory Complexity O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, c := range s {
		m[c] += 1
	}
	for _, c := range t {
		m[c] -= 1
	}
	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
