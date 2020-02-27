/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftMax := maxDepth(root.Left)
    rightMax := maxDepth(root.Right)
    return int(math.Max(float64(leftMax), float64(rightMax)) + 1)
}