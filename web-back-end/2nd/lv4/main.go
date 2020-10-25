package main

import (
	"fmt"
)

type Painter interface {
	Paint()
}

type Circle struct {

}

type Heart struct {

}

func Paint (painter Painter){
	painter.Paint()
}
func (circle Circle) Paint (){
	for y := 10; y > -10; y-- {
		for x := 10; x > -10; x-- {
			if x*x+y*y < 100 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}
func (heart Heart) Paint (){
	for y := 1.4; y > -1.4; y-=0.1 {
		for x := 1.4; x > -1.4; x-= 0.1{
			if (x*x+y*y-1)*(x*x+y*y-1)*(x*x+y*y-1)-x*x*y*y*y < 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}


func main(){
	var circle Circle
	var heart Heart
	Paint(circle)
	Paint(heart)
}

