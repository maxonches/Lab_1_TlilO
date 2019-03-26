package brown_robinson

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"sort"
)

type BrownRobinsonTable struct {
	payoffMatrix	*mat.Dense
	steps			[]*BrownRobinsonStep
}

func NewBrownRobinsonTable(
	payoffMatrix *mat.Dense,
	firstA int,
	firstB int) *BrownRobinsonTable {
	brt := &BrownRobinsonTable{
		payoffMatrix,
		make([]*BrownRobinsonStep, 0),
	}
	brt.makeStep(firstA, firstB)

	return brt
}


func (brt *BrownRobinsonTable) Solve(threshold float64, maxSteps int) {
	for {
		nextA := brt.getNextAStrategy()
		nextB := brt.getNextBStrategy()
		step := brt.makeStep(nextA, nextB)
		shouldBreak := threshold > step.epsilon || len(brt.steps) > maxSteps
		if shouldBreak {
			break
		}
	}
	for _, step := range brt.steps {
		fmt.Println(step.String())
	}
}

func (brt *BrownRobinsonTable) getAMixedStrategy() *mat.VecDense {
	counter := make(map[int]int)
	for _, step := range brt.steps {
		counter[step.aChoice]++
	}
	keys := make([]int, 0, len(counter))
	for key := range counter {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result := make([]float64, 0, len(counter))
	for key := range keys {
		prob := float64(counter[key]) / float64(brt.getPreviousStep().stepNumber)
		result = append(result, prob)
	}
	return mat.NewVecDense(len(result), result)
}

func (brt *BrownRobinsonTable) getBMixedStrategy() *mat.VecDense {
	counter := make(map[int]int)
	for _, step := range brt.steps {
		counter[step.bChoice]++
	}
	keys := make([]int, 0, len(counter))
	for key := range counter {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result := make([]float64, 0, len(counter))
	for key := range keys {
		prob := float64(counter[key]) / float64(brt.getPreviousStep().stepNumber)
		result = append(result, prob)
	}
	return mat.NewVecDense(len(result), result)
}

func (brt *BrownRobinsonTable) getPreviousStep() *BrownRobinsonStep {
	return brt.steps[len(brt.steps) - 1]
}

func (brt *BrownRobinsonTable) getNextAStrategy() int {
	prevStep := brt.getPreviousStep()
	vecLen := prevStep.aGain.Len()
	max := prevStep.aGain.AtVec(0)
	idxOfMax := 0
	for i := 0; i < vecLen; i++ {
		elem := prevStep.aGain.AtVec(i)
		if elem > max {
			max = elem
			idxOfMax = i
		}
	}
	return idxOfMax
}

func (brt *BrownRobinsonTable) getNextBStrategy() int {
	prevStep := brt.getPreviousStep()
	vecLen := prevStep.bGain.Len()
	min := prevStep.bGain.AtVec(0)
	idxOfMin := 0
	for i := 0; i < vecLen; i++ {
		elem := prevStep.bGain.AtVec(i)
		if elem < min {
			min = elem
			idxOfMin = i
		}
	}
	return idxOfMin
}

func (brt *BrownRobinsonTable) makeStep(aStrategy int, bStrategy int) *BrownRobinsonStep {
	var prevStep *BrownRobinsonStep
	if len(brt.steps) != 0 {
		prevStep = brt.steps[len(brt.steps)-1]
	}
	step := NewStep(
		aStrategy,
		bStrategy,
		brt.payoffMatrix,
		prevStep,
	)
	brt.steps = append(brt.steps, step)
	return step
}