package Q2530

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]

	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	h := IntHeap(nums)
	heap.Init(&h)

	count := int64(0)
	for ; k > 0; k-- {
		max := heap.Pop(&h).(int)
		count += int64(max)

		max = int(math.Ceil(float64(max) / 3))
		heap.Push(&h, max)
	}

	return count
}
