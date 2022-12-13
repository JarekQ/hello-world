/*
Followed along a Pluralsight course on GO:
https://app.pluralsight.com/library/courses/getting-started-with-go/table-of-contents

Original Author: Mike Van Sickle

Modified By: JarekQ Aloisio
*/
package main

import "fmt"

func main() {
	helloWorld()
	introPDT()
	introPointers()
	introConstants()
	introConstExpress_Iota()
}

// First Constant Block
const (
	pi  = 3.1415
	fst = 1
	snd = "second"
)

// Second Constant Block
const (
	first = iota
	//iota + 6
	//iota

	second // = iota
	//2 << iota
	//
)

// Third Constant Block
const (
	third = iota
	//fourth
)

func introConstExpress_Iota() {
	fmt.Println(pi, fst, snd)
	fmt.Println(first, second, third) //, fourth)
}

func introConstants() {
	const pi = 3.1415
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)

	// a bunch of code

	fmt.Println(c + 1.2)
}

// Hello World :)
func helloWorld() {
	fmt.Println("Hello, World!")
}

// Intro to Pointers
func introPointers() {
	var newFirstN *string = new(string)

	*newFirstN = "Bob"

	// Memory Address
	fmt.Println(newFirstN)
	// Memory Value | Used the (*) dereference operator
	fmt.Println(*newFirstN)

	fName := "Charlie"
	fmt.Println(fName)

	// Pointer Reference Declaration & Assignment
	ptr := &fName
	// Pointer Reference with Pointer's value
	fmt.Println(ptr, *ptr)

	// Overwrite existing Pointer Value
	fName = "Beth"
	// Pointer Reference with Pointer's value
	fmt.Println(ptr, *ptr)
}

// Intro to Primitive Data Types
func introPDT() {
	// Declare variable int | Ignore Warning
	var i int32
	// Assign variable int
	i = 42
	// Print variable int
	fmt.Println(i)

	// Declare and Assign variable float32
	var f float32 = 3.14
	// Print variable float32
	fmt.Println(f)

	// Declare and Assign implicit string
	firstName := "Arthur"
	// Print string
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
