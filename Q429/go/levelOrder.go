package levelOrder

type Node struct {
	Val      int
	Children []*Node
}

type levelNode struct {
	Node
	Level int
}

func levelOrder(root *Node) [][]int {

	queue := make([]levelNode, 0)
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	queue = append(queue, levelNode{
		Level: 0,
		Node:  *root,
	})

	for len(queue) != 0 {
		nowNode := queue[0]

		if len(ans) == nowNode.Level {
			ans = append(ans, []int{nowNode.Val})
		} else {
			ans[nowNode.Level] = append(ans[nowNode.Level], nowNode.Val)
		}

		for i := 0; i < len(nowNode.Children); i++ {
			queue = append(queue, levelNode{
				Node:  *nowNode.Children[i],
				Level: nowNode.Level + 1,
			})
		}

		queue = queue[1:]
	}
	return ans
}
