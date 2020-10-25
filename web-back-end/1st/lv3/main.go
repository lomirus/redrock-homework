package main

import "fmt"

func main(){
	var arr []float64 = []float64{-88, -25, -12, -8, -0.5, 1, 5, 6, 77}
	fmt.Printf("%f", getAbsMin(arr))
}

func getAbsMin(arr []float64)(minAbs float64){
	minAbs = abs(arr[0])
	for _, i := range arr {
		if abs(minAbs) > abs(i){
			minAbs = i
		}
	}
	return minAbs
}
func abs(a float64)(abs float64){
	if a < 0 {
		return -a
	} else {
		return a
	}

}
