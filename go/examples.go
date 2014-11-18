package main

import "fmt"

func main() {
	arrays1()
	// names := [2]string{"ada", "lovelace"}
	// println("Names address:", &names)
	// f1(names)
	// f2(&names)
	// fmt.Println(names[0]) // still prints "ada"
}

func f1(array [2]string) {
	fmt.Printf("Value: %s", array)
	println(" Addr:", &array)
	array[0] = "marie"
}

func f2(array *[2]string) {
	fmt.Printf("Value: %s", array)
	println(" Addr:", &array)
	array[0] = "marie"
}

func arrays1() {
	var a [4]int

	println("a addr:", &a)

	for i, e := range a {
		fmt.Printf("Value[%d] ", e)
		println("IndexAddr:", &a[i])
	}
}

func arrays2() {
	// declare an array of int pointers
	pointers := [2]*int{0: new(int), 1: new(int)}
	*pointers[0] = 20
	*pointers[1] = 30

	fmt.Printf("address of pointers array: %p \n", &pointers)

	for index, ptr := range pointers {
		fmt.Printf("Value[%d] IndexAddr[%p] \n", *pointers[index], ptr)
	}
}
