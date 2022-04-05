package main21

import "fmt"

type CarDoor interface {
	OpenDoor()
}
//Key - структура, реализующая интерфейс
type Key struct {}

func (k *Key) OpenDoor() {
	fmt.Println("You use key and open door")
}
//SmartKey - структура, для использования которой необходим паттерн адаптер
type SmartKey struct {}

func (s *SmartKey) Open()  {
	fmt.Println("You use smart key and open door")
}
//SmartKeyAdapter - адаптер для структуры SmartKey под интерфейс CarDoor
type SmartKeyAdapter struct {
	*SmartKey
}

func (k *SmartKeyAdapter) OpenDoor()  {
	k.Open()
}
