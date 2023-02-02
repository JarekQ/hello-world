/*
Followed along a Pluralsight course on GO:
https://app.pluralsight.com/library/courses/getting-started-with-go/table-of-contents

Original Author: Mike Van Sickle

Modified By: JarekQ Aloisio
*/
package main

import (
	"JQ_GO/controllers"
	"JQ_GO/models"
	"errors"
	"fmt"
	"net/http"
)

func main() {
	/*
		helloWorld()
		introPDT()
		introPointers()
		introConstants()
		introConstExpress_Iota()
		modelUser()
		startWebServer_v1(3000,2)

		isStarted := startWebServer_v2(3000, 2)
		fmt.Println(isStarted)

		err := startWebServer_v3(3000, 2)
		fmt.Println(err)

		port, err := startWebServer_v4(3000, 2)
		fmt.Println(port, err)
	*/

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

func startWebServer_v4(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting Server...")
	// do something
	fmt.Println("Server Started on port:", port)
	fmt.Println("Number of retries:", numberOfRetries)
	//return nil
	return port, nil
}

func startWebServer_v3(port int, numberOfRetries int) error {
	fmt.Println("Starting Server...")
	// do something
	fmt.Println("Server Started on port:", port)
	fmt.Println("Number of retries:", numberOfRetries)
	//return nil
	return errors.New("something went wrong")
}

func startWebServer_v2(port int, numberOfRetries int) bool {
	fmt.Println("Starting Server...")
	// do something
	fmt.Println("Server Started on port:", port)
	fmt.Println("Number of retries:", numberOfRetries)
	return true
}

func startWebServer_v1(port int, numberOfRetries int) {
	fmt.Println("Starting Server...")
	// do something
	fmt.Println("Server Started on port:", port)
	fmt.Println("Number of retries:", numberOfRetries)
}

func modelUser() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}

func introStructs() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{
		ID:        1,
		FirstName: "John",
		LastName:  "Connor",
	}

	fmt.Println(u2)
}

func introMaps() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}

func introSlices() {
	arr3 := [3]int{1, 2, 3}
	slice := arr3[:]
	fmt.Println(arr3, slice)

	arr3[1] = 42
	slice[2] = 27

	fmt.Println(arr3, slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2)

	s2 := slice2[1:]
	s3 := slice2[:2]
	s4 := slice2[1:2]
	fmt.Println(s2, s3, s4)
}

func introArrays() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	//fmt.Println(arr[4]) // error thrown out of bounds
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
