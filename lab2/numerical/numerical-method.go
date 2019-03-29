package numerical

import (
	"../analytical"
	an "../../lab1/analytical"
	br "../../lab1/brown-robinson"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func SolveNumerical(kernelStruct *analytical.KernelStruct, threshold float64) *analytical.AnalyticalSolveStruct {
	var prevNumericalSolve *analytical.AnalyticalSolveStruct
	n := 2
	for {
		fmt.Printf("N=%d\n", n)
		numericalSolve := &analytical.AnalyticalSolveStruct{}
		m, _ := MakeGridMatrix(n, kernelStruct)
		an.MatPrint(m)
		rPos, cPos, maximin := GetMaximin(m)
		_, _, minimax := GetMinimax(m)
		if maximin == minimax {
			numericalSolve.H = maximin
			numericalSolve.X = float64(rPos) / float64(n)
			numericalSolve.Y = float64(cPos) / float64(n)

			fmt.Println("Есть седловая точка:")
		} else {
			fmt.Println("Седловой точки нет, решение методом Брауна-Робинсона:")
			brs := br.NewBrownRobinsonSolver(m)
			brs.Solve(0.1)
			numericalSolve.H = brs.CalculateGameCost()
			mixedStrA := brs.CalculateStrategyA()
			mixedStrB := brs.CalculateStrategyB()
			strNumA, _ := GetMaxFromVector(mixedStrA)
			strNumB, _ := GetMaxFromVector(mixedStrB)
			numericalSolve.X = float64(strNumA) / float64(n)
			numericalSolve.Y = float64(strNumB) / float64(n)

		}
		fmt.Printf("X=%.3f, Y=%.3f, H=%.3f\n\n", numericalSolve.X, numericalSolve.Y, numericalSolve.H)

		if prevNumericalSolve != nil {
			if math.Abs(prevNumericalSolve.H - numericalSolve.H) < threshold {
				fmt.Printf("На итерации N=%d значение H по сравнению с предыдущей итерацией изменилось менее чем на %.3f\n", n, threshold)
				fmt.Println("Конец цикла\n")
				return numericalSolve
			}
		}
		prevNumericalSolve = numericalSolve
		n++
	}
}

func MakeGridMatrix(n int, kernelStruct *analytical.KernelStruct) (m *mat.Dense, err error) {
	if n <= 0 {
		return m, err
	}

	matElems := make([]float64, 0, (n+1)*(n+1))

	for i := 0; i < n+1; i++ {
		x := float64(i) / float64(n)
		for j := 0; j < n+1; j++ {
			y := float64(j) / float64(n)
			h, _ := analytical.Hxy(x, y, kernelStruct)
			matElems = append(matElems, h)
		}
	}
	m = mat.NewDense(n+1, n+1, matElems)

	return m, nil
}

func GetMaximin(m *mat.Dense) (int, int, float64) {
	r, _ := m.Dims()
	minimums := make([]float64, 0, r)
	cPositions := make([]int, 0, r)
	for i := 0; i < r; i++ {
		row := m.RowView(i)
		cPos, min := GetMinFromVector(row)
		minimums = append(minimums, min)
		cPositions = append(cPositions, cPos)
	}

	maximin := minimums[0]
	rPos := 0
	for i, min := range minimums {
		if min > maximin {
			maximin = min
			rPos = i
		}
	}
	cPos := cPositions[rPos]

	return rPos, cPos, maximin
}

func GetMinimax(m *mat.Dense) (int, int, float64) {
	_, c := m.Dims()
	maximums := make([]float64, 0, c)
	rPositions := make([]int, 0, c)
	for i := 0; i < c; i++ {
		col := m.ColView(i)
		rPos, max := GetMaxFromVector(col)
		maximums = append(maximums, max)
		rPositions = append(rPositions, rPos)
	}

	minimax := maximums[0]
	cPos := 0
	for i, max := range maximums {
		if max < minimax {
			minimax = max
			cPos = i
		}
	}
	rPos := rPositions[cPos]

	return rPos, cPos, minimax
}

func GetMinFromVector(v mat.Vector) (int, float64) {
	min := v.AtVec(0)
	idx := 0
	for i := 0; i < v.Len(); i++ {
		elem := v.AtVec(i)
		if elem < min {
			min = elem
			idx = i
		}
	}

	return idx, min
}

func GetMaxFromVector(v mat.Vector) (int, float64) {
	max := v.AtVec(0)
	idx := 0
	for i := 0; i < v.Len(); i++ {
		elem := v.AtVec(i)
		if elem > max {
			max = elem
			idx = i
		}
	}

	return idx, max
}
