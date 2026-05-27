package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    leftInverted := invertTree(root.Left)
    rightInverted := invertTree(root.Right)
    root.Left = rightInverted
    root.Right = leftInverted
    return root
}
