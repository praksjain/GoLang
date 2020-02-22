///////////////////////////////////////////////////////////////////////////////////////////////////
/* Module for BASICS OF GO ITERATION. */
///////////////////////////////////////////////////////////////////////////////////////////////////


/*Copyright © https://github.com/praksjain/golang
Author : Prakhar Jain
Version : 1.0
Description : Script to understand basics of go iteration.
File Name : iteration.go
*/

// Main
package main

// Built -in imports
import "fmt"

// Main method
func main() {
	for i:=1;i<=10;i++{
		s := ""
		if i % 3 == 0 {
			s += "Jain"
		}
		if i % 5 == 0 {
			s += "Prakhar"
		}
		if s != "" {
			fmt.Println(s)
		}else{
			fmt.Println(i)
		}
	}
}