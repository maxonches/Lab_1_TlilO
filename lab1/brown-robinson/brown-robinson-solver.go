package brown_robinson

import "gonum.org/v1/gonum/mat"

type BrownRobinsonSolver struct {
	payoffMatrix		*mat.Dense
	brownRobinsonTable	*BrownRobinsonTable
}

func NewBrownRobinsonSolver(inputMatrix *mat.Dense) *BrownRobinsonSolver {
	return &BrownRobinsonSolver{
		inputMatrix,
		nil,
	}
}

func (brs *BrownRobinsonSolver) Solve(threshold float64) {
	r, c := brs.payoffMatrix.Dims()
	brt := NewBrownRobinsonTable(
		brs.payoffMatrix,
		RandFromRange(0, r),
		RandFromRange(0, c),
	)
	brs.brownRobinsonTable = brt
	brt.Solve(threshold, 1000)
}

func (brs *BrownRobinsonSolver) CalculateStrategyA() *mat.VecDense {
	return brs.brownRobinsonTable.getAMixedStrategy()
}

func (brs *BrownRobinsonSolver) CalculateStrategyB() *mat.VecDense {
	return brs.brownRobinsonTable.getBMixedStrategy()
}

func (brs *BrownRobinsonSolver) CalculateGameCost() float64 {
	aMixedStrategy := brs.CalculateStrategyA()
	bMixedStrategy := brs.CalculateStrategyB()

	gameCost := 0.0
	for j := 0; j < bMixedStrategy.Len(); j++ {
		for i := 0; i < aMixedStrategy.Len(); i++ {
			gameCost += brs.payoffMatrix.At(i, j) * aMixedStrategy.AtVec(i) * bMixedStrategy.AtVec(j)
		}
	}

	return gameCost
}