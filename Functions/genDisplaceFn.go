package main

import "fmt"

func main() {
	var A, V0, S0 float64
	var t float64
	fmt.Printf("Input value of Acceleration(a) \t\t\t:=>")
	fmt.Scanln(&A)

	fmt.Printf("Input value of Initial Velocity(V0) \t\t:=>")
	fmt.Scanln(&V0)

	fmt.Printf("Input value of Initial Displacement(S0) \t:=>")
	fmt.Scanln(&S0)

	fmt.Printf("Input value of Time(t) \t\t\t\t:=>")
	fmt.Scanln(&t)
	fmt.Println("------------Isaac---on---go-----------")

	fn := GenDisplaceFn(A, V0, S0)
	fmt.Printf("Displacement after %v Seconds \t\t\t:=>%v\n", t, fn(t))
	fmt.Println("------------Isaac---on---go-----------------")

	
}

func GenDisplaceFn(acc, iniVelocity, iniDisplacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*(t*t) + iniVelocity*t + iniDisplacement
	}
}