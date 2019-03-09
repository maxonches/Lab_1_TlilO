package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"lab1/analytical"
	"lab1/brown-robinson"
	"math"
)

func main() {
	payMatrix := mat.NewDense(3, 3, []float64{
		17,   4,  9,
		 0,  16,  9,
		12,   2, 19,
	})

	x, y, v := analytical.SolveAnalyticalMethod(payMatrix)
	fmt.Println("Probability vector x* =")
	analytical.MatPrint(x)
	fmt.Println("Probability vector y* =")
	analytical.MatPrint(y)
	fmt.Printf("Game cost v = %.3f\n\n", v)

	fmt.Println("METHOD OF BROWN-ROBINSON")
	brownRobinsonSolver := brown_robinson.NewBrownRobinsonSolver(payMatrix)
	brownRobinsonSolver.Solve(0.1)
	gameCost := brownRobinsonSolver.CalculateGameCost()
	strategyA := brownRobinsonSolver.CalculateStrategyA()
	strategyB := brownRobinsonSolver.CalculateStrategyB()

	// погрешность для смешанной стратегии А
	inaccuracyForA := mat.NewDense(3,1, nil)
	inaccuracyForA.Sub(x.T(), strategyA)
	inaccuracyForA.Apply(func(i, j int, v float64) float64 {
		return math.Abs(v)
	}, inaccuracyForA)

	// погрешность для смешанной стратегии B
	inaccuracyForB := mat.NewDense(3,1, nil)
	inaccuracyForB.Sub(y.T(), strategyB)
	inaccuracyForB.Apply(func(i, j int, v float64) float64 {
		return math.Abs(v)
	}, inaccuracyForB)

	fmt.Printf("Game cost = %.3f\n", gameCost)
	fmt.Printf("A mixed strategy = %.3f\n", strategyA.RawVector().Data)
	fmt.Printf("B mixed strategy = %.3f\n", strategyB.RawVector().Data)
	fmt.Printf("Inaccuracy for a mixed strategy A = %.3f\n", inaccuracyForA.RawMatrix().Data)
	fmt.Printf("Inaccuracy for a mixed strategy B = %.3f\n", inaccuracyForB.RawMatrix().Data)

}

