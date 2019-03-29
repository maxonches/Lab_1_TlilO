package main

import (
	"fmt"
	"../../labs_TlilO/lab2/analytical"
)


func main() {
	kernelParams := &analytical.KernelStruct {
		-15,
		6.667,
		40,
		-12,
		-24,
	}
	fmt.Println("АНАЛИТИЧЕСКОЕ РЕШЕНИЕ")
	analyticalSolve, err := analytical.AnalyticalSolverFunc(*kernelParams)
	if err != nil {
		fmt.Println("Error of analytical method!", err)
	}
	err = analytical.Solution(analyticalSolve)
	if err != nil {
		fmt.Println("Error of analytical method!", err)
	}
	fmt.Printf("Седловая точка x=%.3f, y=%.3f, H(x,y) = %.3f\n", analyticalSolve.X, analyticalSolve.Y, analyticalSolve.H)

/*=====================================================================================================================================*/

	fmt.Println("ЧИСЛЕННЫЙ МЕТОД")

}