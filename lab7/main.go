package main

import (
	"devprojects/labs_TlilO/lab7/service"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(102)
	g := service.NewGame(10)
	g.MakeGame()
	fmt.Println()
	g.MakeGameInfluence()
}