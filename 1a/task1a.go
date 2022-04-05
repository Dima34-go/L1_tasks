package main1

import "fmt"

//Структура Human, является частью структуры Action, из нее можно вызвать методы принадлежащие Human

type Human struct {
	Hands   string
	Legs    string
}
type Action struct {
	 Human   Human
	 Object  string
}
func (h *Human) Dance() {
	fmt.Println(h.Hands,h.Legs)
}
func (h *Human) Run() {
	fmt.Println(h.Legs)
}
func (a *Action) Use(){
	fmt.Println("Human use",a.Object)
}
func (a *Action) Dance() {
	a.Human.Dance()
}
func (a *Action) Run() {
	a.Human.Run()
}

func Main1a(){
	Tom:=Human{
		Hands: "top-top",
		Legs:  "hlop-hlop",
	}
	obj:=Action{
		Human:  Tom,
		Object: "ball",
	}
	//Можно определить метод для структуры Action
	obj.Run()
	obj.Human.Run()
}
