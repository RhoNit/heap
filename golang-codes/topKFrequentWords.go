// https://leetcode.com/problems/top-k-frequent-words

import (
    "container/heap"
)

type Pair struct {
    Word string
    Frequency int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    if h[i].Frequency == h[j].Frequency {
        return h[i].Word < h[j].Word
    }
    
    return h[i].Frequency > h[j].Frequency
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)

    x := old[n-1]

    // slice off the element
    *h = old[:n-1]

    return x
}


func topKFrequent(words []string, k int) []string {
    mp := make(map[string]int)
    for _, val := range words {
        mp[val]++
    }

    h := &MaxHeap{}
    for key, count := range mp {
        pair := Pair {
            Word: key,
            Frequency: count, 
        }

        heap.Push(h, pair)
    }

    var res []string
    for i := 0; i < k; i++ {
        res = append(res, heap.Pop(h).(Pair).Word)
    }

    return res
}