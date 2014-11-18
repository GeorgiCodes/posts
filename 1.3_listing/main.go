package main

func main() {
	var a [4]int

	println("a addr:", &a)

	for i, e := range a {
		print("Value:", e)
		println(" IndexAddr:", &a[i])
	}
}
