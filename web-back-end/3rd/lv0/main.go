package main

import (
	"fmt"
	"sync"
)

var (
	myres = make(map[int]int, 20)
	ch    = make(chan int)
	mu    sync.Mutex
)

func factorial() {
	var n int
	n = <-ch
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	mu.Lock()
	myres[n] = res
	mu.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial()
		ch <- i
	}
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}
