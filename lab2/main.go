package main

import (
	"./numerical"
	"../../labs_TlilO/lab2/analytical"
	"fmt"
	"math"
)


func main() {
	kernelParams := &analytical.KernelStruct {
		A: -15,
		B: 6.667,
		C: 40,
		D: -12,
		E: -24,
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
	fmt.Printf("Седловая точка x=%.3f, y=%.3f, H(x,y) = %.3f\n",
		analyticalSolve.X, analyticalSolve.Y, analyticalSolve.H)

/*=====================================================================================================================================*/

	fmt.Println("ЧИСЛЕННЫЙ МЕТОД")

	numericalSolve := numerical.SolveNumerical(kernelParams, 0.001)

	fmt.Printf("Погрешность численного метода err(H) = |%.4f - %.4f| = %.4f\n",
		numericalSolve.H, analyticalSolve.H, math.Abs(numericalSolve.H - analyticalSolve.H))

}