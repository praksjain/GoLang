package main

import "fmt"

const unsixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

// Struct - a type of class in golang
type car struct {
	gas_pedal uint64 // 0 to 65535
	brake_pedal uint64
	steering_wheel int16 //-32k to 32k
	top_speed_kmh float64
}

// Value receiver methods
func (c car) kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/unsixteenbitmax)
}

// Value receiver methods
func (c car) mph() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/unsixteenbitmax/kmh_multiple)
}

//Pointer receiver
func (c *car) new_top_speed(newspeed float64){
	c.top_speed_kmh = newspeed
}

// Normal way
func newer_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

// Execution main method
func main() {
	a_car := car{gas_pedal: 22341,
				brake_pedal: 0,
				steering_wheel: 12561,
				top_speed_kmh: 225.0}
	
	fmt.Println("Gas pedal value:", a_car.gas_pedal)
	fmt.Println("Kilometers per hour value:", a_car.kmh())
	fmt.Println("Miles per hour value:", a_car.mph())
	a_car.new_top_speed(5000)
	fmt.Println("Kilometers per hour value:", a_car.kmh())
	fmt.Println("Miles per hour value:", a_car.mph())
	a_car = newer_speed(a_car, 5000)
	fmt.Println("Kilometers per hour value:", a_car.kmh())
	fmt.Println("Miles per hour value:", a_car.mph())
}