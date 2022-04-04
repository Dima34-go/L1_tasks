package main

import (
	"fmt"
	"math/rand"
)

// нужно занести инициализацию переменной в тело функции
// 1) justString - не очищается после завершения работы функции
// 2) При одновременной работе нескольких  функций с данной могут возникать проблемы, так как justString
// в глобальной области видимости
func someFunc() {
	var justString string
	v := createHugeString(1 << 7)
	fmt.Println(v,&v)
	justString = v[:2]
	fmt.Println(justString,&justString)
	justString=v
	fmt.Println(v,&v)
	fmt.Println(justString,&justString)
}

//func main() {
//	someFunc()
//}


func createHugeString(size int) string{
	sumbs:=make([]rune,size)
	for i := 0; i < size; i++ {
		sumbs[i]=int32(62 + rand.Int()%27)
	}
	return string(sumbs)
}