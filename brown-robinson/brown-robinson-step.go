package brown_robinson

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type BrownRobinsonStep struct {
	stepNumber			int
	aChoice				int
	bChoice				int
	aGain				*mat.VecDense
	bGain				*mat.VecDense
	curUpperGameCost	float64
	curLowerGameCost	float64
	minUpperGameCost	float64
	maxLowerGameCost	float64
	epsilon				float64
}


func NewStep(
	aChoice int,
	bChoice int,
	payoffMatrix *mat.Dense,
	prevStep *BrownRobinsonStep) *BrownRobinsonStep {
	var stepNum int
	if prevStep != nil {
		stepNum = prevStep.stepNumber + 1
	} else {
		stepNum = 1
	}
	currentAGain := payoffMatrix.ColView(bChoice)
	currentBGain := payoffMatrix.RowView(aChoice)


	aGain := mat.NewVecDense(currentAGain.Len(), nil)
	bGain := mat.NewVecDense(currentBGain.Len(), nil)
	var curUpperGameCost float64
	var curLowerGameCost float64
	var minUpperGameCost float64
	var maxLowerGameCost float64
	if prevStep != nil {
		aGain.AddVec(prevStep.aGain, currentAGain)
		bGain.AddVec(prevStep.bGain, currentBGain)
		curUpperGameCost = mat.Max(aGain) / float64(stepNum)
		curLowerGameCost = mat.Min(bGain) / float64(stepNum)
		minUpperGameCost = Min([]float64{curUpperGameCost, prevStep.minUpperGameCost})
		maxLowerGameCost = Max([]float64{curLowerGameCost, prevStep.maxLowerGameCost})
	} else {
		aGain.CopyVec(currentAGain)
		bGain.CopyVec(currentBGain)
		curUpperGameCost = mat.Max(aGain) / float64(stepNum)
		curLowerGameCost = mat.Min(bGain) / float64(stepNum)
		minUpperGameCost = curUpperGameCost
		maxLowerGameCost = curLowerGameCost
	}
	epsilon := minUpperGameCost - maxLowerGameCost

	return &BrownRobinsonStep{
		stepNum,
		aChoice,
		bChoice,
		aGain,
		bGain,
		curUpperGameCost,
		curLowerGameCost,
		minUpperGameCost,
		maxLowerGameCost,
		epsilon,
	}
}

func (s *BrownRobinsonStep) String() string {
	return fmt.Sprintf("%d | %d | %d | %v | %v | %.3f | %.3f | %.3f",
		s.stepNumber, s.aChoice + 1, s.bChoice + 1, s.aGain.RawVector().Data, s.bGain.RawVector().Data,
		s.curUpperGameCost, s.curLowerGameCost, s.epsilon)
}