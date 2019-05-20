package service

import (
	"bytes"
	"fmt"
	"github.com/disiqueira/gotree"
	"github.com/fatih/color"
	"math/rand"
)

type Joint struct {
	Index    int
	Player   int
	Parent   *Joint
	Children []*Joint
	Wins     [][]int
}

type TreeRandomGame struct {
	Deep        int
	Root        *Joint
	RangeOfWins Couple
	PlayersNum  int
	Strategies  []int
	colors      []ColorFunc
}

type ReverseInductionSolver struct {
	tree	*TreeRandomGame
}

type Couple struct {
	a, b int
}

type ColorFunc func(a ...interface{}) string

func NewJoint(player int, parent *Joint) *Joint {
	joint := &Joint{
		Player:   player,
		Parent:   parent,
		Children: make([]*Joint, 0),
		Wins:     make([][]int, 0),
	}
	if parent != nil {
		parent.Children = append(parent.Children, joint)
	}

	return joint
}

func (joint *Joint) String() string {
	var bb bytes.Buffer
	bb.WriteString(fmt.Sprintf("index = %d, player = %d", joint.Index, joint.Player))
	if len(joint.Wins) > 0 {
		bb.WriteString(fmt.Sprintf(", %v", joint.Wins))
	}

	return bb.String()
}

func (joint *Joint) WinContain(win []int) bool {
	if win == nil {
		return false
	}
	for _, jointWin := range joint.Wins {
		if len(win) != len(jointWin) {
			continue
		}
		match := true
		for i := range jointWin {
			if jointWin[i] != win[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}

func NewTree() *TreeRandomGame {
	colors := []ColorFunc{
		color.New(color.FgYellow).SprintFunc(),
		color.New(color.FgGreen).SprintFunc(),
		color.New(color.FgBlue).SprintFunc(),
		color.New(color.FgCyan).SprintFunc(),
	}
	tree := &TreeRandomGame{
		Deep:        7,
		Root:        nil,
		RangeOfWins: Couple{0,16},
		PlayersNum:  3,
		Strategies:  []int{2,2,2},
		colors:      colors,
	}
	GenerateRandomTree(tree)

	return tree
}

func GenerateRandomTree(tree *TreeRandomGame) {
	tree.Root = NewJoint(PlayerByLayer(0, tree), nil)
	tree.Root.Index = 0
	index := 1
	for i := 1; i < tree.Deep; i++ {
		jointsFromLevelI := tree.GetJointFromLayer(i-1)
		for _, joints := range jointsFromLevelI {
			for _, joint := range joints {
				for newChild := range tree.GenerateChildrenForLayer(i) {
					newChild.Index = index
					index++
					newChild.Parent = joint
					joint.Children = append(joint.Children, newChild)
				}
			}
		}
	}
}

func (tree *TreeRandomGame) GetJointFromLayer(layer int) [][]*Joint {
	if layer == 0 {
		return [][]*Joint{
			{tree.Root},
		}
	}
	visitedJoints := make(map[*Joint]bool)
	tmpJoint := tree.Root
	tmpDeep := 0
	jointsFromLayer := make([][]*Joint, 0)

	for tmpJoint != nil {
		curJoint := tmpJoint
		for _, child := range tmpJoint.Children {
			if _, visited := visitedJoints[child]; !visited {
				tmpJoint = child
				tmpDeep++
				break
			}
		}
		if curJoint == tmpJoint {
			if tmpDeep == layer - 1 {
				jointsFromLayer = append(jointsFromLayer, tmpJoint.Children)
			}
			visitedJoints[tmpJoint] = true
			tmpDeep--
			tmpJoint = tmpJoint.Parent
		}
	}

	return jointsFromLayer
}

func (tree *TreeRandomGame) GenerateChildrenForLayer(layer int) <-chan *Joint {
	player := PlayerByLayer(layer, tree)
	ch := make(chan *Joint)
	go func() {
		for i := 0; i < tree.Strategies[player - 1]; i++ {
			var joint *Joint
			if layer + 1 == tree.Deep {
				joint = NewJoint(0, nil)
				wins := make([]int, tree.PlayersNum)
				for i := range wins {
					wins[i] = GenerateRandomInt(tree.RangeOfWins.a, tree.RangeOfWins.b)
				}
				joint.Wins = append(joint.Wins, wins)
			} else {
				joint = NewJoint(player, nil)
			}
			ch <- joint
		}
		close(ch)
	}()

	return ch
}

func InverseInductionStep(layer int, tree *TreeRandomGame) {
	jointsFromLayer := tree.GetJointFromLayer(layer)
	player := PlayerByLayer(layer, tree)

	for _, joints := range jointsFromLayer {
		for _, joint := range joints {
			strategiesToComp := make([][]int, 0, len(joints))
			for _, child := range joint.Children {
				for _, win := range child.Wins {
					strategiesToComp = append(strategiesToComp, win)
				}
			}
			if len(strategiesToComp) == 0 {
				continue
			}
			maxValueForPlayer := strategiesToComp[0][player-1]
			for _, strategy := range strategiesToComp {
				if strategy[player-1] > maxValueForPlayer {
					maxValueForPlayer = strategy[player-1]
				}
			}
			for _, strategy := range strategiesToComp {
				if strategy[player-1] == maxValueForPlayer {
					joint.Wins = append(joint.Wins, strategy)
				}
			}
		}
	}
}

func PlayerByLayer(layer int, tree *TreeRandomGame) int {
	return (layer % tree.PlayersNum) + 1
}

func PrintTree(tree *TreeRandomGame) string {
	return MakeTreeFromJoint(tree.Root, tree.colors, tree.Root.Wins).Print()
}

func MakeTreeFromJoint(joint *Joint, colors []ColorFunc, rootWins [][]int) gotree.Tree {
	var printJoint gotree.Tree = nil
	var sprintFunc ColorFunc = nil
	for idx, win := range rootWins {
		if joint.WinContain(win) {
			sprintFunc = colors[idx]
			break
		}
	}
	if sprintFunc != nil {
		printJoint = gotree.New(sprintFunc(joint.String()))
	} else {
		printJoint = gotree.New(joint.String())
	}
	for _, child := range joint.Children {
		printJoint.AddTree(MakeTreeFromJoint(child, colors, rootWins))
	}

	return printJoint
}

func NewSolver(tree *TreeRandomGame) *ReverseInductionSolver {
	return &ReverseInductionSolver{
		tree: tree,
	}
}

func Solve(ris *ReverseInductionSolver) {
	for i := ris.tree.Deep -1; i >=0; i-- {
		InverseInductionStep(i, ris.tree)
	}
}

func GenerateRandomInt(min, max int) int {
	return rand.Intn(max - min) + min
}