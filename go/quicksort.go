package main

import (
	"log"
)

func main() {
	arr := []int{3, 1, 2, 7, 4, 6, 5}
	log.Println("before", arr)
	quickSort(arr)
	log.Println("after", arr)
}

func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	lo, hi := 0, len(a)-1

	for i := 0; i < len(a); i++ {
		if a[i] < a[hi] {
			a[lo], a[i] = a[i], a[lo]
			lo++
		}
	}
	a[lo], a[hi] = a[hi], a[lo]
	quickSort(a[:lo])
	quickSort(a[lo+1:])
}
