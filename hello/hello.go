package main

import (
	"fmt"
	"log"

	"eyalliebermann.com/greetings"
	"rsc.io/quote/v4"
)

func main() {

	greeting, _ := greetings.Hello("Uno!")
	fmt.Println("1. " + greeting + "\n" + quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	greeting, err := greetings.Hello("Second")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(greeting)
	}

	greeting, err = greetings.Hello("")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(greeting)
	}

	greeting, err = greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(greeting)
	}

	fmt.Println("This line won't excute if a log.Fatal has been called")

}
