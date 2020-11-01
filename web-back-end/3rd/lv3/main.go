package main

import "fmt"

func main() {
	over := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			//错误：goroutine的位置错误；
			//修改：提升至for循环外。
			//go func() {
				fmt.Println(i)
			//}()
			if i == 9 {
				over <- true
			}
		}
	}()
	<-over
	fmt.Println("over!!!")
}
