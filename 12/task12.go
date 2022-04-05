package main12

import "fmt"

func Main12() {
	allStr:=[]string{ "cat", "cat", "dog", "cat", "tree"}
	type void struct{}
	var exist void
	m:=make(map[string]void)
	for _,str := range allStr {
		m[str]=exist
	}
	fmt.Println(m)
}