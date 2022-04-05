package main13

import "fmt"

func Main13() {
	a:=4
	b:=3
	fmt.Println(a,b)
	a,b=b,a
	fmt.Println(a,b)
}