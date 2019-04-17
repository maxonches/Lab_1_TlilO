package main

import (
	"fmt"
	"github.com/dglo/java2go/parser"
	"go/format"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
	"time"
)

const bOUNDARIES = 50
const mATRIX_SIZE = 10
const aNSI_RESET = "\u001B[0m"
const aNSI_RED = "\u001B[31m"
const aNSI_GREEN = "\u001B[32m"
const aNSI_BLUE = "\u001B[34m"

type LR9Main struct {

}

func main() {
	inputBiMatrix := [][][]float64{
		{
			{4, 7}, {5, 2},
		},
		{
			{0, 2}, {7, 3},
		},
	}
	crossroadBiMatrix := [][][]float64{
		{
			{1, 1}, {0.5, 2},
		},
		{
			{2, 0.5}, {0, 0},
		},
	}
	battleOFMalesBiMatrix := [][][]float64{
		{
			{4, 1}, {0, 0},
		},
		{
			{0, 0}, {1, 4},
		},
	}
	pRISONER_DILEMMA_BI_MATRIX := [][][]float64{
		{
			{-5, -5}, {0, -10},
		},
		{
			{-10, 0}, {-1, -1},
		},
	}

	fmt.Println(fmt.Sprintf("%v%v%v", aNSI_RED, "Равновесие по Нэшу", aNSI_RESET))
	fmt.Println(fmt.Sprintf("%v%v%v", aNSI_BLUE, "Парето эффективность", aNSI_RESET))
	fmt.Println(fmt.Sprintf("%v%v%v", aNSI_GREEN, "Все вместе", aNSI_RESET))
	fmt.Println("Перекресток:")
	findAllOptimalSituations(crossroadBiMatrix)
	fmt.Println("Семейный спор:")
	findAllOptimalSituations(battleOFMalesBiMatrix)
	fmt.Println("Дилемма заключенного:")
	findAllOptimalSituations(pRISONER_DILEMMA_BI_MATRIX)
	inputBiMatrix := generateBiMatrix(mATRIX_SIZE, bOUNDARIES)
	fmt.Println()
	fmt.Println("Случайно сгенерированная матрица 10х10:")
	findAllOptimalSituations(inputBiMatrix)
	fmt.Println("Матрица по варианту")
	findAllOptimalSituations(inputBiMatrix)
	solveAnalitycally(inputBiMatrix)
}

func NewLR9Main() (rcvr *LR9Main) {
	rcvr = &LR9Main{}
	return
}
func checkForNashOptimality(matrix [][][]float64, i int, j int) bool {
	isThereBetterStrategyForA := false
	isThereBetterStrategyForB := false
	for jIter := 0; jIter < len(matrix); jIter++ {
		if matrix[i][jIter][1] > matrix[i][j][1] {
			isThereBetterStrategyForA = true
		}
	}
	for iter := 0; iter < len(matrix); iter++ {
		if matrix[iter][j][0] > matrix[i][j][0] {
			isThereBetterStrategyForB = true
		}
	}
	return !(isThereBetterStrategyForA || isThereBetterStrategyForB)
}
func checkForParretoEfficiency(matrix [][][]float64, i int, j int) (bool) {
	isThereBetterStrategy := false
	for iter := 0; iter < len(matrix); iter++ {
		for jIter := 0; jIter < len(matrix); jIter++ {
			if (matrix[iter][jIter][0] > matrix[i][j][0] && matrix[iter][jIter][1] >= matrix[i][j][1]) ||
				(matrix[iter][jIter][1] > matrix[i][j][1] && matrix[iter][jIter][0] >= matrix[i][j][0]) {
				isThereBetterStrategy = true
			}
		}
	}

	return !isThereBetterStrategy
}
func findAllOptimalSituations(inputMatrix [][][]float64) {
	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[0]); j++ {
			isNashEfficient := checkForNashOptimality(inputMatrix, i, j)
			isParetoOptimal := checkForParretoEfficiency(inputMatrix, i, j)
			result := fmt.Sprintf("%v%v%v%v%v", "(", inputMatrix[i][j][0], ";", inputMatrix[i][j][1], ")")
			if isNashEfficient && isParetoOptimal {
				result = fmt.Sprintf("%v%v%v", aNSI_GREEN, result, aNSI_RESET)
			}
			if isNashEfficient && !isParetoOptimal {
				result = fmt.Sprintf("%v%v%v", aNSI_RED, result, aNSI_RESET)
			}
			if !isNashEfficient && isParetoOptimal {
				result = fmt.Sprintf("%v%v%v", aNSI_BLUE, result, aNSI_RESET)
			}
			fmt.Print(fmt.Sprintf("%v%v", result, " "))
		}
		fmt.Println()
	}
}

func generateBiMatrix(size int, boundaries float64) ([][][]float64) {
	rand.Seed(time.Now().Unix())
	min := -1 * boundaries
	max := boundaries
	matrix := make([][][]float64, size)
	for i := 0; i < len(matrix); i++ {
		matrixOneDimensional := make([][]float64, size)
		for j := 0; j < len(matrix); j++ {
			matrixTwoDimensional := make([]float64, 2)
			matrix[i][j][0] = rand.Float64() * (max - min) + min
			matrix[i][j][1] = rand.Float64() * (max - min) + min
		}
		matrix = append(matrix, matrixOneDimensional)
	}

	return matrix
}

func generateBiMatrix(size int, boundaries float64) ([][][]float64) {
	rand.Seed(time.Now().Unix())
	min := -1 * boundaries
	max := boundaries
	matrix := make([][][]float64, size)
	for i := 0; i < len(matrix); i++ {
		matrixOneDimensional := make([][]float64, size)
		for j := 0; j < len(matrix); j++ {
			matrixTwoDimensional := make([]float64, 2)
			matrix[i][j][0] = rand.Float64() * (max - min) + min
			matrix[i][j][1] = rand.Float64() * (max - min) + min
		}
		matrix = append(matrix, matrixOneDimensional)
	}

	return matrix
}

func GENERATE() {
	size := 10
	boundaries := 50
	rand.Seed(time.Now().Unix())
	matrix := make([][][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([][]float64, size + 1)
		for j := 0; j < size + 1; j++ {
			matrix[i][j] = make([]float64, boundaries)
			for k := 0; k < int(boundaries); k++ {
				matrix[i][j][k] = 1
			}
		}
	}
	fmt.Println(matrix)
}

func solveAnalitycally(matrix [][][]float64) {
	matrixA := make([]float64, len(matrix), len(matrix))
	matrixB := make([]float64, len(matrix), len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrixA[i][j] = matrix[i][j][0]
			matrixB[i][j] = matrix[i][j][1]
		}
	}
	uRow := make([]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		uRow[i] = 1
	}
	u := []float64{uRow}
	u_t := make([]float64, len(matrix), 1)
	for i := 0; i < len(matrix); i++ {
		u_t[i][0] = 1
	}
	aReverse := mat.calculateReverseMatrix(matrixA)
	bReverse := mat.calculateReverseMatrix(matrixB)
	aReverseUT := mat.multiplicate(aReverse, u_t)
	bReverseUT := mat.multiplicate(bReverse, u_t)
	vX := 1 / mat.multiplicate(u, aReverseUT)[0][0]
	vY := 1 / mat.multiplicate(u, bReverseUT)[0][0]
	format := NewDecimalFormat("#0.000")
	fmt.Println(fmt.Sprintf("%v%v", "VX:", format.format(vX)))
	fmt.Println(fmt.Sprintf("%v%v", "VY:", format.format(vY)))
	x := []float64{MatrixTools.multiplicate(aReverse, u_t)[0][0] * vX, MatrixTools.multiplicate(aReverse, u_t)[1][0] * vX}
	y := []float64{MatrixTools.multiplicate(u, bReverse)[0][0] * vY, MatrixTools.multiplicate(u, bReverse)[0][1] * vY}
	fmt.Println(fmt.Sprintf("%v%v", "X:", fmt.Sprintf("%v", x)))
	fmt.Println(fmt.Sprintf("%v%v", "Y:", fmt.Sprintf("%v", y)))
}