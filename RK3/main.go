package main

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"sort"
)

func main() {
	charFunc := make(map[string]int)
	charFunc[SetToString(hashset.New())] = 0
	charFunc[SetToString(hashset.New(1))] = 1
	charFunc[SetToString(hashset.New(2))] = 1
	charFunc[SetToString(hashset.New(3))] = 3
	charFunc[SetToString(hashset.New(4))] = 2
	charFunc[SetToString(hashset.New(1, 2))] = 4
	charFunc[SetToString(hashset.New(1, 3))] = 4
	charFunc[SetToString(hashset.New(1, 4))] = 5
	charFunc[SetToString(hashset.New(2, 3))] = 4
	charFunc[SetToString(hashset.New(2, 4))] = 4
	charFunc[SetToString(hashset.New(3, 4))] = 7
	charFunc[SetToString(hashset.New(1, 2, 3))] = 9
	charFunc[SetToString(hashset.New(1, 2, 4))] = 9
	charFunc[SetToString(hashset.New(1, 3, 4))] = 8
	charFunc[SetToString(hashset.New(2, 3, 4))] = 8
	charFunc[SetToString(hashset.New(1, 2, 3, 4))] = 12

	keys := []string{}
	for k := range charFunc {
		keys = append(keys, k)
	}
	sort.Sort(ByLenAndValues(keys))
	for _, k := range keys {
		fmt.Println("Key:", k, "		Value:", charFunc[k])
	}

	isSuperadditive := true
	for setString1, value1 := range charFunc {
		set1 := StringToSet(setString1)
		for setString2, value2 := range charFunc {
			set2 := StringToSet(setString2)
			if setString1 == setString2 {
				continue
			}
			intersection := Intersection(set1,set2)
			if intersection.Size() > 0 {
				continue
			}
			union := Union(set1,set2)
			unionVal := charFunc[SetToString(union)]
			if unionVal < value1 + value2 {
				fmt.Println("Not superadditive", union, unionVal, set1, value1, set2, value2)
				isSuperadditive = false
			}
		}
	}
	if isSuperadditive {
		fmt.Println("Game is superadditive")
	} else {
		fmt.Println("Game is not superadditive")
	}

	isConvex := true
	for setString1, value1 := range charFunc {
		set1 := StringToSet(setString1)
		for setString2, value2 := range charFunc {
			set2 := StringToSet(setString2)
			if setString1 == setString2 {
				continue
			}
			intersection := Intersection(set1, set2)
			union := Union(set1,set2)
			unionVal := charFunc[SetToString(union)]
			intersectionVal := charFunc[SetToString(intersection)]
			if unionVal + intersectionVal < value1 + value2 {
				fmt.Println("Not convex on ", intersection, intersectionVal, union, unionVal, set1, value1, set2, value2)
				isConvex = false
			}
		}
	}
	if isConvex {
		fmt.Println("Game is convex")
	} else {
		fmt.Println("Game is not convex")
	}

	vectSheply := VectorSheply(charFunc)
	fmt.Printf("Vector Sheply = %.2f\n", vectSheply)
}

func VectorSheply(charFunc map[string]int) (vectorSheply []float64) {
	vectorSheply = make([]float64, 0, 4)
	for i := 1; i <= 4; i++ {
		vectorComponent := 0.0
		for setStr, value := range charFunc {
			set := StringToSet(setStr)
			if set.Contains(float64(i)) {
				setMinusI := StringToSet(setStr)
				setMinusI.Remove(float64(i))
				setMinusIStr := SetToString(setMinusI)
				setMinusIVal := charFunc[setMinusIStr]
				vectorComponent += float64(Fact(set.Size() - 1)) * float64(Fact(4 - set.Size())) * float64(value - setMinusIVal)
			}
		}

		vectorComponent /= float64(Fact(4))
		vectorSheply = append(vectorSheply, vectorComponent)
	}

	return vectorSheply
}


func Union(set1, set2 *hashset.Set) *hashset.Set {
	var union *hashset.Set
	union = hashset.New(set1.Values()...)
	union.Add(set2.Values()...)

	return union
}

func Intersection(set1, set2 *hashset.Set) *hashset.Set {
	var left, intersection *hashset.Set
	left = hashset.New(set1.Values()...)
	left.Remove(set2.Values()...)
	intersection = hashset.New(set1.Values()...)
	intersection.Remove(left.Values()...)

	return intersection
}

func IsSubset(set1, set2  *hashset.Set) bool {
	return set2.Contains(set1.Values()...)
}

func SetToString(set *hashset.Set) string {
	bytes, _ := set.ToJSON()
	arrIface := make([]interface{},0)
	_ = json.Unmarshal(bytes, &arrIface)

	arrInts := make([]int, len(arrIface))
	for i, vIface := range arrIface {
		arrInts[i] = int(vIface.(float64))
	}
	sort.Ints(arrInts)
	jsonBytes, _ := json.Marshal(&arrInts)

	return string(jsonBytes)
}

func Fact(num int) int {
	if num == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	return result
}

func StringToSet(str string) *hashset.Set {
	set := hashset.New()
	_ = set.FromJSON([]byte(str))

	return set
}

type ByLenAndValues []string
func (a ByLenAndValues) Len() int           { return len(a) }
func (a ByLenAndValues) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLenAndValues) Less(i, j int) bool {
	if len(a[i]) < len(a[j])  {
		return true
	} else if len(a[i]) == len(a[j]) {
		return a[i] < a[j]
	} else {
		return false
	}
}
