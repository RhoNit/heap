package main

import (
	"fmt"
)

type MinHeap struct {
	arr []int
}

func getParent(idx int) int {
	return (idx-1) / 2
}

func getLeftChild(idx int) int {
	return 2*idx + 1
}

func getRightChild(idx int) int {
	return 2*idx + 2
}

func (h *MinHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MinHeap) Insert(elem int) {
	h.arr = append(h.arr, elem)

	// move up the node to the correct position
	h.HeapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) HeapifyUp(idx int) {
	for h.arr[getParent(idx)] > h.arr[idx] {
		h.Swap(idx, getParent(idx))
		idx = getParent(idx)
	}
}

func (h *MinHeap) Remove() int {
	removed := h.arr[0]
	lastIdx := len(h.arr) - 1

	if len(h.arr) == 0 {
		return -1
	}

	// move the last node to the beginning
	h.arr[0] = h.arr[lastIdx]

	// slice off the last position since there is one less node in the queue
	h.arr = h.arr[:lastIdx]

	h.HeapifyDown(0)

	return removed
}

func (h *MinHeap) HeapifyDown(idx int) {
	lastIdx := len(h.arr) - 1
	left, right := getLeftChild(idx), getRightChild(idx)
	childIdxToCompare := 0

	for left <= lastIdx {
		if left == lastIdx {
			childIdxToCompare = left
		} else if h.arr[left] < h.arr[right] {
			childIdxToCompare = left
		} else {
			childIdxToCompare = right
		}

		// swap the smallest child node with its parent node
		if h.arr[childIdxToCompare] < h.arr[idx] {
			h.Swap(childIdxToCompare, idx)
			idx = childIdxToCompare
			left, right = getLeftChild(idx), getRightChild(idx)
		} else {
			return
		}
	}
}

func main() {
	h := &MinHeap{}

	h.Insert(56)
	h.Insert(11)
	h.Insert(9)
	h.Insert(67)
	h.Insert(22)
	h.Insert(37)
	h.Insert(5)
	h.Insert(12)

	fmt.Println("Initial Min Heap : ", h.arr)

	removedElem := h.Remove()
	fmt.Printf("Min elem %d is reomved\n", removedElem)
	fmt.Println("After removal, Min Heap : ", h.arr)

	h.Insert(2)
	fmt.Println("After insertion, Min Heap becomes --> ", h.arr)
}
