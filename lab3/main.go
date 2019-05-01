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
	X 				float64
	Y 				float64
	Z 				float64
	RadialDistance	float64
	ZenithAngle		float64
	AzimuthAngle	float64
}

func (point *Point) GetDistance(otherPoint *Point) (distance float64) {
	distance = math.Sqrt(math.Pow(point.X - otherPoint.X,2) + math.Pow(point.Y - otherPoint.Y,2) + math.Pow(point.Z - otherPoint.Z,2))

	return distance
}

type Sphere struct {
	Radius	float64
	Center	*Point
}

func (s *Sphere) GeneratePoint() *Point {
	z := GetRandFloat(s.Center.Z - s.Radius, s.Center.Z + s.Radius)
	azimuthAngle := GetRandFloat(0, 2 * math.Pi)
	zenithAngle := math.Acos(z / s.Radius)
	point := GetPointFromSpherical(s.Radius, azimuthAngle, zenithAngle)

	return point
}

type Game struct {
	Apoints		[]Point
	Bpoints		  Point
	PointsNum	  int
	Epsilon		  float64
	Sphere        Sphere
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	counterWinsA := 0; counterWinsB := 0; gamesNum := 100000; pointsNum := 10; epsilon := 0.1; radius := 1.0
	center := &Point{X: 0, Y: 0, Z: 0}
	sphere := &Sphere{radius,center}
	for i := 1; i <= gamesNum; i++ {
		result := MakeGame(pointsNum, epsilon, sphere)
		if result > 0 {
			counterWinsA++
		} else {
			counterWinsB++
		}
	}
	fmt.Println("Параметры игры:")
	fmt.Println(" Центр сферы:(", center.X, ",", center.Y, ",", center.Z,")\n", "Радиус сферы:", radius, "\n Количество точек:", pointsNum, "\n Максимальное расстояние:", epsilon)

	fmt.Println("\nСтатистика:\n", "Количество игр:", gamesNum)
	fmt.Println(" Количество побед игрока А:", counterWinsA, fmt.Sprintf("или %#.2f%%", (float64(counterWinsA) / float64(gamesNum)) * 100))
	fmt.Println(" Количество побед игрока B:", counterWinsB, fmt.Sprintf("или %#.2f%%", (float64(counterWinsB) / float64(gamesNum)) * 100))
}

func MakeGame(pointsNum int, epsilon float64, sphere *Sphere) int {
	pointB := sphere.GeneratePoint()
	result := 0
	for i := 0; i < pointsNum; i++ {
		pointA := sphere.GeneratePoint()
		distanceFromBtoA := pointB.GetDistance(pointA)
		if distanceFromBtoA <= epsilon {
			result++
		}
	}

	return result
}

func GetPointFromSpherical(RadialDistance float64, AzimuthAngle float64, ZenithAngle float64) *Point {
	point := &Point{}
	point.X = RadialDistance * math.Sin(ZenithAngle) * math.Cos(AzimuthAngle)
	point.Y = RadialDistance * math.Sin(ZenithAngle) * math.Sin(AzimuthAngle)
	point.Z = RadialDistance * math.Cos(ZenithAngle)

	return point
}
