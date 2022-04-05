package main10

import (
	"fmt"
	"sort"
)

func Main10() {
	arr:=[]float64{ -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]<arr[j]
	})
	m:=make(map[int][]float64)
	//за o(n) проходим и записываем в map, потом выводи map
	for _,val:=range arr {
		var key int
		//получение ключа
		if val < 0 {
			key = (int(val/10)-1)*10
		}else {
			key=int(val/10)*10
		}
		//поиск нужного места
		m[key] = append(m[key], val)
	}
	fmt.Println(m)
}