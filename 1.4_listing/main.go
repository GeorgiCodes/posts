package main

import "fmt"

func main() {
	names := [2]string{"ada", "lovelace"}
	println("names address:", &names)
	f1(names)
	fmt.Println(names[0]) // still prints "ada"
}

func f1(a [2]string) {
	fmt.Printf("value: %s,", a)
	println(" a address:", &a)
	a[0] = "marie"
}
