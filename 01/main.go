package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide an input file")
	}

	fmt.Printf("%s\n%s\n", os.Args[0], os.Args[1])
}

