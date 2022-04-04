package main

import (
	"fmt"
	"sync"
)

//возведение в квадрат примерно в 1000 раз быстрее чем вывод числа
//
//func main(){
//	runtime.GOMAXPROCS(100)
//	n:=1<<16
//	arr:=make([]int,n)
//	for i:=0;i<n;i++{
//		arr[i]=i
//	}
//	ts:=time.Now()
//	f5(arr)
//	fmt.Println("Прошедшее время",time.Now().Sub(ts).Milliseconds())
//}

func f5(arr []int) {
	ch := make(chan int, 10)
	go func(c chan int) {
		for _, val := range arr {
			c <- val * val
		}
		close(c)
	}(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

//лучший вариант, но все еще хуже простого вывода подряд)))
func f4(arr []int) {
	k := 0
	go func() {
		for ; k < len(arr); k++ {
			arr[k] *= arr[k]
		}
		k++
	}()
	for i := 0; i < len(arr); {
		if i < k {
			fmt.Println(arr[i])
			i++
		}
	}
}

//без учета порядка
func f3(arr []int) {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	for k := 0; k < len(arr); k++ {
		i := k
		go func() {
			fmt.Println(arr[i] * arr[i])
			wg.Done()
		}()
	}
	wg.Wait()
}

//с учетом порядка
func f2(arr []int) {
	arrBool := make([]bool, len(arr))
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for k := 0; k < len(arr); k++ {
		i := k
		go func() {
			arr[i] *= arr[i]
			arrBool[i] = true
			wg.Done()
		}()
	}
	for i := 0; i < len(arr); {
		if arrBool[i] {
			fmt.Println(arr[i])
			i++
		}
	}
}

//обычный счет
func f1(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i] * arr[i])
	}
}
