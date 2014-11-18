package main

import "fmt"

func main() {
	names := [2]string{"ada", "lovelace"}
	println("names address:", &names)
	f1(&names)
	println(names[0]) // now prints "marie"
}

func f1(a *[2]string) {
	println("a value:", a)
	println("a address:", &a)
	a[0] = "marie"

	// Do this to prevent inlining.
	var x int
	fmt.Sprintf("Prevent Inlining: %d", x)
}
