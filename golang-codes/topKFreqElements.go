// https://leetcode.com/problems/top-k-frequent-elements/

type Pair struct {
    number int
    frequency int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i].frequency < h[j].frequency
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    
    // store the element before popping out
    x := old[n-1]

    // slice off the last element as the heap now has one element less
    *h = old[:n-1]

    return x
}

func topKFrequent(nums []int, k int) []int {
    mp := make(map[int]int)

    for _, val := range nums {
        mp[val]++
    }

    h := &MinHeap{}
    heap.Init(h)

    for elem, count := range mp {
        pair := Pair {
            number: elem,
            frequency: count,
        }

        heap.Push(h, pair)

        if h.Len() > k {
            heap.Pop(h)
        }
    }

    arr := make([]int, k)

    for i := 0; i < k; i++ {
        arr[i] = heap.Pop(h).(Pair).number
    }

    return arr
}