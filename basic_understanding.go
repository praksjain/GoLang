///////////////////////////////////////////////////////////////////////////////////////////////////
/* Module for BASICS OF GO . */
///////////////////////////////////////////////////////////////////////////////////////////////////


/*Copyright © https://github.com/praksjain/golang
Author : Prakhar Jain
Version : 1.0
Description : Script to understand basics of go (adding, concatenating, pointers).
File Name : basic_understanding.go
*/

// Main
package main

// Built -in imports
import "fmt"


// adding two numbers and returning float value
func add(x float64, y float64) float64{
	return x+y
}

// concatenating two strings
func concat(x string, y string) (string, string) {
	return x,y
}

// reassigning the values
func reassignment() float64 {
	var num1 int = 61
	var num2 float64 = float64(num1)
	return num2
}

// use of pointers
func point () {
	x := 5
	a := &x
	fmt.Println("memory address: ", a)
	fmt.Println("Value:", *a)
	fmt.Println("New address of a:", &a)
	fmt.Println("New value:", *&a)
	fmt.Println("New value:", **&a)
}

// main method to call operations
func main() {
	fmt.Println("welcome go..!!")
	num1, num2 := 3.2, 3.2
	fmt.Println("Sum of two numbers is:", add(num1, num2))
	a1, a2 := "Prakhar", "Jain"
	fmt.Println(concat(a1,a2))
	fmt.Println(reassignment())
	point()
}
