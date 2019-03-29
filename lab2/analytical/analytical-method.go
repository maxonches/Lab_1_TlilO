package analytical

import "fmt"

type KernelStruct struct  {
	A float64
	B float64
	C float64
	D float64
	E float64
}

type AnalyticalSolveStruct struct {
	KernelStruct KernelStruct
	X            float64
	Y            float64
	H            float64
}

func Hxy(x, y float64, kernelStruct *KernelStruct) float64 {
	return kernelStruct.A* x * x + kernelStruct.B* y * y + kernelStruct.C* x * y + kernelStruct.D* x + kernelStruct.E* y
}

func AnalyticalSolverFunc(kernelStruct KernelStruct) *AnalyticalSolveStruct {
	analyticalSolve := &AnalyticalSolveStruct{}
	analyticalSolve.KernelStruct = kernelStruct

	return analyticalSolve
}

func Solution(analyticalSolve *AnalyticalSolveStruct) error {
	err := ConvexConcaveChecking(analyticalSolve)
	if err != nil {
		return err
	}
	x, y := SaddlePoint(&analyticalSolve.KernelStruct)
	if err != nil {
		return err
	}

	analyticalSolve.X = x
	analyticalSolve.Y = y
	analyticalSolve.H = Hxy(x, y, &analyticalSolve.KernelStruct)

	return nil
}

func SaddlePoint(kernelStruct *KernelStruct) (float64, float64) {
	x := (kernelStruct.E*kernelStruct.C -2*kernelStruct.B*kernelStruct.D)/(4*kernelStruct.A*kernelStruct.B -kernelStruct.C*kernelStruct.C)
	y := -(kernelStruct.C*x+kernelStruct.E)/(2*kernelStruct.B)

	return x, y
}


func ConvexConcaveChecking(kernelStruct *AnalyticalSolveStruct) error {
	Hxx := 2 * kernelStruct.KernelStruct.A
	Hyy := 2 * kernelStruct.KernelStruct.B

	if Hxx < 0 {
		return nil
	} else {
		fmt.Println("The game is not convex - concave!")
	}
	if Hyy > 0 {
		return nil
	} else {
		fmt.Println("The game is not convex - concave!")
	}

	return nil
}
