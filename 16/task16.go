package main16

import (
	"fmt"
	"math/rand"
)

func Main16() {
	arr:=make([]int,1<<5)
	for i := 0; i < len(arr); i++ {
		arr[i]=rand.Int()%(1<<4)
	}
	quicksort(arr)
	fmt.Println(arr)
}
func quicksort(arr []int){
	quickSortTemplate(arr,0,len(arr)-1)
}
func quickSortTemplate(arr []int,l,r int) {
	if r-l < 1 {
		return
	}
	rv:=arr[l+rand.Int()%(r-l+1)]
	//подсчет
	i,j:=l,r
	for ;i < j; {
		if arr[i]<rv {
			i++
		} else if arr[j] >= rv {
			j--
		} else{
			arr[i],arr[j]=arr[j],arr[i]
		}
	}
	midl:=i-1
	for j=r;i<j; {
		if arr[i] == rv{ i++ }else if arr[j]>rv { j-- } else{
			arr[i],arr[j]=arr[j],arr[i]
		}
	}
	midr:=i
	quickSortTemplate(arr,l,midl)
	quickSortTemplate(arr,midr,r)
}