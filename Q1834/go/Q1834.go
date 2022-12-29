package Q1834

import (
	"container/heap"
	"sort"
)

type TaskHeap [][]int

func (h TaskHeap) Len() int {
	return len(h)
}

func (h TaskHeap) Less(i, j int) bool {
	if h[i][1] != h[j][1] {
		return h[i][1] < h[j][1]
	}

	return h[i][2] < h[j][2]
}

func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]

	*h = old[0 : n-1]
	return x
}

func getOrder(tasks [][]int) []int {
	ans := []int{}
	taskHeap := &TaskHeap{}
	heap.Init(taskHeap)

	for index := range tasks {
		tasks[index] = append(tasks[index], index)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0] < tasks[j][0]
	})

	nowTime := tasks[0][0]
	for i := 0; i < len(tasks); {
		for ; i < len(tasks) && tasks[i][0] <= nowTime; i++ {
			heap.Push(taskHeap, tasks[i])
		}

		if taskHeap.Len() != 0 {
			nowWork := heap.Pop(taskHeap).([]int)
			ans = append(ans, nowWork[2])
			nowTime = nowTime + nowWork[1]
		} else if i != len(tasks) {
			nowTime = tasks[i][0]
		}
	}

	for taskHeap.Len() > 0 {
		ans = append(ans, heap.Pop(taskHeap).([]int)[2])
	}

	return ans
}
