package main25

import (
	"fmt"
	"time"
)

//
func Main25() {
	st:=time.Now()
	//Sleep(time.After(1*time.Second))
	//time.Sleep(1*time.Second)
	fmt.Println("Времени прошло:", time.Now().Sub(st).Milliseconds())
}

func Sleep(ch <- chan time.Time){
	<-ch
	return
}