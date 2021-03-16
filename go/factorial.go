package main

import "log"

func main() {
	log.Println(factorial(0))
	log.Println(factorial(1))
	log.Println(factorial(3))
	log.Println(factorial(5))
	log.Println(factorial(10))
	log.Println(factorial(30))
	log.Println(factorial(50))
	//log.Println(factorial(100))
	log.Println(factorial2(0))
	log.Println(factorial2(1))
	log.Println(factorial2(3))
	log.Println(factorial2(5))
	log.Println(factorial2(10))
	log.Println(factorial2(30))
	log.Println(factorial2(50))
}

func factorial(n uint64) uint64 {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}

func factorial2(n uint64) uint64 {
	res := uint64(1)

	for i := n; i > 1; i -= 1 {
		res *= i
	}

	return res
}
