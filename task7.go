package main

import (
	"sync"
)


type AsyncMap struct {
	m map[int]int
	mx sync.Mutex
}
func NewAsyncMap() *AsyncMap{
	return &AsyncMap{
		m:  make(map[int]int),
		mx: sync.Mutex{},
	}
}
func (a *AsyncMap) Get(key int) int {
	a.mx.Lock()
	defer a.mx.Unlock()
	return a.m[key]
}
func (a *AsyncMap) Set(key,val int) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.m[key]=val
}

//RWMap struct

type AsyncRWMap struct {
	m map[int]int
	mx sync.RWMutex
}
func NewAsyncRWMap() *AsyncRWMap{
	return &AsyncRWMap{
		m:  make(map[int]int),
		mx: sync.RWMutex{},
	}
}
func (a *AsyncRWMap) Get(key int) int {
	a.mx.RLock()
	defer a.mx.RUnlock()
	return a.m[key]
}
func (a *AsyncRWMap) Set(key,val int) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.m[key]=val
}
//пример в котором может возникнуть ошибка из - за неконкурентной записи
//func main(){
//	m:=make(map[int]int)
//	go func() {
//		m[3]=4
//	}()
//	go func() {
//		fmt.Println(m[3])
//	}()
//	m[3]=3
//}
//работа asyncMap и asyncRWMap
//func main()  {
//	am:=NewAsyncRWMap()
//	go func() {
//		am.Set(3,3)
//		am.Set(3,3)
//		am.Set(3,3)
//	}()
//	go func() {
//		am.Get(3)
//	}()
//	am.Get(3)
//	go func() {
//		am.Get(3)
//	}()
//	am.Set(3,4)
//
//	go func() {
//		am.Set(3,8)
//	}()
//	am.Set(3,5)
//}


