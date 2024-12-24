package main

import "fmt"

type Human struct {
	Age     int
	Name    string
	Surname string
	Gender  string
}

//Встраивание в структуру
type Action struct {
	Human
	Job   string
	Hobby string
}

//Метод для Human
func (h *Human) Born() {

	h.Age = 52
	h.Name = "Ivan"
	h.Surname = "Ivanov"
	h.Gender = "Woman"
	fmt.Printf("%s %s is %s and %d years old\n", h.Name, h.Surname, h.Gender, h.Age)
}

// Метод для Action
func (a *Action) Action(h Human) {
	a.Job = "Go Developer"
	a.Hobby = "Go developping"
	fmt.Printf("He is %s and love %s", a.Job, a.Hobby)
}

func main() {
	var h = Human{}
	var a = Action{}
	h.Born()
	a.Action(h)
}
