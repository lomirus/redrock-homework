package main

import "fmt"

func main() {
	over := make(chan bool)
		for i := 0; i < 10; i++ {
			go func() {
				fmt.Println(i)
			}()
			if i == 9 {
				over <- true
			}
		}
	<-over
	fmt.Println("over!!!")
}
