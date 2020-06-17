package main

import (
	"log"
	"sort"
)

func main() {
	tst := []int{1, 3, 5, 7, 9, 11, 13, 15}
	log.Println(binarySearch(tst, 9))
	log.Println(binarySearch(tst, 4))

	log.Println(sort.Search(len(tst), func(i int) bool { return tst[i] == 9 }))
	log.Println(sort.Search(len(tst), func(i int) bool { return tst[i] == 4 }))
}

func binarySearch(arr []int, val int) int {
	low := 0
	high := len(arr)
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] < val {
			low = mid + 1
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			return mid
		}
	}
	return len(arr)
}
