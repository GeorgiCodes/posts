package main

func main() {
	names := [2]string{"ada", "lovelace"}
	println("names address:", &names)
	f1(&names)
	println(names[0]) // prints "marie"
}

func f1(a *[2]string) {
	println("value:", a)
	println("a address:", &a)
	a[0] = "marie"
}
