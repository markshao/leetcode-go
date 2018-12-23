package easy

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


 // https://leetcode.com/problems/middle-of-the-linked-list/submissions/

func middleNode(head *ListNode) *ListNode {
	one , two := head, head

	for one != nil && two != nil {
		if one != nil && two.Next != nil { // 判断2倍速的节点是否已经到NodeList的尾部
			one = one.Next
		}

		for i:=0;i<2 && two !=  nil;i++ {
			two = two.Next
		}
	}

	return one
}
