package main

import "fmt"

type Animal interface {
	vocalise()
	move()
}

type Dog struct {
	name  string
	breed string
}

func (d Dog) vocalise() {
	fmt.Printf("Woof!\n")
}

func (d Dog) move() {
	fmt.Printf(d.name + " is sprinting...\n")
}

type Cat struct {
	name  string
	breed string
}

func (c Cat) vocalise() {
	fmt.Printf("Meow!\n")
}

func (c Cat) move() {
	fmt.Printf(c.name + " is slinking away...\n")
}

type Human struct {
	name    string
	age     uint
	country string
}

func (h Human) vocalise() {
	fmt.Printf("Yarrrhggggghhhh!\n")
}

func (h Human) move() {
	fmt.Printf(h.name + " is sneaking off...\n")
}

func main() {
	roofus := Dog{name: "Roofus", breed: "Labrador"}
	sparkles := Cat{name: "Sparkles", breed: "Persian"}
	bob := Human{name: "Bob", age: 42, country: "U.S.A"}

	animals := []Animal{roofus, sparkles, bob}

	for _, animal := range animals {
		animal.vocalise()
		animal.move()
		fmt.Println()
	}
}
