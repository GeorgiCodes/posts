package main

import "fmt"

func main() {
	var a [4]int

	println("a addr:", &a)

	for i, e := range a {
		fmt.Printf("Value[%d] ", e)
		println("IndexAddr:", &a[i])
	}
}
