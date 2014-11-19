package main

import "fmt"

func main() {
	names := [2]string{"ada", "lovelace"}
	fmt.Printf("Names address: %p \n", &names)
	// f1(names)
	f2(&names)
	fmt.Println(names[0])
}

func f1(array [2]string) {
	fmt.Printf("Value: %s Addr: %p \n", array, &array)
	array[0] = "marie"
}

func f2(array *[2]string) {
	fmt.Printf("Value %p Addr: %p \n", array, &array)
	array[0] = "marie"
}

func arrays1() {
	var el [4]int

	fmt.Printf("el addr: %p \n", &el)

	for i, e := range el {
		fmt.Printf("Value[%s] IndexAddr[%p] \n", e, &el[i])
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
