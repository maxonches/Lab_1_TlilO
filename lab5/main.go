package main

import (
	"devprojects/labs_TlilO/lab5/service"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(101)

	t := service.NewTree()
	fmt.Printf("Random position game (%d):\n Located in file 'RandomPosGame.txt'\n", t.Deep)
	fmt.Println(service.PrintTree(t))
	f0, err := os.Create("RandomPosGame.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f0.WriteString(service.PrintTree(t))
	if err != nil {
		fmt.Println(err)
		f0.Close()
		return
	}

	fmt.Println("Solution by method of inverse induction:\n Located in file 'InverseInduction.txt'")
	ris := service.NewSolver(t)
	service.Solve(ris)
	fmt.Println(service.PrintTree(t))
	f1, err := os.Create("InverseInduction.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f1.WriteString(service.PrintTree(t))
	if err != nil {
		fmt.Println(err)
		f1.Close()
		return
	}
}