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

	var pointerRanni *Dog = &ranni
	pointerRanni.Age = 4

	fmt.Println(pointerRanni.Age)
	fmt.Println(ranni.Age)

	ranni.Age = 5

	fmt.Println(pointerRanni.Age)
	fmt.Println(ranni.Age)

	go showMessage()
	fmt.Println("This message will appear before the message from showMessage")
	// time.Sleep(1 * time.Second)

	go meet.SayHello()
	time.Sleep(500 * time.Millisecond)

	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)

		fmt.Println("End of write")
	}()

	value := <-ch
	fmt.Println("Value from chanel:", value)
	fmt.Println("Nexts values from chanel")

	for i := range ch {
		fmt.Println("Value from chanel:", i)
	}

	channel := make(chan int)
	go producer(channel)
	go consumer(channel)
	go consumer(channel)
	go consumer(channel)

	time.Sleep(time.Second * 1)
}

func showMessage() {
	fmt.Println("Hello from goroutine!")
}

func producer(ch chan int) {
	for i := 0; i < 30; i++ {
		ch <- i
	}
	close(ch)

	fmt.Println("End of producer write")
}

func consumer(ch chan int) {
	for i := range ch {
		fmt.Println("Value from new chanel:", i)
	}
	fmt.Println("End of consumer reading")
}
