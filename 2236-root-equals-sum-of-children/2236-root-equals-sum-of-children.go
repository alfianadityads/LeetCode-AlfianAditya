/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
    // Check if the root is nil
  if root == nil {
    return false
  }

  // Get the values of the left and right children
  leftValue := root.Left.Val
  rightValue := root.Right.Val

  // Check if the sum of the children equals the value of the root
  return leftValue + rightValue == root.Val
}