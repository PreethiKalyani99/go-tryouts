package main

import (
	"fmt"
	"rsc.io/quote"
	"log"
)

func main() {
	fmt.Println(quote.Go());
	fmt.Println(quote.Opt());
	fmt.Println(quote.Hello());
	
	message, err := Greet("Preethi");

	if err != nil {
		log.Fatal(err); 
	}
		
	fmt.Println(message);
}