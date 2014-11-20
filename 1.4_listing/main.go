package main

func main() {
	names := [2]string{"ada", "lovelace"}
	println("names address:", &names)
	f1(names)
	println(names[0]) // still prints "ada"
}

func f1(a [2]string) {
	println("value:", a[0], a[1])
	println("a address:", &a)
	a[0] = "marie"
}
