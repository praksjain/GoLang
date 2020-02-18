package main


import "fmt"
		// "math")
// import "math"

func add(x float64, y float64) float64{
	return x+y
}

func concat(x string, y string) (string, string) {
	return x,y
}

func reassignment() float64 {
	var num1 int = 61
	var num2 float64 = float64(num1)
	return num2
}

func point () {
	x := 5
	a := &x
	fmt.Println("memory address: ", a)
	fmt.Println("Value:", *a)
	fmt.Println("New address of a:", &a)
	fmt.Println("New mohammed value:", *&a)
	fmt.Println("New mohammed value:", **&a)
}

func main() {
	fmt.Println("welcome go..!!")
	num1, num2 := 3.2, 3.2
	fmt.Println("Sum of two numbers is:", add(num1, num2))

	a1, a2 := "Prakhar", "Jain"
	fmt.Println(concat(a1,a2))

	fmt.Println(reassignment())

	point()
}
