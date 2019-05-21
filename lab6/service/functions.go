package service

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
	"sort"
	"strings"
)

type Game struct {
	MatrixA          *mat.Dense
	VectorX          *mat.VecDense
	Dimension        int
	RangeOfInfluence pair
	Eps              float64
}

type pair struct {
	min, max int
}

func NewGame(dim int) *Game {
	randomMatrix := GenerateRandomStochasticTrustMatrix(dim)
	initialOpinions := GenerateRandomVector(0, 20, dim)

	return &Game{
		MatrixA:          randomMatrix,
		VectorX:          initialOpinions,
		Dimension:        dim,
		RangeOfInfluence: pair{
			0,
			100,
	},
		Eps:              10e-6,
	}
}

func (g *Game) MakeGame() {
	Xt, t := g.multByAUntilAccuracyReached(g.VectorX)
	//AInf := computeAInf(g.MatrixA, t)
	fmt.Println("Initial agent opinions:")
	PrintMatrix(g.VectorX.T(), "Vx(0)")
	fmt.Printf("Iterations: %d\n", t)
	fmt.Println("Resulting agent opinion without influence:")
	PrintMatrix(Xt.T(), "Vx(t->inf)")
	//PrintMatrix(g.MatrixA, "A")
	//PrintMatrix(AInf, "AInf")

}

func (g *Game) MakeGameInfluence() {
	agents := make([]int, g.Dimension)
	for i := 0; i < g.Dimension; i++ {
		agents[i] = i
	}
	rand.Shuffle(len(agents), func(i, j int) {
		agents[i],
		agents[j] = agents[j],
		agents[i]
		},
	)

	uSize, vSize := len(agents), len(agents)

	for uSize + vSize > len(agents) {
		uSize = GenerateRandomInt(1, len(agents))
		vSize = GenerateRandomInt(1, len(agents))
	}

	uAgents := agents[:uSize]
	vAgents := agents[uSize:uSize+vSize]
	sort.Ints(uAgents)
	sort.Ints(vAgents)
	fmt.Printf("Agents of the first player: %v, Agents of the second player: %v\n", uAgents, vAgents)

	uInflValue := float64(GenerateRandomInt(g.RangeOfInfluence.min, g.RangeOfInfluence.max))
	vInflValue := -float64(GenerateRandomInt(g.RangeOfInfluence.min, g.RangeOfInfluence.max))

	var xInfl = mat.VecDenseCopyOf(g.VectorX)
	for _, agentNum := range uAgents {
		xInfl.SetVec(agentNum, uInflValue)
	}
	for _, agentNum := range vAgents {
		xInfl.SetVec(agentNum, vInflValue)
	}
	fmt.Printf("Formed initial opinion of the first player: %.0f\n", uInflValue)
	fmt.Printf("Formed initial opinion of the second player: %.0f\n", vInflValue)
	fmt.Printf("Initial opinions with regard to the formed:\n")
	PrintMatrix(xInfl.T(), "Vx(0)")
	Xt, t := g.multByAUntilAccuracyReached(xInfl)
	fmt.Printf("Iterations: %d\n", t)
	fmt.Println("Resulting opinion:")
	PrintMatrix(Xt.T(), "Vx(t->inf)")
}

func computeAInf(A *mat.Dense, itersNum int) (*mat.Dense) {
	AInf := mat.Dense{}
	AInf.Clone(A)
	for i := 0; i < itersNum; i++ {
		AInf.Mul(&AInf, A)
	}
	return &AInf
}

func (g *Game) multByAUntilAccuracyReached(x *mat.VecDense) (*mat.Dense, int) {
	t := 0
	var prevXt, curXt, deltaXt mat.Dense
	var accuracyReached bool

	curXt.Clone(x)
	for {
		t++
		accuracyReached = true
		prevXt.Clone(&curXt)
		curXt.Mul(g.MatrixA, &curXt)
		deltaXt.Sub(&curXt, &prevXt)

		deltaXt.Apply(func(i, j int, value float64) float64 {
			if math.Abs(value) >= g.Eps {
				accuracyReached = false
			}
			return value
		}, &deltaXt)
		if accuracyReached {
			break
		}
	}

	return &curXt, t
}

func GenerateRandomStochasticTrustMatrix(dim int) *mat.Dense {
	data := make([]float64, 0, dim * dim)
	for i := 0; i < dim; i++ {
		row := GenerateStochasticRow(dim)
		data = append(data, row...)
	}

	return mat.NewDense(dim, dim, data)
}

func GenerateStochasticRow(size int) []float64 {
	row := make([]float64, 0, size)
	sum := 0.0
	for i := 0; i < size; i++ {
		value := rand.Float64()
		sum += value
		row = append(row, value)
	}
	for i := range row {
		row[i] /= sum
	}

	return row
}

func GenerateRandomVector(min, max, size int) *mat.VecDense {
	data := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, float64(GenerateRandomInt(min, max)))
	}

	return mat.NewVecDense(size, data)
}

func GenerateRandomInt(min, max int) int {
	return rand.Intn(max - min) + min
}

func PrintMatrix(matrix mat.Matrix, name string) {
	fmt.Printf("%s = % .2f\n", name, mat.Formatted(matrix, mat.Prefix(strings.Repeat(" ", len(name) + 3)), mat.DotByte('0')))
}