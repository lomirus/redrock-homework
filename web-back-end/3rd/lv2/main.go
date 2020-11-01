package main

import (
	"fmt"
	"os"
)

func main() {
	//Init
	buf := make([]byte, 1024)
	//Create
	proverb, _ := os.Create("proverb.txt")
	defer proverb.Close()
	//Write
	proverb.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	//Read
	proverb, _ = os.Open("proverb.txt")
	n, _ := proverb.Read(buf)
	//Output
	fmt.Println(string(buf[:n]))
}
