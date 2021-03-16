package main

import "log"

func main() {
	arr := []int{4, 2, 5, 1, 3}
	log.Println("Before: ", arr)
	insertionSort(arr)
	log.Println("After: ", arr)
}

func insertionSort(arr []int) {
	var i, j, k int
	for i = 1; i < len(arr); i++ {
		k = arr[i]
		for j = i - 1; (j >= 0) && (arr[j] > k); j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = k
	}
}
