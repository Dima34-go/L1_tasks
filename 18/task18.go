package main18

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mx   sync.Mutex
}

func (c *Counter) Add() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.value++
}

func Main18() {
	c:=new(Counter)
	for i:=0;i<1000;i++{
		go func() {
			c.Add()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(c.value)
}
