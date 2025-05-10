package main

import "fmt"

// Animal本质上是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "Dog"
}

func Polumorphism(animal Animal) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}

func main() {
	cat := Cat{color: "black"}
	Polumorphism(&cat)
	dog := Dog{"Yellow"}
	Polumorphism(&dog)
}
