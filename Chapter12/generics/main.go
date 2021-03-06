// This code demonstrates how we can use generics in GO using code generators
// Here we use genny as the code generator
package main

import (
	"./list"
	"fmt"
)

// To install genny do 'go get github.com/cheekybits/genny'
// For more information refer : https://github.com/cheekybits/genny

func main() {
	// Create a unit List
	// Note that this struct and related methods are autogenerated
	ul := list.NewUintList()

	// insert 1
	ul.Add(1)

	// ensure we get it back!
	fmt.Println(ul.Get())
}
