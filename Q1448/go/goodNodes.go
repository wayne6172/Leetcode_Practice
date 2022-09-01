package goodNodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, maxValue int) int {

	if node == nil {
		return 0
	}

	if maxValue > node.Val {
		return dfs(node.Left, maxValue) + dfs(node.Right, maxValue)
	} else {
		return dfs(node.Left, node.Val) + dfs(node.Right, node.Val) + 1
	}
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}
