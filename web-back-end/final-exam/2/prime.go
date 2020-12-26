package main

import "fmt"

func main() {
	var exit = make(chan bool)
	var primeNumbers []int
	primeNumbers = append(primeNumbers, 2)
	go func() {
		for i := 3; i < 123456; i++ {
			var ok = true
			for j := range primeNumbers {
				if i%primeNumbers[j] == 0 {
					ok = false
				}
			}
			if ok {
				primeNumbers = append(primeNumbers, i)
				fmt.Println(i)
			}
		}
		exit <- true
	}()
	<-exit
}
