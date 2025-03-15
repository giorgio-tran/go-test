package main

import "fmt"

// importing from external package
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}
