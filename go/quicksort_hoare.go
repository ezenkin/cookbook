package main

import (
	"log"
)

func main() {
	arr := []int{3, 1, 2, 7, 4, 6, 5}
	log.Println("before", arr)
	quickSort3(arr, 0, len(arr)-1)
	log.Println("after", arr)
}

func quickSort3(arr []int, lo, hi int) {
	if lo < hi {
		p := partition3(arr, lo, hi)
		quickSort3(arr, lo, p)
		quickSort3(arr, p+1, hi)
	}
}

func partition3(a []int, lo, hi int) int {
	pivot := a[lo+(hi-lo)/2]
	i, j := lo-1, hi+1
	for {
		for {
			i++
			if a[i] < pivot {
				continue
			}
			break
		}
		for {
			j--
			if a[j] > pivot {
				continue
			}
			break
		}
		if i >= j {
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}
