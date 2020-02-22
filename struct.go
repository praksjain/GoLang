///////////////////////////////////////////////////////////////////////////////////////////////////
/* Module for BASICS OF GO STRUCT. */
///////////////////////////////////////////////////////////////////////////////////////////////////


/*Copyright © https://github.com/praksjain/golang
Author : Prakhar Jain
Version : 1.0
Description : Script to understand basics of go struct.
File Name : struct.go
*/

// Main
package main

// Built -in imports
import "fmt"

type car struct {
	gas_pedal uint64 // 0 to 65535
	brake_pedal uint64
	steering_wheel int16 //-32k to 32k
	top_speend_kmh float64
}

// Main method
func main() {
	a_car := car{gas_pedal: 22341,
				brake_pedal: 0,
				steering_wheel: 12561,
				top_speend_kmh: 225.0}

	fmt.Println(a_car.gas_pedal)
}