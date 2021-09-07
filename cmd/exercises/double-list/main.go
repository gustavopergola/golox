package main

import (
	"fmt"

	"github.com/gustavopergola/golox/src/list"
)

/*

+-> [nil]X[next] -> [prev]Y[next] -> [prev]Z[next]

*/

func main() {
	list := list.List{}
	list.Insert("the first one!")
	list.Insert("the second one!")
	list.Insert("the third one!")
	list.Print()

	itemFound := list.Find("the third one!")

	if itemFound != nil {
		fmt.Printf("Found %s\n", itemFound.Value)
	}
}
