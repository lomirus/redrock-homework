package main

import (
	"fmt"
)

//Receiver ...
func Receiver(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case float32:
		fmt.Println("float32")
	case float64:
		fmt.Println("float64")
	case int:
		fmt.Println("int")
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case int32:
		fmt.Println("int32")
	case int64:
		fmt.Println("int64")
	}
}

func main() {
	Receiver(true)
}
