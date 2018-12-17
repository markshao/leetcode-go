package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	var t1Val, t2Val int
	// set the value
	if t1 != nil {
		t1Val = t1.Val
	}

	if t2 != nil {
		t2Val = t2.Val
	}

	newNode := &TreeNode{}
	newNode.Val = t1Val + t2Val

	newLeft := mergeTrees(t1.Left, t2.Left)
	newRight := mergeTrees(t1.Right, t2.Right)
	newNode.Left = newLeft
	newNode.Right = newRight

	return newNode
}

func main() {

}
