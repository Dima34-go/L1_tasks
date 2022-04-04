package main

import (
	"runtime"
)

type Summator struct {
	first int
	ch    chan int
	sum   int
}

//func main() {
//	arr := make([]int, 1<<20)
//	for i:=0;i<len(arr);i++ {
//		arr[i]=i
//	}
//	ts:=time.Now()
//	sum:=fs1(arr)
//	fmt.Println(time.Now().Sub(ts).Microseconds())
//	fmt.Println(sum)
//}
func fs2(arr [] int) int {
	numCPU := runtime.NumCPU()
	ch:=make(chan int)
	s := make([]Summator, numCPU)
	for i := 0; i < numCPU; i++ {
		s[i].first = i
		s[i].ch = ch
		s[i].sum = 0
	}
	for k := 0; k < numCPU; k++ {
		i := k
		go func() {
			for j := s[i].first; j < len(arr); j += numCPU {
				s[i].sum+=arr[j]*arr[j]
			}
			s[i].ch<-s[i].sum
		}()
	}
	fullSum:=0
	for i:=0;i<numCPU;i++{
		fullSum+=<-ch
	}
	return fullSum
}
func fs1(arr []int) int {
	sum:=0
	for _,val:=range arr {
		sum+=val*val
	}
	return sum
}
