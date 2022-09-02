package averageOfLevels

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	LevelSum   []int64
	LevelNodes []int
}

func NewSolution() Solution {
	return Solution{
		LevelSum:   make([]int64, 10000),
		LevelNodes: make([]int, 10000),
	}
}

func (solution *Solution) dfs(root *TreeNode, level int) {
	if root != nil {
		solution.LevelSum[level] += int64(root.Val)
		solution.LevelNodes[level] += 1

		solution.dfs(root.Left, level+1)
		solution.dfs(root.Right, level+1)
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func averageOfLevels(root *TreeNode) []float64 {
	solution := NewSolution()
	solution.dfs(root, 0)

	maxLevel := 0
	for i := 0; i < 10000; i++ {
		if solution.LevelNodes[i] == 0 {
			maxLevel = i
			break
		}
	}

	ans := make([]float64, maxLevel)
	for i := 0; i < maxLevel; i++ {
		ans[i] = float64(solution.LevelSum[i]) / float64(solution.LevelNodes[i])
		ans[i] = toFixed(ans[i], 5)
	}

	return ans
}
