// https://leetcode.com/problems/kth-largest-element-in-an-array

import (
    "container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
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

func findKthLargest(nums []int, k int) int {
    h := &MinHeap{}

    for _, val := range nums {
        heap.Push(h, val)

        for h.Len() > k {
            heap.Pop(h)
        }
    }

    //return heap.Pop(h).(int)
    return (*h)[0]
}