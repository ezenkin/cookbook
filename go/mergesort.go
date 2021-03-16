package main

import "fmt"

func mergeSort(arr []element) []element {
	if len(arr) < 2 {
		return arr
	}
	m := len(arr) / 2
	a := mergeSort(arr[:m])
	b := mergeSort(arr[m:])
	return merge(a, b)
}

func merge(arr1, arr2 []element) []element {
	res := make([]element, 0, len(arr1)+len(arr2))
	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] < arr2[0] {
			res = append(res, arr1[0])
			arr1 = arr1[1:]
		} else {
			res = append(res, arr2[0])
			arr2 = arr2[1:]
		}
	}

	if len(arr1) > 0 {
		res = append(res, arr1...)
	} else if len(arr2) > 0 {
		res = append(res, arr2...)
	}

	return res
}

func main() {
	a := []element{3, 2, 1}
	fmt.Println(a)
	b := mergeSort(a)
	fmt.Println(b)
}
