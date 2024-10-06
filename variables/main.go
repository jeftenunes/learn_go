package main

import "fmt"

func main() {
	var a = "initial value"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Printf("%v, %v\n", b, c)

	d := "shortcut for declaring and initializing a variable"
	fmt.Println(d)
}
