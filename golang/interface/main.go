package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Woof, my name is %s", d.name)
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{
		name: "Buddy",
	}
	cat := Cat{}

	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
}
