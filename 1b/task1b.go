package main1b

import "fmt"

type Human struct {
	Hands   string
	Legs    string
}
type Action struct {
   Human
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
	fmt.Println("action Dance")
}
func (a *Action) Run() {
	fmt.Println("action Run")
}
func Main1b(){
	Tom:=Human{
		Hands: "top-top",
		Legs:  "hlop-hlop",
	}
	obj:=Action{
		Human:  Tom,
		Object: "ball",
	}
	//реализовано встраивание, метод для структур выше можно будет вызвать
	obj.Run()
	obj.Human.Run()
}
