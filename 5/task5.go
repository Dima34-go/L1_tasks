package main5

import (
	"fmt"
	"math/rand"
	"time"
)

func Main5 (){
	N:=5
	ch :=make(chan int)
	go func() {
		for true {
			ch<-rand.Int()
		}
	}()
	go func() {
		for val := range ch{
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Duration(N)*time.Second)
}
