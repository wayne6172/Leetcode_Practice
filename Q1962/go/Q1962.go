package Q1962

import (
	"container/heap"
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

func minStoneSum(piles []int, k int) int {
	h := IntHeap(piles)
	heap.Init(&h)

	for ; k > 0; k-- {
		h[0] -= h[0] / 2
		heap.Fix(&h, 0)
	}

	ans := 0
	for _, value := range piles {
		ans += value
	}

	return ans
}
