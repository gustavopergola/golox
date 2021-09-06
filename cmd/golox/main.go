package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	
	if (len(args) != 1) {
		fmt.Println("number of args must be exactly one")
		os.Exit(64)
	}
	fmt.Println(args)	
}