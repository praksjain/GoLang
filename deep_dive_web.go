///////////////////////////////////////////////////////////////////////////////////////////////////
/* Module for BASICS OF GO MORE UNDERSTANDING ON WEB. */
///////////////////////////////////////////////////////////////////////////////////////////////////


/*Copyright © https://github.com/praksjain/golang
Author : Prakhar Jain
Version : 1.0
Description : Script for go more web understanding.
File Name : deep_dive_web.go
*/

// Main
package main

// Built -in imports
import ("fmt"
		"net/http")
		// "io/ioutil")


var washPostXML = []byte(`
<sitemap)
func main() {
	resp, value := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	fmt.Println("response is :", resp)
	fmt.Println("dash is :", value)
}
