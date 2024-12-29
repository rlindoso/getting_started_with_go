package main

import (
	"fmt"
	"time"

	"github.com/rlindoso/getting_started_with_go/meet"
)

type Dog struct {
	Name  string
	Age   int
	Breed string
}

func (dog Dog) printDog() {
	fmt.Printf("Hello, my name is %s, I`m %d years old and my breed is %s\n", dog.Name, dog.Age, dog.Breed)
}

func main() {
	meet.Say("Hello, my friend!")

	test := "Hello"
	test = "Hello again"

	meet.Say(test)

	const testConst = "Lindoso"

	meet.Say(testConst)

	var slices []string

	slices = append(slices, "index 0", "index 1")
	slices = append(slices, "index 2", "index 3", "index 4")

	fmt.Println(slices)
	fmt.Println(slices[3])
	fmt.Println(slices[2:4])
	fmt.Println(slices[:3])
	fmt.Println(slices[2:])

	var dogs = map[string]int{}
	dogs["Ranni"] = 3
	dogs["Melina"] = 7

	if age, ok := dogs["Branco"]; ok {
		fmt.Println("Dog exists on map", age, ok)
	} else {
		fmt.Println("Dog doesn`t exists")
	}

	if age, ok := dogs["Melina"]; ok {
		fmt.Println("Dog exists on map", age, ok)
	} else {
		fmt.Println("Dog doesn`t exists")
	}

	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Is today")
	case today + 1:
		fmt.Println("It's tomorrow")
	case today + 2:
		fmt.Println("It's in two days")
	default:
		fmt.Println("It will take a long time to arrive")
	}

	melina := Dog{
		Name:  "Melina",
		Age:   7,
		Breed: "Stray",
	}

	ranni := Dog{
		Name:  "Ranni",
		Age:   3,
		Breed: "English Pointer",
	}

	melina.printDog()
	ranni.printDog()
}
