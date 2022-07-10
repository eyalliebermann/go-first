package main

import (
	"fmt"

	"eyalliebermann.com/greetings"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(greetings.Hello("World!") + "\n" + quote.Go())

}
