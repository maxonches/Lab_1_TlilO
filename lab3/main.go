package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GetRandFloat(min, max float64) float64 {
	return rand.Float64() * (max - min) + min
}

type Point struct {
	X float64
	Y float64
	Z float64
	RadialDistance float64
	ZenithAngle float64
	AzimuthAngle float64
	PointsA []Point
}

func (point *Point) GetDistance(otherPoint *Point) (distance float64) {
	distance = math.Sqrt(math.Pow(point.X - otherPoint.X,2) + math.Pow(point.Y - otherPoint.Y,2) + math.Pow(point.Z - otherPoint.Z,2))

	return distance
}

type Sphere struct {
	Radius float64
	Center *Point
}

func (s *Sphere) GeneratePoint() *Point {
	z := GetRandFloat(s.Center.Z - s.Radius, s.Center.Z + s.Radius)
	azimuthAngle := GetRandFloat(0, 2 * math.Pi)
	zenithAngle := math.Acos(z / s.Radius)
	point := GetPointFromSpherical(s.Radius, azimuthAngle, zenithAngle)

	return point
}

type Game struct {
	Apoints []Point
	Bpoints Point
	PointsNum int
	Epsilon float64
	Sphere Sphere
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	pointsNumA := 3; counterWinsA := 0; counterWinsB := 0; gamesNum := 1000; epsilon := 0.4; radius := 1.0
	center := &Point{X: 0, Y: 0, Z: 0}
	sphere := &Sphere{radius,center}
	var gameCostBetter float64
	for i := 0; i < gamesNum; i++ {
		counterWinsA, counterWinsB = GameCost(pointsNumA, gamesNum, epsilon, sphere)
		gameCostA := float64(counterWinsA)/float64(gamesNum)
		if gameCostA > gameCostBetter {
			gameCostBetter = gameCostA
		}
	}
	gameCostAnalytical := (float64(pointsNumA) / 2.0)*(1.0 - math.Sqrt(1.0 - (math.Pow((epsilon / radius), 2))))

	fmt.Println(" Параметры игры:")
	fmt.Println(" Центр сферы:(", center.X, ",", center.Y, ",", center.Z,")\n", "Радиус сферы:", radius, "\n Количество точек:", pointsNumA, "\n Максимальное расстояние:", epsilon)

	fmt.Println("\nСтатистика:\n", "Количество игр:", gamesNum)
	fmt.Println(" Количество побед игрока А:", counterWinsA, fmt.Sprintf("или %#.2f%%", (float64(counterWinsA) / float64(gamesNum)) * 100))
	fmt.Println(" Количество побед игрока B:", counterWinsB, fmt.Sprintf("или %#.2f%%", (float64(counterWinsB) / float64(gamesNum)) * 100))
	fmt.Println(" Цена игры (аналитически): ", fmt.Sprintf("%.4f", gameCostAnalytical))
	fmt.Println(" Цена игры (численно): ", fmt.Sprintf("%.4f", gameCostBetter))
	fmt.Println(" Погрешность: ", fmt.Sprintf("%.4f", gameCostAnalytical - gameCostBetter))
}

func GameCost(pointsNumA int, gamesNum int, epsilon float64, sphere *Sphere) (int, int) {
	pointsA := []*Point{}
	winsA := 0; winsB := 0

	for i := 0; i < pointsNumA; i++ {
		pointA := sphere.GeneratePoint()
		pointsA = append(pointsA, pointA)
	}

	for j := 0; j < gamesNum; j++ {
		pointB := sphere.GeneratePoint()
		for m := 0; m < pointsNumA; m++ {
			distanceBetweenPoints := pointB.GetDistance(pointsA[m])
			if distanceBetweenPoints <= epsilon {
				winsA++
				break
			} else {
				winsB++
				break
			}
		}
	}

	return winsA, winsB
}

func GetPointFromSpherical(RadialDistance float64, AzimuthAngle float64, ZenithAngle float64) *Point {
	point := &Point{}
	point.X = RadialDistance * math.Sin(ZenithAngle) * math.Cos(AzimuthAngle)
	point.Y = RadialDistance * math.Sin(ZenithAngle) * math.Sin(AzimuthAngle)
	point.Z = RadialDistance * math.Cos(ZenithAngle)

	return point
}