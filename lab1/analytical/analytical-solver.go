package analytical

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func SolveAnalyticalMethod(matrix *mat.Dense) (x, y *mat.Dense, v float64) {
	// инвертируем матрицу
	var inverseMatrix mat.Dense
	inverseMatrix.Inverse(matrix)
	formatMatrix := mat.Formatted(&inverseMatrix, mat.Prefix(""), mat.Squeeze())
	fmt.Println("ANALYTICAL METHOD")
	fmt.Printf("Inverse matrix:\n%.3f\n\n", formatMatrix)

	// находим вектор вероятностей y*
	u := mat.NewVecDense(3, []float64{
		1, 1, 1,
	})
	numeratorY := mat.NewDense(3, 1, nil)
	numeratorY.Product(&inverseMatrix, u)
	denominatorY := mat.NewDense(1, 1, nil)
	denominatorY.Product(u.T(), numeratorY)
	y = mat.NewDense(3,1,nil)
	y.Scale(1 / denominatorY.At(0,0), numeratorY)

	// находим вектор вероятностей x*
	numeratorX := mat.NewDense(1, 3, nil)
	numeratorX.Product(u.T(), &inverseMatrix)
	denominatorX := mat.NewDense(1, 1, nil)
	denominatorX.Product(u.T(), numeratorY)
	x = mat.NewDense(1,3,nil)
	x.Scale(1 / denominatorX.At(0,0), numeratorX)

	// находим цену игры v
	v = 1 / denominatorX.At(0,0)

	r, c := y.Dims()
	y = mat.NewDense(c, r, y.RawMatrix().Data)

	return x, y, v
}

func MatPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%.3f\n", fa)
}