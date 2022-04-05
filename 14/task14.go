package main14

import "fmt"

func Main14(){
	var secret interface {}
	secret="00"
	fmt.Println(secret)
	switch secret.(type) {
	case int:
		fmt.Println("It is int",secret.(int))
	case string:
		fmt.Println("It is string",secret.(string))
	case bool:
		fmt.Println("It is bool",secret.(bool))
	case chan int:
		fmt.Println("It is chan int",secret.(chan int))
	default:
		fmt.Printf("Untypical type interface %T",secret)
	}
}