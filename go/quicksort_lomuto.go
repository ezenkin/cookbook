package main

import (
	"log"
)

func main() {
	arr := []int{3, 1, 2, 7, 4, 6, 5}
	log.Println("before", arr)
	quickSort2(arr, 0, len(arr)-1)
	log.Println("after", arr)
}

func quickSort2(arr []int, lo, hi int) {
	if lo < hi {
		p := partition2(arr, lo, hi)
		quickSort2(arr, lo, p-1)
		quickSort2(arr, p+1, hi)
	}
}

func partition2(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
