package analytical

type KernelStruct struct  {
	A 	float64
	B 	float64
	C 	float64
	D 	float64
	E 	float64
}

type AnalyticalSolveStruct struct {
	KernelStruct	KernelStruct
	X				float64
	Y				float64
	H				float64
}

func Hxy(x, y float64, kernelStruct *KernelStruct) (kernelFunction float64, err error) {
	kernelFunction = kernelStruct.A * x * x + kernelStruct.B * y * y + kernelStruct.C * x * y + kernelStruct.D * x + kernelStruct.E * y

	return kernelFunction, err
}

func AnalyticalSolverFunc(kernelStruct KernelStruct) (analyticalSolve *AnalyticalSolveStruct, err error) {
	analyticalSolve = &AnalyticalSolveStruct{}
	analyticalSolve.KernelStruct = kernelStruct

	return analyticalSolve, err
}

func Solution(analyticalSolve *AnalyticalSolveStruct) (err error) {
	err = ConvexConcaveChecking(analyticalSolve)
	if err != nil {
		return err
	}
	x, y, err := SaddlePoint(&analyticalSolve.KernelStruct)
	if err != nil {
		return err
	}

	analyticalSolve.X = x
	analyticalSolve.Y = y
	analyticalSolve.H, err = Hxy(x, y, &analyticalSolve.KernelStruct)
	if err != nil {
		return err
	}

	return err
}

func SaddlePoint(kernelStruct *KernelStruct) (x float64, y float64, err error) {
	x = (kernelStruct.E*kernelStruct.C -2*kernelStruct.B*kernelStruct.D)/(4*kernelStruct.A*kernelStruct.B -kernelStruct.C*kernelStruct.C)
	y = -(kernelStruct.C*x+kernelStruct.E)/(2*kernelStruct.B)

	return x, y, err
}

func ConvexConcaveChecking(kernelStruct *AnalyticalSolveStruct) (err error) {
	Hxx := 2 * kernelStruct.KernelStruct.A
	Hyy := 2 * kernelStruct.KernelStruct.B

	if Hxx < 0 && Hyy > 0 {
		return nil
	} else {
		return err
	}

	return err
}
