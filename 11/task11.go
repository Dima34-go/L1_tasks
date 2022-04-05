package main11

import "fmt"

func Main11() {
	arr1:=[]int{10, 12, 34, -12, 30}
	arr2:=[]int{13, 10, -12, 43, 65}
	m1,m2:=make(map[int]bool),make(map[int]bool)
	//два массива данных превращаются в map
	for _, val := range arr1 {
		m1[val]=true
	}
	for _, val := range arr2 {
		m2[val]=true
	}
	// в третью map закидываем элементы, которые есть в каждой из map1 и map2
	m3:=make(map[int]bool)
	for key := range m1 {
		if m1[key] == m2[key]{
			m3[key] = true
		}
	}
	fmt.Println(m3)
}
