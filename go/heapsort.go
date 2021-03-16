package main

import (
	"fmt"
)

type Heap struct {
	arr []element
}

func (h *Heap) make(input []element) {
	h.arr = append(h.arr, 0)
	for _, a := range input {
		h.insert(a)
	}
}

func (h *Heap) insert(a element) {
	h.arr = append(h.arr, a)
	h.bubbleUp(len(h.arr) - 1)
}

func (h *Heap) bubbleUp(idx int) {
	parentIdx := h.parent(idx)

	if parentIdx == -1 {
		return
	}

	if h.arr[parentIdx] > h.arr[idx] {
		h.arr[parentIdx], h.arr[idx] = h.arr[idx], h.arr[parentIdx]
		h.bubbleUp(parentIdx)
	}
}

func (h *Heap) parent(idx int) int {
	if idx > 1 {
		return idx / 2
	}

	return -1
}

func (h *Heap) child(idx int) int {
	return idx * 2
}

func (h *Heap) extractMin() (el element, err error) {
	if len(h.arr) < 2 {
		return el, fmt.Errorf("empty heap")
	}

	min := h.arr[1]
	h.arr[1] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.bubbleDown(1)
	return min, nil
}

func (h *Heap) bubbleDown(idx int) {
	var minElemIdx int
	childIdx1, childIdx2 := h.child(idx), h.child(idx)+1
	if childIdx1 > (len(h.arr) - 1) {
		return
	} else if childIdx2 > (len(h.arr) - 1) {
		minElemIdx = childIdx1
	} else if h.arr[childIdx1] < h.arr[childIdx2] {
		minElemIdx = childIdx1
	} else {
		minElemIdx = childIdx2
	}

	if h.arr[minElemIdx] < h.arr[idx] {
		h.arr[minElemIdx], h.arr[idx] = h.arr[idx], h.arr[minElemIdx]
		h.bubbleDown(minElemIdx)
	}
}

func heapSort(arr []element) {
	var h Heap
	h.make(arr[:])
	for i := range arr {
		arr[i], _ = h.extractMin()
	}
}

func main() {
	a := []element{5, 4, 2, 3, 1}
	fmt.Println("before sorting:", a)
	heapSort(a)
	fmt.Println("after sorting:", a)
}
