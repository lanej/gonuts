package median

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int64

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Root() int64        { return h[0] }
func (h MinHeap) Root() int64        { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func nextNumber(r *bufio.Reader) int64 {
	num_s, _ := r.ReadString('\n')
	num, err := strconv.ParseInt(strings.Trim(num_s, "\n"), 10, 64)

	if nil != err {
		panic(err)
	}

	return num
}

func solve(r *os.File, w *os.File) {
	reader := bufio.NewReader(r)

	count := int(nextNumber(reader))
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}

	heap.Init(minHeap)
	heap.Init(maxHeap)

	var median float64

	for i := 0; i < count; i++ {
		num := int64(nextNumber(reader))

		// Push on correct heap
		if 2 == i && maxHeap.Root() > minHeap.Root() {
			min := heap.Pop(minHeap)
			max := heap.Pop(maxHeap)
			heap.Push(minHeap, max)
			heap.Push(maxHeap, min)
		}
		if minHeap.Len() == 0 {
			heap.Push(minHeap, num)
		} else if (maxHeap.Len() == 0) || (num < maxHeap.Root()) {
			heap.Push(maxHeap, num)
		} else {
			heap.Push(minHeap, num)
		}

		// Rebalance
		if minHeap.Len()-maxHeap.Len() > 1 {
			heap.Push(maxHeap, heap.Pop(minHeap))
		} else if maxHeap.Len()-maxHeap.Len() > 1 {
			heap.Push(minHeap, heap.Pop(maxHeap))
		}

		if maxHeap.Len() == minHeap.Len() {
			median = (float64(maxHeap.Root()) + float64(minHeap.Root())) / 2
		} else if maxHeap.Len() > minHeap.Len() {
			median = float64(maxHeap.Root())
		} else {
			median = float64(minHeap.Root())
		}

		w.WriteString(strconv.FormatFloat(median, 'f', 1, 32))

		if (i + 1) < count {
			w.WriteString("\n")
		}
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
