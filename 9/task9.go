package main9

import (
	"fmt"
	"sync"
	"time"
)

func Main9(){
	ch12:=make(chan int)
	ch23:=make(chan int)
	//создание массива
	arr:=make([]int,1<<4)
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < len(arr); i++ {
		arr[i]=i
	}
	//генератор чисел
	go func() {
		for _,val:=range arr {
			time.Sleep(200*time.Millisecond)
			ch12<-val
		}
		close(ch12)
	}()
	//возведение в квадрат
	go func() {
		for val:=range ch12{
			ch23<-val*val
		}
		close(ch23)
	}()
	//вывод в консоль
	go func() {
		for val:=range ch23{
			fmt.Println(val)
		}
		wg.Done()
	}()
	wg.Wait()
}