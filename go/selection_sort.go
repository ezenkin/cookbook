package main

import "log"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	log.Println(arr)
	selectionSort(arr)
	log.Println(arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
