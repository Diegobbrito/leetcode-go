package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to act as the start of the merged list
	dummy := &ListNode{}
	current := dummy // Pointer to build the merged list

	// Merge both lists until one is exhausted
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1 // Append list1's node
			list1 = list1.Next   // Move to the next node in list1
		} else {
			current.Next = list2 // Append list2's node
			list2 = list2.Next   // Move to the next node in list2
		}
		current = current.Next // Advance the pointer to the newly added node
	}

	// Append any remaining nodes from list1 or list2
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	// Return the head of the merged list, skipping the dummy node
	return dummy.Next
}
