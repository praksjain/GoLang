///////////////////////////////////////////////////////////////////////////////////////////////////
/* Module for BASICS OF GO WEB. */
///////////////////////////////////////////////////////////////////////////////////////////////////


/*Copyright © https://github.com/praksjain/golang
Author : Prakhar Jain
Version : 1.0
Description : Script to understand basics of go web.
File Name : web.go
*/

// Main
package main

// Built -in imports
import ("fmt"
		"net/http")


// index handler (localhost:8000/)
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whoa, Go is neat")
}

// about handler (localhost:8000/about/)
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Designed by Prakhar Jain</h1>")
}

// main method
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000",nil)
}