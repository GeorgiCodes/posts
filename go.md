## Grokking Pass by Value in Go

One of the lines you will hear over and over again as you learn go is that it is **pass by value**. This post describes how arrays are represented in memory in Go and what it means to pass by value.

This post was spurred on by the fact that last week I had the opportunity to attend the [GothamGo](http://gothamgo.com/) conference here in New York which included a Go Workshop taught by [Bill Kennedy](http://www.goinggo.net/), one of the authors of [Go in Action](http://www.manning.com/ketelsen/). The best part of the workshop was that we dove down into some of these fun lower level language details. 

### Arrays are stored contiguously in memory
An array in Go is a container data structure of fixed length with values of a single type. The declaration below creates a zero-value array whose elements themselves are zero.

```go
var elements [4]int
```
In Go, when you declare a value of type `int` then the actual size of the `int` will be determined based on the type of architecture the program is run on. In my case I am running this program on my mac which is 64bit architecture so each `int` will be 8 bytes. It is important to note that `int` is its own type and is not an alias for `int64`. 

Let's print out the memory address of the array and of each of its elements. 
```go
var elements [4]int

fmt.Printf("address of elements array: %p \n", &elements)

for index, element := range elements {
	fmt.Printf("Value[%d] IndexAddr[%p]\n", element, &elements[index])
}

address of elements array: 0x2081ac000
Value[0] IndexAddr[0x2081ac000]
Value[0] IndexAddr[0x2081ac008]
Value[0] IndexAddr[0x2081ac010]
Value[0] IndexAddr[0x2081ac018]
```

Interanlly an array is stored **contiguously** in memory.

![](images/go_initialized_array.jpg)

These memory addresses are in hexidecimal with each index located 8 bytes ahead of the last. See [Hexadecimal to Decimal](http://www.binaryhexconverter.com/hex-to-decimal-converter) converter.

In Go, the length of the array forms part of its type. This allows us to loop through the array and access individual elements very quickly using index arithmetic. The below assignment will throw an error:

```go
var elements [4]int
var longElements [8]int

// ERROR: cannot use longElements (type [8]int) as type [4]int in assignment
elements = longElements		
```

Contiguous memory is an advantage because it assists with keeping the used memory in the CPU cache. This in turn has performance benefits because the CPU doesn't have to look all the way back to the RAM to access that memory.

#### Go is pass by value
In Go, everything is **pass by value**. This means that when we pass an array as an argument, we pass a copy of the value of the array not the reference to the array.

Lets say we have the following program:
```go
func main() {
	names := [2]string{"ada", "lovelace"}
	fmt.Printf("Names Address[%p] \n", &names)
	f1(names)
	fmt.Println(names[0])	// "ada"
}

func f1(array [2]string) {
	fmt.Printf("Array Address[%p] \n", &array)
	array[0] = "marie"
}
```
In Go terminology we would say that `array` is a **value receiver**.

![](images/call_stack_1.png)

##### There are two important things to take note of here: 
1. A copy of the value of `names` array is made when the `f1` function is called. <br/>
1. This means in `f1` when we change the value of the first element, we are making a change to the copy <br/>

Copying the value of the array might be ok for small sized arrays, but what if the `names` array had millions of strings? The stack is starting to have to do a lot of work - creating and releasing megs of memory each time the `f1` function is called. Passing by value here also doesn't allow us to alter the contents of the original array.

##### Use a pointer!
One way to overcome this would be to instead pass a pointer to the `names` array (ie. `array *[2]string`). Pointers in Go only take up 8 bytes. 

```go
func main() {
	names := [2]string{"ada", "lovelace"}
	fmt.Printf("Names Address[%p] \n", &names)
	f1(&names)
	fmt.Println(names[0])	// "marie"
}

func f1(array *[2]string) {
	fmt.Printf("Array Address[%p] \n", &array)
	*array[0] = "marie"
}
```
In Go terminology we would say that `array` is a **pointer receiver**.

![](images/call_stack_2.png)

By using a pointer we reduce the size of the `f1` stack frame and this also allows us to change the value that the pointer points too ie. the orginal array.

##### When do I use a `pointer` vs `value` receiver?
There are a couple of scenarios for when using a pointer receiver is the right choice: <br/>
1. Modification. If you need to modify the contents of the receiver. <br/>
2. Performance. If the receiver is large then using a `pointer` receiver will be more efficient. <br/>

This has been a brief introduction into arrays internals and pass by value in Go. I've tried to use digarams to further illustrate these concepts. 

Next up I will talk about slices!

## References & Reading
* [Go Data Structures](http://research.swtch.com/godata)
* [Go Slices: usage and internals](http://blog.golang.org/go-slices-usage-and-internals)
* [Go in Action (book)](http://www.manning.com/ketelsen/)
