package main

import (
	"fmt"
	"gitlab.echelon.lan/lab2/labs_TlilO/lab2/analytical"
)


func main() {
	kernelParams := &analytical.KernelStruct{-15, 6.667, 40, -12, -24}
	fmt.Println("АНАЛИТИЧЕСКОЕ РЕШЕНИЕ")
	analyticalSolve := analytical.AnalyticalSolverFunc(*kernelParams)
	err := analytical.Solution(analyticalSolve)
	if err != nil {
		fmt.Println("Error of analytical method!")
	}
	fmt.Printf("Седловая точка x=%.3f, y=%.3f, H(x,y) = %.3f\n", analyticalSolve.X, analyticalSolve.Y, analyticalSolve.H)

	fmt.Println("ЧИСЛЕННЫЙ МЕТОД")

}