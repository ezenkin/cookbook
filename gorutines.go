package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second * time.Duration(i))
			fmt.Println("i =", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
