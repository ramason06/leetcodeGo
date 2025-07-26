package medium

import "container/heap"

type Edge struct {
	node int
	prob float64
}

type MaxHeap []Edge

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].prob > h[j].prob } // 최대 힙이므로 큰 확률이 앞에 오게
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	l := len(old)
	cur := old[l-1]
	*h = old[0 : l-1]
	return cur
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]Edge, n)

	//graph 초기화
	for i, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], Edge{node: edge[1], prob: succProb[i]})
		graph[edge[1]] = append(graph[edge[1]], Edge{node: edge[0], prob: succProb[i]})
	}

	prob := make([]float64, n)
	// 초기화
	for i, _ := range prob {
		prob[i] = 0.0
	}

	prob[start] = 1.0

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Edge{node: start, prob: 1.0})

	for maxHeap.Len() > 0 {
		cur := heap.Pop(maxHeap).(Edge)
		curNode := cur.node
		curProb := cur.prob

		if curNode == end {
			return curProb
		}

		for _, next := range graph[curNode] {
			nextProb := next.prob * curProb
			if nextProb > prob[next.node] {
				prob[next.node] = nextProb
				heap.Push(maxHeap, Edge{
					node: next.node,
					prob: nextProb,
				})
			}
		}
	}

	return 0.0
}
