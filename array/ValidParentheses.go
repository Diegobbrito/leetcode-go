package array

// https://leetcode.com/problems/valid-parentheses/
// Time Complexity O(n)
// Memory Complexity O(n)
type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return ' '
	}
	temp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return temp
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := Stack{}
	for _, c := range s {
		switch c {
		case '(':
			stack.Push(c)
		case '{':
			stack.Push(c)
		case '[':
			stack.Push(c)
		case ')':
			i := stack.Pop()
			if i != '(' {
				return false
			}
		case '}':
			i := stack.Pop()
			if i != '{' {
				return false
			}
		case ']':
			i := stack.Pop()
			if i != '[' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func isValid2(s string) bool {
	stack := []rune{}
	// Define the mapping of closing to opening brackets
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// Check if the character is a closing bracket
		if open, exists := brackets[char]; exists {
			// If the stack is not empty and the top of the stack matches the corresponding opening bracket
			if len(stack) > 0 && stack[len(stack)-1] == open {
				stack = stack[:len(stack)-1] // Pop the stack
			} else {
				return false // Mismatched or unbalanced bracket
			}
		} else {
			// Push opening brackets onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets are matched
	return len(stack) == 0
}
