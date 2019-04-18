package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

const BOUNDARIES = 50
const MATRIX_SIZE = 10
const SET = "\u001B[0m"
const COLOR_RED = "\u001B[31m"
const COLOR_GREEN = "\u001B[32m"
const COLOR_BLUE = "\u001B[34m"

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
	battleOfMalesBiMatrix := [][][]float64{
		{
			{4, 1}, {0, 0},
		},
		{
			{0, 0}, {1, 4},
		},
	}
	prisonerDilemmaBiMatrix := [][][]float64{
		{
			{-5, -5}, {0, -10},
		},
		{
			{-10, 0}, {-1, -1},
		},
	}

	fmt.Println(fmt.Sprintf("%v%v%v", COLOR_RED, "Равновесие по Нэшу", SET))
	fmt.Println(fmt.Sprintf("%v%v%v", COLOR_BLUE, "Парето эффективность", SET))
	fmt.Println(fmt.Sprintf("%v%v%v", COLOR_GREEN, "Все вместе", SET))
	fmt.Println("Перекресток:")
	SearchAllOptimalSituations(crossroadBiMatrix)
	fmt.Println("Семейный спор:")
	SearchAllOptimalSituations(battleOfMalesBiMatrix)
	fmt.Println("Дилемма заключенного:")
	SearchAllOptimalSituations(prisonerDilemmaBiMatrix)
	inputMatrix := GenerateBiMatrix(MATRIX_SIZE, BOUNDARIES)
	fmt.Println()
	fmt.Println("Случайно сгенерированная матрица 10х10:")
	SearchAllOptimalSituations(inputMatrix)
	fmt.Println("Матрица по варианту:")
	SearchAllOptimalSituations(inputBiMatrix)
	SolveAnalitycal(MATRIX_SIZE, inputBiMatrix)
}

func CheckEquilibriumByNash(matrix [][][]float64, i int, j int) bool {
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

func СheckParetoOptimality(matrix [][][]float64, i int, j int) (bool) {
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

func SearchAllOptimalSituations(inputMatrix [][][]float64) {
	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[0]); j++ {
			isNashEfficient := CheckEquilibriumByNash(inputMatrix, i, j)
			isParetoOptimal := СheckParetoOptimality(inputMatrix, i, j)
			result := fmt.Sprintf("%v%v%v%v%v", "(", inputMatrix[i][j][0], ";", inputMatrix[i][j][1], ")")
			if isNashEfficient && isParetoOptimal {
				result = fmt.Sprintf("%v%v%v", COLOR_GREEN, result, SET)
			}
			if isNashEfficient && !isParetoOptimal {
				result = fmt.Sprintf("%v%v%v", COLOR_RED, result, SET)
			}
			if !isNashEfficient && isParetoOptimal {
				result = fmt.Sprintf("%v%v%v", COLOR_BLUE, result, SET)
			}
			fmt.Print(fmt.Sprintf("%v%v", result, " "))
		}
		fmt.Println()
	}
}

func GenerateBiMatrix(size int, boundaries float64) ([][][]float64) {
	rand.Seed(time.Now().Unix())
	size = 10
	boundaries = 50.0
	min := -1 * boundaries
	max := 50.0
	matrix := make([][][]float64, 0, size)
	for i := 0; i < size; i++ {
		matrixOneDimensional := make([][]float64, 0, size)
		for j := 0; j < size; j++ {
			matrixTwoDimensional := make([]float64, 0, 2)
			for k := 0; k < 2; k++ {
				matrixTwoDimensional = append(matrixTwoDimensional, float64(rand.Intn(int(max)-int(min))+int(min)))
			}
			matrixOneDimensional = append(matrixOneDimensional, matrixTwoDimensional)
		}
		matrix = append(matrix, matrixOneDimensional)
	}

	return matrix
}

func SolveAnalitycal(size int, inputBiMatrix [][][]float64) {
	matrixAOneDim := make([]float64, 0, size*size)
	matrixBOneDim := make([]float64, 0, size*size)
	for i := 0; i < len(inputBiMatrix); i++ {
		for j := 0; j < len(inputBiMatrix); j++ {
			matrixAOneDim = append(matrixAOneDim, inputBiMatrix[i][j][0])
			matrixBOneDim = append(matrixBOneDim, inputBiMatrix[i][j][1])
		}
	}
	matrixA := mat.NewDense(len(inputBiMatrix), len(inputBiMatrix), matrixAOneDim)
	matrixB := mat.NewDense(len(inputBiMatrix), len(inputBiMatrix), matrixBOneDim)

	uRow := make([]float64, len(inputBiMatrix))
	for i := 0; i < len(inputBiMatrix); i++ {
		uRow[i] = 1
	}
	u := mat.NewVecDense(len(inputBiMatrix), uRow)

	var inverseMatrixA, inverseMatrixB mat.Dense
	inverseMatrixA.Inverse(matrixA)
	inverseMatrixB.Inverse(matrixB)

	uAinv := mat.Dense{}
	uAinv.Product(u.T(), &inverseMatrixA)
	uAinvUt := mat.Dense{}
	uAinvUt.Product(&uAinv, u)
	v1 := 1 / uAinvUt.At(0,0)

	uBinv := mat.Dense{}
	uBinv.Product(u.T(), &inverseMatrixB)
	uBinvUt := mat.Dense{}
	uBinvUt.Product(&uBinv, u)
	v2 := 1 / uBinvUt.At(0,0)

	fmt.Println(fmt.Sprintf("%v%.3f", "v1 = ", v1))
	fmt.Println(fmt.Sprintf("%v%.3f", "v2 = ", v2))

	x := mat.Dense{}
	x.Product(u.T(), &inverseMatrixB)
	x.Scale(v2, &x)

	y := mat.Dense{}
	y.Product(&inverseMatrixA, u)
	y.Scale(v1, &y)

	fmt.Println("x = ")
	MatPrint(&x)
	fmt.Println("y = ")
	MatPrint(y.T())
}

func MatPrint(x mat.Matrix) {
	fa := mat.Formatted(x, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%.3f\n", fa)
}