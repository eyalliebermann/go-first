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

	greeting, err := greetings.Hello("Segundo")
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

	greets, err := greetings.Hellos([]string{"Anton", "Berta", "Caesar", "Dora", "Emil", "Friedrich"})

	for _, pair := range greets {
		fmt.Println(pair)
	}

	_, err = greetings.Hellos([]string{""})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("This line won't excute if a log.Fatal has been called")

}
