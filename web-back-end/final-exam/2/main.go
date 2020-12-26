package main

import (
	"fmt"
	"math"
)

func main() {
	var exit = make(chan bool)
	go func() {
		for i := 2; i < 123456; i++ {
			var factors []int
			for j := 1; j < int(math.Sqrt(float64(i)))+1; j++ {
				if i%j == 0 {
					factors = append(factors, j)
					factors = append(factors, i/j)
				}
			}
			var sum = 0
			for j := range factors {
				sum += factors[j]
			}
			if sum-i == i {
				fmt.Println(i)
			}
		}
		exit <- true
	}()
	<-exit
}
