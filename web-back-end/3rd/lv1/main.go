package main

import "fmt"

var oddNum, evenNum = make(chan int), make(chan int)
var exit = make(chan bool)

func main(){
	go odd()
	go even()
	oddNum <- 1
	<-exit
}
func odd(){
	for i := 1;i < 52;i++{
		var temp = <- oddNum
		fmt.Printf("odd%d\n",temp)
		evenNum <- temp+1
	}
}

func even(){
	for i := 1;i < 51;i++{
		var temp = <- evenNum
		fmt.Printf("even%d\n",temp)
		oddNum <- temp+1
	}
	exit <- true
}